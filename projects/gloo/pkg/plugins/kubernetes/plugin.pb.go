// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: plugin.proto

/*
Package kubernetes is a generated protocol buffer package.

It is generated from these files:
	plugin.proto

It has these top-level messages:
	UpstreamSpec
*/
package kubernetes

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import gloo_plugins_core "github.com/solo-io/solo-kit/projects/gloo/pkg/plugins/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Upstream Spec for Kubernetes Upstreams
// Kubernetes Upstreams represent a set of one or more addressable pods for a Kubernetes Service
// the Gloo Kubernetes Upstream maps to a single service port. Because Kubernetes Services support mulitple ports,
// Gloo requires that a different upstream be created for each port
// Kubernetes Upstreams are typically generated automatically by Gloo from the Kubernetes API
type UpstreamSpec struct {
	// The name of the Kubernetes Service
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The namespace where the Service lives
	ServiceNamespace string `protobuf:"bytes,2,opt,name=service_namespace,json=serviceNamespace,proto3" json:"service_namespace,omitempty"`
	// The port where the Service is listening. If the service only has one port, this can be left empty
	ServicePort int32 `protobuf:"varint,3,opt,name=service_port,json=servicePort,proto3" json:"service_port,omitempty"`
	// Selector allow finer-grained filtering of pods for the Upstream. Gloo will select pods based on their labels if
	// any are provided here.
	// (see [Kubernetes labels and selectors](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/)
	Selector map[string]string `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec *gloo_plugins_core.ServiceSpec `protobuf:"bytes,5,opt,name=service_spec,json=serviceSpec" json:"service_spec,omitempty"`
}

func (m *UpstreamSpec) Reset()                    { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string            { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()               {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) { return fileDescriptorPlugin, []int{0} }

func (m *UpstreamSpec) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *UpstreamSpec) GetServiceNamespace() string {
	if m != nil {
		return m.ServiceNamespace
	}
	return ""
}

func (m *UpstreamSpec) GetServicePort() int32 {
	if m != nil {
		return m.ServicePort
	}
	return 0
}

func (m *UpstreamSpec) GetLabels() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *UpstreamSpec) GetServiceSpec() *gloo_plugins_core.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "gloo.plugins.kubernetes.UpstreamSpec")
}
func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ServiceName != that1.ServiceName {
		return false
	}
	if this.ServiceNamespace != that1.ServiceNamespace {
		return false
	}
	if this.ServicePort != that1.ServicePort {
		return false
	}
	if len(this.Selector) != len(that1.Selector) {
		return false
	}
	for i := range this.Selector {
		if this.Selector[i] != that1.Selector[i] {
			return false
		}
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("plugin.proto", fileDescriptorPlugin) }

var fileDescriptorPlugin = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0x6b, 0x0b, 0x6e, 0x7a, 0xa8, 0x4b, 0xc1, 0xd0, 0x43, 0x89, 0x7a, 0x09, 0x88,
	0xbb, 0x58, 0x2f, 0xea, 0x49, 0x05, 0x05, 0x41, 0x44, 0x5a, 0xbc, 0x78, 0x91, 0x64, 0x19, 0x62,
	0xcc, 0x9f, 0x59, 0x76, 0x37, 0x85, 0xbe, 0x91, 0xcf, 0xe2, 0x63, 0xf8, 0x24, 0xb2, 0x49, 0xda,
	0xae, 0x07, 0x4f, 0x99, 0x19, 0x7e, 0xf9, 0xf6, 0xfb, 0x66, 0xc8, 0x48, 0x16, 0x75, 0x9a, 0x55,
	0x4c, 0x2a, 0x34, 0x48, 0x0f, 0xd3, 0x02, 0x91, 0xb5, 0x23, 0xcd, 0xf2, 0x3a, 0x01, 0x55, 0x81,
	0x01, 0x3d, 0x9d, 0xa4, 0x98, 0x62, 0xc3, 0x70, 0x5b, 0xb5, 0xf8, 0xf4, 0x44, 0xe6, 0x29, 0xef,
	0x68, 0x2e, 0x50, 0x01, 0xd7, 0xa0, 0x56, 0x99, 0x80, 0x77, 0x2d, 0x41, 0xb4, 0xd0, 0xf1, 0x77,
	0x8f, 0x8c, 0x5e, 0xa5, 0x36, 0x0a, 0xe2, 0x72, 0x29, 0x41, 0xd0, 0x23, 0x32, 0xda, 0x60, 0x55,
	0x5c, 0x42, 0xe0, 0x85, 0x5e, 0xb4, 0xbf, 0xf0, 0xbb, 0xd9, 0x73, 0x5c, 0x02, 0x3d, 0x25, 0x07,
	0x2e, 0xa2, 0x65, 0x2c, 0x20, 0xe8, 0x35, 0xdc, 0xd8, 0xe1, 0x9a, 0xb9, 0xab, 0x27, 0x51, 0x99,
	0xa0, 0x1f, 0x7a, 0xd1, 0x60, 0xab, 0xf7, 0x82, 0xca, 0xd0, 0x47, 0x32, 0x2c, 0xe2, 0x04, 0x0a,
	0x1d, 0xec, 0x85, 0xfd, 0xc8, 0x9f, 0x9f, 0xb3, 0x7f, 0x82, 0x32, 0xd7, 0x29, 0x7b, 0x6a, 0xfe,
	0xb9, 0xaf, 0x8c, 0x5a, 0x2f, 0x3a, 0x01, 0x7a, 0xbb, 0x7b, 0xcd, 0x86, 0x0c, 0x06, 0xa1, 0x17,
	0xf9, 0xf3, 0xd9, 0x5f, 0x41, 0xbb, 0x0b, 0xb6, 0x6c, 0x31, 0xab, 0xb4, 0x75, 0x63, 0x9b, 0xe9,
	0x15, 0xf1, 0x1d, 0x65, 0x3a, 0x26, 0xfd, 0x1c, 0xd6, 0xdd, 0x1a, 0x6c, 0x49, 0x27, 0x64, 0xb0,
	0x8a, 0x8b, 0x7a, 0x13, 0xb9, 0x6d, 0xae, 0x7b, 0x97, 0xde, 0xdd, 0xc3, 0xd7, 0xcf, 0xcc, 0x7b,
	0xbb, 0x49, 0x33, 0xf3, 0x51, 0x27, 0x4c, 0x60, 0xc9, 0x35, 0x16, 0x78, 0x96, 0x61, 0xfb, 0xcd,
	0x33, 0xc3, 0xa5, 0xc2, 0x4f, 0x10, 0x46, 0x73, 0xeb, 0x88, 0xbb, 0x17, 0xda, 0xc5, 0x4c, 0x86,
	0xcd, 0x6d, 0x2e, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x46, 0x16, 0x85, 0x94, 0xff, 0x01, 0x00,
	0x00,
}
