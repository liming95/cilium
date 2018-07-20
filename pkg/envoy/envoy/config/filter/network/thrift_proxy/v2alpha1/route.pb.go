// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto
	envoy/config/filter/network/thrift_proxy/v2alpha1/thrift_proxy.proto

It has these top-level messages:
	RouteConfiguration
	Route
	RouteMatch
	RouteAction
	ThriftProxy
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// [#comment:next free field: 3]
type RouteConfiguration struct {
	// The name of the route configuration. Reserved for future use in asynchronous route discovery.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The list of routes that will be matched, in order, against incoming requests. The first route
	// that matches will be used.
	Routes []*Route `protobuf:"bytes,2,rep,name=routes" json:"routes,omitempty"`
}

func (m *RouteConfiguration) Reset()                    { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string            { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()               {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// [#comment:next free field: 3]
type Route struct {
	// Route matching prarameters.
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match" json:"match,omitempty"`
	// Route request to some upstream cluster.
	Route *RouteAction `protobuf:"bytes,2,opt,name=route" json:"route,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

// [#comment:next free field: 2]
type RouteMatch struct {
	// If specified, the route must exactly match the request method name. As a special case, an
	// empty string matches any request method name.
	Method string `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
}

func (m *RouteMatch) Reset()                    { *m = RouteMatch{} }
func (m *RouteMatch) String() string            { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()               {}
func (*RouteMatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RouteMatch) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

// [#comment:next free field: 2]
type RouteAction struct {
	// Indicates the upstream cluster to which the request should be routed.
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *RouteAction) Reset()                    { *m = RouteAction{} }
func (m *RouteAction) String() string            { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()               {}
func (*RouteAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RouteAction) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.thrift_proxy.v2alpha1.RouteAction")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v2alpha1/route.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0x3f, 0xa7, 0x49, 0x3f, 0xf5, 0x2a, 0x21, 0x64, 0x21, 0xa8, 0xba, 0x50, 0x05, 0x86,
	0x4e, 0xb6, 0x1a, 0x16, 0x96, 0x22, 0x11, 0x66, 0x96, 0x0c, 0x0c, 0x2c, 0xc8, 0xa4, 0x6e, 0x13,
	0x91, 0xc6, 0x95, 0xeb, 0x04, 0xba, 0x31, 0xf3, 0x93, 0x98, 0x18, 0x99, 0xf9, 0x01, 0xec, 0xfc,
	0x0b, 0xe4, 0x73, 0x22, 0x2a, 0x31, 0xd1, 0xed, 0xcd, 0x25, 0xef, 0xf3, 0x9c, 0x2e, 0x30, 0x95,
	0x65, 0xad, 0x36, 0x3c, 0x55, 0xe5, 0x3c, 0x5f, 0xf0, 0x79, 0x5e, 0x18, 0xa9, 0x79, 0x29, 0xcd,
	0xa3, 0xd2, 0x0f, 0xdc, 0x64, 0x3a, 0x9f, 0x9b, 0xbb, 0x95, 0x56, 0x4f, 0x1b, 0x5e, 0x47, 0xa2,
	0x58, 0x65, 0x62, 0xc2, 0xb5, 0xaa, 0x8c, 0x64, 0x2b, 0xad, 0x8c, 0xa2, 0x13, 0xac, 0x33, 0x57,
	0x67, 0xae, 0xce, 0x9a, 0x3a, 0xdb, 0xae, 0xb3, 0xb6, 0x3e, 0x3c, 0xaa, 0x45, 0x91, 0xcf, 0x84,
	0x91, 0xbc, 0x0d, 0x8e, 0x35, 0x3c, 0x58, 0xa8, 0x85, 0xc2, 0xc8, 0x6d, 0x72, 0xd3, 0xf0, 0x99,
	0x00, 0x4d, 0xac, 0xf1, 0x0a, 0x1d, 0x95, 0x16, 0x26, 0x57, 0x25, 0xa5, 0xe0, 0x97, 0x62, 0x29,
	0x07, 0x64, 0x44, 0xc6, 0xbd, 0x04, 0x33, 0xbd, 0x81, 0x2e, 0xee, 0xb6, 0x1e, 0x78, 0xa3, 0xce,
	0xb8, 0x1f, 0x9d, 0xb3, 0x3f, 0x6f, 0xc7, 0x50, 0x15, 0xfb, 0xef, 0x9f, 0xc7, 0xff, 0x92, 0x86,
	0x16, 0x7e, 0x10, 0x08, 0x70, 0x4e, 0x53, 0x08, 0x96, 0xc2, 0xa4, 0x19, 0x6a, 0xfb, 0xd1, 0x74,
	0x57, 0xc1, 0xb5, 0x85, 0xc4, 0x7b, 0xd6, 0xf2, 0xfa, 0xf5, 0xd6, 0x09, 0x5e, 0x88, 0xb7, 0x4f,
	0x12, 0xc7, 0xa6, 0x33, 0x08, 0x50, 0x3c, 0xf0, 0x50, 0x72, 0xb1, 0xab, 0xe4, 0x32, 0xb5, 0x97,
	0xfa, 0x6d, 0x41, 0x78, 0x78, 0x0a, 0xf0, 0xb3, 0x0a, 0x3d, 0x84, 0xee, 0x52, 0x9a, 0x4c, 0xcd,
	0x9a, 0x83, 0x36, 0x4f, 0x61, 0x04, 0xfd, 0x2d, 0x16, 0x3d, 0x81, 0xff, 0x69, 0x51, 0xad, 0x8d,
	0xd4, 0xee, 0xbb, 0xb8, 0x67, 0xc1, 0xbe, 0xf6, 0x46, 0x24, 0x69, 0xdf, 0xc4, 0xfe, 0xad, 0x57,
	0x47, 0xf7, 0x5d, 0xfc, 0x7d, 0x67, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80, 0x5c, 0x16, 0xe2,
	0x61, 0x02, 0x00, 0x00,
}
