// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package envoy

import (
	"context"
	"fmt"
	"strconv"

	envoy_config_core "github.com/cilium/proxy/go/envoy/config/core/v3"
	envoy_config_endpoint "github.com/cilium/proxy/go/envoy/config/endpoint/v3"

	"github.com/cilium/cilium/pkg/loadbalancer"
	"github.com/cilium/cilium/pkg/lock"
	"github.com/cilium/cilium/pkg/logging/logfields"
	"github.com/cilium/cilium/pkg/service"
	"github.com/cilium/cilium/pkg/slices"
)

const anyPort = "*"

// EnvoyServiceBackendSyncer syncs the backends of a Service as Endpoints to the Envoy L7 proxy.
type EnvoyServiceBackendSyncer struct {
	envoyXdsServer XDSServer

	l7lbSvcsMutex lock.RWMutex
	l7lbSvcs      map[loadbalancer.ServiceName]*backendSyncInfo
}

var _ service.BackendSyncer = &EnvoyServiceBackendSyncer{}

func (*EnvoyServiceBackendSyncer) ProxyName() string {
	return "Envoy"
}

func newEnvoyServiceBackendSyncer(envoyXdsServer XDSServer) *EnvoyServiceBackendSyncer {
	return &EnvoyServiceBackendSyncer{
		envoyXdsServer: envoyXdsServer,
		l7lbSvcs:       map[loadbalancer.ServiceName]*backendSyncInfo{},
	}
}

func (r *EnvoyServiceBackendSyncer) Sync(svc *loadbalancer.SVC) error {
	r.l7lbSvcsMutex.RLock()
	l7lbInfo, exists := r.l7lbSvcs[svc.Name]
	r.l7lbSvcsMutex.RUnlock()

	if !exists {
		return nil
	}

	// Filter backend based on list of port numbers, then upsert backends
	// as Envoy endpoints
	be := filterServiceBackends(svc, l7lbInfo.GetAllFrontendPorts())

	log.
		WithField("filteredBackends", be).
		WithField(logfields.L7LBFrontendPorts, l7lbInfo.GetAllFrontendPorts()).
		Debug("Upsert envoy endpoints")
	if err := r.upsertEnvoyEndpoints(svc.Name, be); err != nil {
		return fmt.Errorf("failed to update backends in Envoy: %w", err)
	}

	return nil
}

func (r *EnvoyServiceBackendSyncer) RegisterServiceUsageInCEC(svcName loadbalancer.ServiceName, resourceName service.L7LBResourceName, frontendPorts []string) {
	r.l7lbSvcsMutex.Lock()
	defer r.l7lbSvcsMutex.Unlock()

	l7lbInfo, exists := r.l7lbSvcs[svcName]

	if !exists {
		l7lbInfo = &backendSyncInfo{}
	}

	if l7lbInfo.backendRefs == nil {
		l7lbInfo.backendRefs = make(map[service.L7LBResourceName]backendSyncCECInfo, 1)
	}

	l7lbInfo.backendRefs[resourceName] = backendSyncCECInfo{
		frontendPorts: frontendPorts,
	}

	r.l7lbSvcs[svcName] = l7lbInfo
}

func (r *EnvoyServiceBackendSyncer) DeregisterServiceUsageInCEC(svcName loadbalancer.ServiceName, resourceName service.L7LBResourceName) bool {
	r.l7lbSvcsMutex.Lock()
	defer r.l7lbSvcsMutex.Unlock()

	l7lbInfo, exists := r.l7lbSvcs[svcName]

	if !exists {
		return false
	}

	if l7lbInfo.backendRefs != nil {
		delete(l7lbInfo.backendRefs, resourceName)
	}

	// Cleanup service if it's no longer used by any CEC
	if len(l7lbInfo.backendRefs) == 0 {
		delete(r.l7lbSvcs, svcName)
		return true
	}

	r.l7lbSvcs[svcName] = l7lbInfo

	return false
}

func (r *EnvoyServiceBackendSyncer) upsertEnvoyEndpoints(serviceName loadbalancer.ServiceName, backendMap map[string][]*loadbalancer.Backend) error {
	var resources Resources

	resources.Endpoints = getEndpointsForLBBackends(serviceName, backendMap)

	// Using context.TODO() is fine as we do not upsert listener resources here - the
	// context ends up being used only if listener(s) are included in 'resources'.
	return r.envoyXdsServer.UpsertEnvoyResources(context.TODO(), resources)
}

