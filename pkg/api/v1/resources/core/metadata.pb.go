// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/v1/metadata.proto

package core

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
<<<<<<< HEAD
	math "math"
=======
	_ "github.com/solo-io/protoc-gen-ext/extproto"
>>>>>>> 20b36ddfc8ad6a0ec448b704b926d8b37783e09c
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

//*
// Metadata contains general properties of resources for purposes of versioning, annotating, and namespacing.
type Metadata struct {
	//
	//Name of the resource.
	//
	//Names must be unique and follow the following syntax rules:
	//
	//One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Namespace is used for the namespacing of resources.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Cluster indicates the cluster this resource belongs to
	// Cluster is only applicable in certain contexts, e.g. Kubernetes
	// An empty string here refers to the local cluster
	Cluster string `protobuf:"bytes,7,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// An opaque value that represents the internal version of this object that can
	// be used by clients to determine when objects have changed.
	ResourceVersion string `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty" testdiff:"ignore"`
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. Some resources contain `selectors` which
	// can be linked with other resources by their labels
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata.
	Annotations map[string]string `protobuf:"bytes,6,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A sequence number representing a specific generation of the desired state.
	// Currently only populated for resources backed by Kubernetes
	Generation int64 `protobuf:"varint,8,opt,name=generation,proto3" json:"generation,omitempty"`
	//List of objects depended by this object.
	// Currently only populated for resources backed by Kubernetes
	OwnerReferences      []*Metadata_OwnerReference `protobuf:"bytes,9,rep,name=owner_references,json=ownerReferences,proto3" json:"owner_references,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e335a10da6fa6ca, []int{0}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Metadata) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Metadata) GetResourceVersion() string {
	if m != nil {
		return m.ResourceVersion
	}
	return ""
}

func (m *Metadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Metadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *Metadata) GetGeneration() int64 {
	if m != nil {
		return m.Generation
	}
	return 0
}

func (m *Metadata) GetOwnerReferences() []*Metadata_OwnerReference {
	if m != nil {
		return m.OwnerReferences
	}
	return nil
}

// proto message representing kubernertes owner reference
// https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.15/#ownerreference-v1-meta
type Metadata_OwnerReference struct {
	ApiVersion           string           `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	BlockOwnerDeletion   *types.BoolValue `protobuf:"bytes,2,opt,name=block_owner_deletion,json=blockOwnerDeletion,proto3" json:"block_owner_deletion,omitempty"`
	Controller           *types.BoolValue `protobuf:"bytes,3,opt,name=controller,proto3" json:"controller,omitempty"`
	Kind                 string           `protobuf:"bytes,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Name                 string           `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Uid                  string           `protobuf:"bytes,6,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Metadata_OwnerReference) Reset()         { *m = Metadata_OwnerReference{} }
func (m *Metadata_OwnerReference) String() string { return proto.CompactTextString(m) }
func (*Metadata_OwnerReference) ProtoMessage()    {}
func (*Metadata_OwnerReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e335a10da6fa6ca, []int{0, 2}
}
func (m *Metadata_OwnerReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata_OwnerReference.Unmarshal(m, b)
}
func (m *Metadata_OwnerReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata_OwnerReference.Marshal(b, m, deterministic)
}
func (m *Metadata_OwnerReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata_OwnerReference.Merge(m, src)
}
func (m *Metadata_OwnerReference) XXX_Size() int {
	return xxx_messageInfo_Metadata_OwnerReference.Size(m)
}
func (m *Metadata_OwnerReference) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata_OwnerReference.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata_OwnerReference proto.InternalMessageInfo

func (m *Metadata_OwnerReference) GetApiVersion() string {
	if m != nil {
		return m.ApiVersion
	}
	return ""
}

func (m *Metadata_OwnerReference) GetBlockOwnerDeletion() *types.BoolValue {
	if m != nil {
		return m.BlockOwnerDeletion
	}
	return nil
}

func (m *Metadata_OwnerReference) GetController() *types.BoolValue {
	if m != nil {
		return m.Controller
	}
	return nil
}

