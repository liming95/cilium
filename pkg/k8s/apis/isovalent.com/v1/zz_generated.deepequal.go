//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by deepequal-gen. DO NOT EDIT.

package v1

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EgressGroup) DeepEqual(other *EgressGroup) bool {
	if other == nil {
		return false
	}

	if (in.NodeSelector == nil) != (other.NodeSelector == nil) {
		return false
	} else if in.NodeSelector != nil {
		if !in.NodeSelector.DeepEqual(other.NodeSelector) {
			return false
		}
	}

	if in.Interface != other.Interface {
		return false
	}
	if in.EgressIP != other.EgressIP {
		return false
	}
	if in.MaxGatewayNodes != other.MaxGatewayNodes {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *EgressRule) DeepEqual(other *EgressRule) bool {
	if other == nil {
		return false
	}

	if (in.NamespaceSelector == nil) != (other.NamespaceSelector == nil) {
		return false
	} else if in.NamespaceSelector != nil {
		if !in.NamespaceSelector.DeepEqual(other.NamespaceSelector) {
			return false
		}
	}

	if (in.PodSelector == nil) != (other.PodSelector == nil) {
		return false
	} else if in.PodSelector != nil {
		if !in.PodSelector.DeepEqual(other.PodSelector) {
			return false
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IsovalentEgressGatewayPolicy) DeepEqual(other *IsovalentEgressGatewayPolicy) bool {
	if other == nil {
		return false
	}

	if !in.Spec.DeepEqual(&other.Spec) {
		return false
	}

	if !in.Status.DeepEqual(&other.Status) {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IsovalentEgressGatewayPolicyGroupStatus) DeepEqual(other *IsovalentEgressGatewayPolicyGroupStatus) bool {
	if other == nil {
		return false
	}

	if ((in.ActiveGatewayIPs != nil) && (other.ActiveGatewayIPs != nil)) || ((in.ActiveGatewayIPs == nil) != (other.ActiveGatewayIPs == nil)) {
		in, other := &in.ActiveGatewayIPs, &other.ActiveGatewayIPs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.HealthyGatewayIPs != nil) && (other.HealthyGatewayIPs != nil)) || ((in.HealthyGatewayIPs == nil) != (other.HealthyGatewayIPs == nil)) {
		in, other := &in.HealthyGatewayIPs, &other.HealthyGatewayIPs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IsovalentEgressGatewayPolicySpec) DeepEqual(other *IsovalentEgressGatewayPolicySpec) bool {
	if other == nil {
		return false
	}

	if ((in.Selectors != nil) && (other.Selectors != nil)) || ((in.Selectors == nil) != (other.Selectors == nil)) {
		in, other := &in.Selectors, &other.Selectors
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if ((in.DestinationCIDRs != nil) && (other.DestinationCIDRs != nil)) || ((in.DestinationCIDRs == nil) != (other.DestinationCIDRs == nil)) {
		in, other := &in.DestinationCIDRs, &other.DestinationCIDRs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.ExcludedCIDRs != nil) && (other.ExcludedCIDRs != nil)) || ((in.ExcludedCIDRs == nil) != (other.ExcludedCIDRs == nil)) {
		in, other := &in.ExcludedCIDRs, &other.ExcludedCIDRs
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if inElement != (*other)[i] {
					return false
				}
			}
		}
	}

	if ((in.EgressGroups != nil) && (other.EgressGroups != nil)) || ((in.EgressGroups == nil) != (other.EgressGroups == nil)) {
		in, other := &in.EgressGroups, &other.EgressGroups
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	if in.BGPEnabled != other.BGPEnabled {
		return false
	}

	return true
}

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *IsovalentEgressGatewayPolicyStatus) DeepEqual(other *IsovalentEgressGatewayPolicyStatus) bool {
	if other == nil {
		return false
	}

	if in.ObservedGeneration != other.ObservedGeneration {
		return false
	}
	if ((in.GroupStatuses != nil) && (other.GroupStatuses != nil)) || ((in.GroupStatuses == nil) != (other.GroupStatuses == nil)) {
		in, other := &in.GroupStatuses, &other.GroupStatuses
		if other == nil {
			return false
		}

		if len(*in) != len(*other) {
			return false
		} else {
			for i, inElement := range *in {
				if !inElement.DeepEqual(&(*other)[i]) {
					return false
				}
			}
		}
	}

	return true
}