func getEndpointsForLBBackends(serviceName loadbalancer.ServiceName, backendMap map[string][]*loadbalancer.Backend) []*envoy_config_endpoint.ClusterLoadAssignment {
	var endpoints []*envoy_config_endpoint.ClusterLoadAssignment

	for port, bes := range backendMap {
		var lbEndpoints []*envoy_config_endpoint.LbEndpoint
		for _, be := range bes {
			if be.Protocol != loadbalancer.TCP {
				// Only TCP services supported with Envoy for now
				continue
			}

			lbEndpoints = append(lbEndpoints, &envoy_config_endpoint.LbEndpoint{
				HostIdentifier: &envoy_config_endpoint.LbEndpoint_Endpoint{
					Endpoint: &envoy_config_endpoint.Endpoint{
						Address: &envoy_config_core.Address{
							Address: &envoy_config_core.Address_SocketAddress{
								SocketAddress: &envoy_config_core.SocketAddress{
									Address: be.L3n4Addr.AddrCluster.String(),
									PortSpecifier: &envoy_config_core.SocketAddress_PortValue{
										PortValue: uint32(be.L3n4Addr.L4Addr.Port),
									},
								},
							},
						},
					},
				},
			})
		}

		endpoint := &envoy_config_endpoint.ClusterLoadAssignment{
			ClusterName: fmt.Sprintf("%s:%s", serviceName.String(), port),
			Endpoints: []*envoy_config_endpoint.LocalityLbEndpoints{
				{
					LbEndpoints: lbEndpoints,
				},
			},
		}
		endpoints = append(endpoints, endpoint)

		// for backward compatibility, if any port is allowed, publish one more
		// endpoint having cluster name as service name.
		if port == anyPort {
			endpoints = append(endpoints, &envoy_config_endpoint.ClusterLoadAssignment{
				ClusterName: serviceName.String(),
				Endpoints: []*envoy_config_endpoint.LocalityLbEndpoints{
					{
						LbEndpoints: lbEndpoints,
					},
				},
			})
		}
	}

	return endpoints
}

// filterServiceBackends returns the list of backends based on given front end ports.
// The returned map will have key as port name/number, and value as list of respective backends.
func filterServiceBackends(svc *loadbalancer.SVC, onlyPorts []string) map[string][]*loadbalancer.Backend {
	if len(onlyPorts) == 0 {
		return map[string][]*loadbalancer.Backend{
			"*": filterPreferredBackends(svc.Backends),
		}
	}

	res := map[string][]*loadbalancer.Backend{}
	for _, port := range onlyPorts {
		// check for port number
		if port == strconv.Itoa(int(svc.Frontend.Port)) {
			return map[string][]*loadbalancer.Backend{
				port: filterPreferredBackends(svc.Backends),
			}
		}
		// check for either named port
		for _, backend := range filterPreferredBackends(svc.Backends) {
			if port == backend.FEPortName {
				res[port] = append(res[port], backend)
			}
		}
	}

	return res
}

// filterPreferredBackends returns the slice of backends which are preferred for the given service.
// If there is no preferred backend, it returns the slice of all backends.
func filterPreferredBackends(backends []*loadbalancer.Backend) []*loadbalancer.Backend {
	res := []*loadbalancer.Backend{}
	for _, backend := range backends {
		if backend.Preferred == loadbalancer.Preferred(true) {
			res = append(res, backend)
		}
	}
	if len(res) > 0 {
		return res
	}

	return backends
}

type backendSyncInfo struct {
	// Names of the L7 LB resources (e.g. CEC) that need this service's backends to be
	// synced to to an L7 Loadbalancer.
	backendRefs map[service.L7LBResourceName]backendSyncCECInfo
}

func (r *backendSyncInfo) GetAllFrontendPorts() []string {
	allPorts := []string{}

	for _, info := range r.backendRefs {
		allPorts = append(allPorts, info.frontendPorts...)
	}

	return slices.SortedUnique(allPorts)
}

type backendSyncCECInfo struct {
	// List of front-end ports of upstream service/cluster, which will be used for
	// filtering applicable endpoints.
	//
	// If nil, all the available backends will be used.
	frontendPorts []string
}