func (m *Metadata_OwnerReference) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Metadata_OwnerReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metadata_OwnerReference) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto.RegisterType((*Metadata)(nil), "core.solo.io.Metadata")
	proto.RegisterMapType((map[string]string)(nil), "core.solo.io.Metadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.solo.io.Metadata.LabelsEntry")
	proto.RegisterType((*Metadata_OwnerReference)(nil), "core.solo.io.Metadata.OwnerReference")
}

func init() { proto.RegisterFile("api/v1/metadata.proto", fileDescriptor_6e335a10da6fa6ca) }

<<<<<<< HEAD
var fileDescriptor_6e335a10da6fa6ca = []byte{
	// 506 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x67, 0x9b, 0x36, 0x6d, 0x5f, 0xc4, 0xc6, 0xa1, 0x85, 0x21, 0x48, 0x1a, 0x02, 0x62, 0x2e,
	0xee, 0x62, 0x45, 0xd0, 0x08, 0xa2, 0x41, 0x0f, 0x42, 0x45, 0xd9, 0x43, 0x0f, 0x5e, 0xc2, 0x64,
	0xf7, 0x65, 0x1d, 0x32, 0x99, 0xb7, 0xcc, 0xcc, 0xa6, 0xf4, 0xee, 0x87, 0xf1, 0x53, 0x79, 0xf0,
	0x23, 0x08, 0xde, 0x65, 0x67, 0xb2, 0xed, 0x56, 0xa8, 0xe0, 0x69, 0xdf, 0x9b, 0xf7, 0x7b, 0xbf,
	0xf7, 0xef, 0xb7, 0x70, 0x22, 0x4a, 0x99, 0x6c, 0x9e, 0x26, 0x6b, 0x74, 0x22, 0x17, 0x4e, 0xc4,
	0xa5, 0x21, 0x47, 0xec, 0x5e, 0x46, 0x06, 0x63, 0x4b, 0x8a, 0x62, 0x49, 0x83, 0xe3, 0x82, 0x0a,
	0xf2, 0x81, 0xa4, 0xb6, 0x02, 0x66, 0x30, 0x2c, 0x88, 0x0a, 0x85, 0x89, 0xf7, 0x16, 0xd5, 0x32,
	0xb9, 0x34, 0xa2, 0x2c, 0xd1, 0xd8, 0x10, 0x1f, 0x7f, 0xeb, 0xc2, 0xc1, 0xc7, 0x2d, 0x2d, 0x63,
	0xb0, 0xab, 0xc5, 0x1a, 0xf9, 0xce, 0x28, 0x9a, 0x1c, 0xa6, 0xde, 0x66, 0x0f, 0xe1, 0xb0, 0xfe,
	0xda, 0x52, 0x64, 0xc8, 0x3b, 0x3e, 0x70, 0xf3, 0xc0, 0x38, 0xec, 0x67, 0xaa, 0xb2, 0x0e, 0x0d,
	0xdf, 0xf7, 0xb1, 0xc6, 0x65, 0x6f, 0xa0, 0x6f, 0xd0, 0x52, 0x65, 0x32, 0x9c, 0x6f, 0xd0, 0x58,
	0x49, 0x9a, 0xef, 0xd6, 0x90, 0xd9, 0xc9, 0xaf, 0x1f, 0xa7, 0x0f, 0x1c, 0x5a, 0x97, 0xcb, 0xe5,
	0x72, 0x3a, 0x96, 0x85, 0x26, 0x83, 0xe3, 0xf4, 0xa8, 0x81, 0x5f, 0x04, 0x34, 0x9b, 0x42, 0x57,
	0x89, 0x05, 0x2a, 0xcb, 0xf7, 0x46, 0x9d, 0x49, 0xef, 0x6c, 0x1c, 0xb7, 0xe7, 0x8d, 0x9b, 0xae,
	0xe3, 0x73, 0x0f, 0x7a, 0xaf, 0x9d, 0xb9, 0x4a, 0xb7, 0x19, 0xec, 0x03, 0xf4, 0x84, 0xd6, 0xe4,
	0x84, 0x93, 0xa4, 0x2d, 0xef, 0x7a, 0x82, 0xc7, 0x77, 0x10, 0xbc, 0xbd, 0x41, 0x06, 0x96, 0x76,
	0x2e, 0x1b, 0x02, 0x14, 0xa8, 0xd1, 0x78, 0x97, 0x1f, 0x8c, 0xa2, 0x49, 0x27, 0x6d, 0xbd, 0xb0,
	0xcf, 0xd0, 0xa7, 0x4b, 0x8d, 0x66, 0x6e, 0x70, 0x89, 0x06, 0x75, 0x86, 0x96, 0x1f, 0xfa, 0x7a,
	0x8f, 0xee, 0xa8, 0xf7, 0xa9, 0x86, 0xa7, 0x0d, 0x3a, 0x3d, 0xa2, 0x5b, 0xbe, 0x1d, 0xbc, 0x84,
	0x5e, 0x6b, 0x26, 0xd6, 0x87, 0xce, 0x0a, 0xaf, 0x78, 0xe4, 0xf7, 0x5b, 0x9b, 0xec, 0x18, 0xf6,
	0x36, 0x42, 0x55, 0xcd, 0xa1, 0x82, 0x33, 0xdd, 0x79, 0x11, 0x0d, 0x5e, 0x43, 0xff, 0xef, 0x69,
	0xfe, 0x2b, 0xff, 0x77, 0x04, 0xf7, 0x6f, 0xb7, 0xc7, 0x4e, 0xa1, 0x27, 0x4a, 0x79, 0x7d, 0xc3,
	0x40, 0x03, 0xa2, 0x94, 0xcd, 0x9d, 0xce, 0xe1, 0x78, 0xa1, 0x28, 0x5b, 0xcd, 0xc3, 0x1a, 0x72,
	0x54, 0xe8, 0x57, 0x55, 0x93, 0xf7, 0xce, 0x06, 0x71, 0x50, 0x60, 0xdc, 0x28, 0x30, 0x9e, 0x11,
	0xa9, 0x8b, 0xba, 0x62, 0xca, 0x7c, 0x9e, 0xaf, 0xf7, 0x6e, 0x9b, 0xc5, 0xa6, 0x00, 0x19, 0x69,
	0x67, 0x48, 0x29, 0x34, 0x5e, 0x70, 0xff, 0xe6, 0x68, 0xa1, 0x6b, 0xfd, 0xae, 0xa4, 0xce, 0x83,
	0xce, 0x52, 0x6f, 0x5f, 0x6b, 0x7a, 0xaf, 0xa5, 0xe9, 0x3e, 0x74, 0x2a, 0x99, 0xf3, 0x6e, 0xd8,
	0x48, 0x25, 0xf3, 0xd9, 0xab, 0xef, 0x3f, 0x87, 0xd1, 0x97, 0xe7, 0x85, 0x74, 0x5f, 0xab, 0x45,
	0x9c, 0xd1, 0x3a, 0xa9, 0xaf, 0xf6, 0x44, 0x52, 0xf8, 0xae, 0xa4, 0x4b, 0xca, 0x55, 0x91, 0x6c,
	0xff, 0xc3, 0x46, 0xa9, 0x36, 0xa9, 0x0f, 0xbc, 0xe8, 0xfa, 0xb6, 0x9e, 0xfd, 0x09, 0x00, 0x00,
	0xff, 0xff, 0x94, 0x96, 0x58, 0x03, 0xa7, 0x03, 0x00, 0x00,
=======
var fileDescriptor_56d9f74966f40d04 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x67, 0xbb, 0x6d, 0xda, 0x4c, 0xa4, 0x8d, 0x43, 0x0e, 0x43, 0x90, 0x36, 0x44, 0xc4, 0x5c,
	0xdc, 0xc5, 0x8a, 0xa0, 0x39, 0x14, 0x5c, 0xf5, 0x20, 0x54, 0x94, 0x3d, 0xf4, 0xe0, 0x25, 0x4c,
	0x76, 0x5f, 0xd6, 0x21, 0x93, 0x79, 0xcb, 0xcc, 0x6c, 0xda, 0x7e, 0x07, 0x3f, 0x88, 0x9f, 0xca,
	0x83, 0x1f, 0x41, 0x10, 0x72, 0x94, 0x9d, 0xc9, 0xb6, 0x5b, 0xa1, 0x82, 0xa7, 0x79, 0x7f, 0x7e,
	0xbf, 0xf7, 0xde, 0xbc, 0xf7, 0x23, 0x87, 0x2b, 0xb0, 0x3c, 0xe7, 0x96, 0x47, 0xa5, 0x46, 0x8b,
	0xf4, 0x41, 0x86, 0x1a, 0x22, 0x83, 0x12, 0x23, 0x81, 0xc3, 0x41, 0x81, 0x05, 0xba, 0x44, 0x5c,
	0x5b, 0x1e, 0x33, 0xa4, 0x70, 0x65, 0x7d, 0x10, 0xae, 0xec, 0x36, 0x76, 0x5c, 0x20, 0x16, 0x12,
	0x62, 0xe7, 0xcd, 0xab, 0x45, 0x7c, 0xa9, 0x79, 0x59, 0x82, 0x36, 0x3e, 0x3f, 0xfe, 0xd6, 0x21,
	0x07, 0x1f, 0xb7, 0xad, 0x28, 0x25, 0xbb, 0x8a, 0xaf, 0x80, 0xed, 0x8c, 0x82, 0x49, 0x37, 0x75,
	0x36, 0x7d, 0x44, 0xba, 0xf5, 0x6b, 0x4a, 0x9e, 0x01, 0x0b, 0x5d, 0xe2, 0x36, 0x40, 0x19, 0xd9,
	0xcf, 0x64, 0x65, 0x2c, 0x68, 0xb6, 0xef, 0x72, 0x8d, 0x4b, 0xdf, 0x92, 0xbe, 0x06, 0x83, 0x95,
	0xce, 0x60, 0xb6, 0x06, 0x6d, 0x04, 0x2a, 0xb6, 0x5b, 0x43, 0x12, 0xb6, 0x49, 0x82, 0x5f, 0x3f,
	0x4e, 0x1e, 0x5a, 0x30, 0x36, 0x17, 0x8b, 0xc5, 0x74, 0x2c, 0x0a, 0x85, 0x1a, 0xc6, 0xe9, 0x51,
	0xc3, 0xb8, 0xf0, 0x04, 0x3a, 0x25, 0x1d, 0xc9, 0xe7, 0x20, 0x0d, 0xdb, 0x1b, 0x85, 0x93, 0xde,
	0xe9, 0x38, 0x6a, 0xaf, 0x21, 0x6a, 0x06, 0x8f, 0xce, 0x1d, 0xe8, 0xbd, 0xb2, 0xfa, 0x3a, 0xdd,
	0x32, 0xe8, 0x07, 0xd2, 0xe3, 0x4a, 0xa1, 0xe5, 0x56, 0xa0, 0x32, 0xac, 0xe3, 0x0a, 0x3c, 0xbd,
	0xa7, 0xc0, 0x9b, 0x5b, 0xa4, 0xaf, 0xd2, 0xe6, 0xd2, 0xc7, 0x84, 0x14, 0xa0, 0x40, 0x3b, 0x97,
	0x1d, 0x8c, 0x82, 0x49, 0x98, 0x84, 0x9b, 0x24, 0x48, 0x5b, 0x61, 0xfa, 0x99, 0xf4, 0xf1, 0x52,
	0x81, 0x9e, 0x69, 0x58, 0x80, 0x06, 0x95, 0x81, 0x61, 0x5d, 0xd7, 0xf4, 0xc9, 0x3d, 0x4d, 0x3f,
	0xd5, 0xf0, 0xb4, 0x41, 0xa7, 0x47, 0x78, 0xc7, 0x37, 0xc3, 0xd7, 0xa4, 0xd7, 0xfa, 0x18, 0xed,
	0x93, 0x70, 0x09, 0xd7, 0x2c, 0x70, 0x7b, 0xae, 0x4d, 0x3a, 0x20, 0x7b, 0x6b, 0x2e, 0xab, 0xe6,
	0x60, 0xde, 0x99, 0xee, 0xbc, 0x0a, 0x86, 0x67, 0xa4, 0xff, 0xf7, 0x97, 0xfe, 0x8b, 0xff, 0x3b,
	0x20, 0x87, 0x77, 0xc7, 0xa3, 0x27, 0xa4, 0xc7, 0x4b, 0x71, 0x73, 0x4b, 0x5f, 0x86, 0xf0, 0x52,
	0x34, 0xc7, 0x3a, 0x27, 0x83, 0xb9, 0xc4, 0x6c, 0x39, 0xf3, 0x6b, 0xc8, 0x41, 0x82, 0xdb, 0x57,
	0x5d, 0xbc, 0x77, 0x3a, 0x8c, 0xbc, 0x12, 0xa3, 0x46, 0x89, 0x51, 0x82, 0x28, 0x2f, 0xea, 0x8e,
	0x29, 0x75, 0x3c, 0xd7, 0xef, 0xdd, 0x96, 0x45, 0xa7, 0x84, 0x64, 0xa8, 0xac, 0x46, 0x29, 0x41,
	0x3b, 0xe1, 0xfd, 0xbb, 0x46, 0x0b, 0x5d, 0xeb, 0x78, 0x29, 0x54, 0xee, 0xf5, 0x96, 0x3a, 0xfb,
	0x46, 0xdb, 0x7b, 0x2d, 0x6d, 0xf7, 0x49, 0x58, 0x89, 0x9c, 0x75, 0xfc, 0x46, 0x2a, 0x91, 0x27,
	0x67, 0x9b, 0x24, 0xf8, 0xfe, 0xf3, 0x38, 0xf8, 0xf2, 0xb2, 0x10, 0xf6, 0x6b, 0x35, 0x8f, 0x32,
	0x5c, 0xc5, 0xf5, 0xe1, 0x9e, 0x09, 0xf4, 0xef, 0x52, 0xd8, 0xb8, 0x5c, 0x16, 0x31, 0x2f, 0x45,
	0xbc, 0x7e, 0x1e, 0x37, 0x8a, 0x35, 0x71, 0x7d, 0xe3, 0x79, 0xc7, 0x4d, 0xf6, 0xe2, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xd2, 0xdd, 0x90, 0x3c, 0xbf, 0x03, 0x00, 0x00,
>>>>>>> 20b36ddfc8ad6a0ec448b704b926d8b37783e09c
}

func (this *Metadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metadata)
	if !ok {
		that2, ok := that.(Metadata)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if this.Cluster != that1.Cluster {
		return false
	}
	if this.ResourceVersion != that1.ResourceVersion {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if len(this.Annotations) != len(that1.Annotations) {
		return false
	}
	for i := range this.Annotations {
		if this.Annotations[i] != that1.Annotations[i] {
			return false
		}
	}
	if this.Generation != that1.Generation {
		return false
	}
	if len(this.OwnerReferences) != len(that1.OwnerReferences) {
		return false
	}
	for i := range this.OwnerReferences {
		if !this.OwnerReferences[i].Equal(that1.OwnerReferences[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Metadata_OwnerReference) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Metadata_OwnerReference)
	if !ok {
		that2, ok := that.(Metadata_OwnerReference)
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
	if this.ApiVersion != that1.ApiVersion {
		return false
	}
	if !this.BlockOwnerDeletion.Equal(that1.BlockOwnerDeletion) {
		return false
	}
	if !this.Controller.Equal(that1.Controller) {
		return false
	}
	if this.Kind != that1.Kind {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Uid != that1.Uid {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
