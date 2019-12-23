// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test/mocks/api/v1alpha1/mock_resources.proto

//
//package Comments
//package Comments a

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
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

//
//Mock resources for goofin off
type MockResource struct {
	Status        core.Status   `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	Metadata      core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	Data          string        `protobuf:"bytes,1,opt,name=data,json=data.json,proto3" json:"data.json"`
	SomeDumbField string        `protobuf:"bytes,100,opt,name=some_dumb_field,json=someDumbField,proto3" json:"some_dumb_field,omitempty"`
	// Types that are valid to be assigned to TestOneofFields:
	//	*MockResource_OneofOne
	//	*MockResource_OneofTwo
	TestOneofFields      isMockResource_TestOneofFields `protobuf_oneof:"test_oneof_fields"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *MockResource) Reset()         { *m = MockResource{} }
func (m *MockResource) String() string { return proto.CompactTextString(m) }
func (*MockResource) ProtoMessage()    {}
func (*MockResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b69b55b4c8a991d, []int{0}
}
func (m *MockResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MockResource.Unmarshal(m, b)
}
func (m *MockResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MockResource.Marshal(b, m, deterministic)
}
func (m *MockResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MockResource.Merge(m, src)
}
func (m *MockResource) XXX_Size() int {
	return xxx_messageInfo_MockResource.Size(m)
}
func (m *MockResource) XXX_DiscardUnknown() {
	xxx_messageInfo_MockResource.DiscardUnknown(m)
}

var xxx_messageInfo_MockResource proto.InternalMessageInfo

type isMockResource_TestOneofFields interface {
	isMockResource_TestOneofFields()
	Equal(interface{}) bool
}

type MockResource_OneofOne struct {
	OneofOne string `protobuf:"bytes,3,opt,name=oneof_one,json=oneofOne,proto3,oneof" json:"oneof_one,omitempty"`
}
type MockResource_OneofTwo struct {
	OneofTwo bool `protobuf:"varint,2,opt,name=oneof_two,json=oneofTwo,proto3,oneof" json:"oneof_two,omitempty"`
}

func (*MockResource_OneofOne) isMockResource_TestOneofFields() {}
func (*MockResource_OneofTwo) isMockResource_TestOneofFields() {}

func (m *MockResource) GetTestOneofFields() isMockResource_TestOneofFields {
	if m != nil {
		return m.TestOneofFields
	}
	return nil
}

func (m *MockResource) GetStatus() core.Status {
	if m != nil {
		return m.Status
	}
	return core.Status{}
}

func (m *MockResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *MockResource) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *MockResource) GetSomeDumbField() string {
	if m != nil {
		return m.SomeDumbField
	}
	return ""
}

func (m *MockResource) GetOneofOne() string {
	if x, ok := m.GetTestOneofFields().(*MockResource_OneofOne); ok {
		return x.OneofOne
	}
	return ""
}

func (m *MockResource) GetOneofTwo() bool {
	if x, ok := m.GetTestOneofFields().(*MockResource_OneofTwo); ok {
		return x.OneofTwo
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MockResource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MockResource_OneofOne)(nil),
		(*MockResource_OneofTwo)(nil),
	}
}

type FakeResource struct {
	Count                uint32        `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FakeResource) Reset()         { *m = FakeResource{} }
func (m *FakeResource) String() string { return proto.CompactTextString(m) }
func (*FakeResource) ProtoMessage()    {}
func (*FakeResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b69b55b4c8a991d, []int{1}
}
func (m *FakeResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FakeResource.Unmarshal(m, b)
}
func (m *FakeResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FakeResource.Marshal(b, m, deterministic)
}
func (m *FakeResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FakeResource.Merge(m, src)
}
func (m *FakeResource) XXX_Size() int {
	return xxx_messageInfo_FakeResource.Size(m)
}
func (m *FakeResource) XXX_DiscardUnknown() {
	xxx_messageInfo_FakeResource.DiscardUnknown(m)
}

var xxx_messageInfo_FakeResource proto.InternalMessageInfo

func (m *FakeResource) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *FakeResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func init() {
	proto.RegisterType((*MockResource)(nil), "testing.solo.io.MockResource")
	proto.RegisterType((*FakeResource)(nil), "testing.solo.io.FakeResource")
}

func init() {
	proto.RegisterFile("test/mocks/api/v1alpha1/mock_resources.proto", fileDescriptor_6b69b55b4c8a991d)
}

var fileDescriptor_6b69b55b4c8a991d = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xae, 0x83, 0x9b, 0x26, 0x47, 0xa3, 0x0a, 0x37, 0x42, 0x56, 0x54, 0xda, 0xa8, 0x03, 0xca,
	0x50, 0x7c, 0x6a, 0x91, 0x10, 0x30, 0x5a, 0xa8, 0x62, 0xa9, 0x90, 0x0c, 0x13, 0x8b, 0x75, 0xb1,
	0x5f, 0xdc, 0xc3, 0xf1, 0xbd, 0xc8, 0x77, 0x4e, 0xdb, 0xb5, 0x7f, 0x0d, 0x7f, 0x02, 0x13, 0x33,
	0x7f, 0x45, 0x07, 0x46, 0xb6, 0x0e, 0xec, 0xe8, 0x7e, 0xd8, 0x01, 0x24, 0x16, 0xb6, 0x77, 0xef,
	0xfb, 0xbe, 0xbb, 0xef, 0xbb, 0xf7, 0xc8, 0x89, 0x02, 0xa9, 0x68, 0x85, 0x59, 0x29, 0x29, 0x5b,
	0x71, 0xba, 0x3e, 0x65, 0xcb, 0xd5, 0x25, 0x3b, 0x35, 0xad, 0xb4, 0x06, 0x89, 0x4d, 0x9d, 0x81,
	0x8c, 0x56, 0x35, 0x2a, 0x0c, 0xf6, 0x34, 0x9b, 0x8b, 0x22, 0x92, 0xb8, 0xc4, 0x88, 0xe3, 0xe4,
	0x00, 0xc4, 0x1a, 0x6f, 0xac, 0xf2, 0x8c, 0xe6, 0x5c, 0x66, 0xb8, 0x86, 0xfa, 0xc6, 0xd2, 0x27,
	0x07, 0x05, 0x62, 0xb1, 0x04, 0x03, 0x33, 0x21, 0x50, 0x31, 0xc5, 0x51, 0xb8, 0xcb, 0x26, 0xe3,
	0x02, 0x0b, 0x34, 0x25, 0xd5, 0x95, 0xeb, 0x06, 0x70, 0xad, 0x6c, 0x13, 0xae, 0x95, 0xeb, 0x1d,
	0xea, 0xe7, 0x9e, 0x95, 0x5c, 0x39, 0x8b, 0xb4, 0x02, 0xc5, 0x72, 0xa6, 0x58, 0xfb, 0xce, 0xdf,
	0xb8, 0x54, 0x4c, 0x35, 0xf2, 0x5f, 0xea, 0xf6, 0x6c, 0xf1, 0xe3, 0xaf, 0x3d, 0xb2, 0x7b, 0x81,
	0x59, 0x99, 0xb8, 0xb0, 0xc1, 0x0b, 0xd2, 0xb7, 0x17, 0x84, 0xfd, 0xa9, 0x37, 0x7b, 0x78, 0x36,
	0x8e, 0x32, 0xac, 0xa1, 0xcd, 0x1c, 0xbd, 0x37, 0x58, 0x3c, 0xf8, 0xf2, 0xd3, 0xf7, 0xbe, 0xdd,
	0x1d, 0x6d, 0x25, 0x8e, 0x1d, 0xbc, 0x24, 0x83, 0xd6, 0x58, 0xb8, 0x63, 0x94, 0x8f, 0xff, 0x54,
	0x5e, 0x38, 0x34, 0xf6, 0x8d, 0xae, 0x63, 0x07, 0x4f, 0x89, 0x6f, 0x54, 0xde, 0xd4, 0x9b, 0x0d,
	0xe3, 0xd1, 0x8f, 0xbb, 0xa3, 0xa1, 0x89, 0xf7, 0x49, 0xa2, 0x48, 0x36, 0x65, 0x70, 0x42, 0xf6,
	0x24, 0x56, 0x90, 0xe6, 0x4d, 0x35, 0x4f, 0x17, 0x1c, 0x96, 0x79, 0x98, 0x1b, 0x89, 0xaf, 0xcd,
	0x24, 0x23, 0x0d, 0xbe, 0x69, 0xaa, 0xf9, 0xb9, 0x86, 0x82, 0x27, 0x64, 0x88, 0x02, 0x70, 0x91,
	0xa2, 0x80, 0xf0, 0x81, 0xe6, 0xbd, 0xdd, 0x4a, 0x06, 0xa6, 0xf5, 0x4e, 0xc0, 0x06, 0x56, 0x57,
	0x18, 0xf6, 0xa6, 0xde, 0x6c, 0xd0, 0xc1, 0x1f, 0xae, 0xf0, 0xf5, 0xfe, 0xed, 0xbd, 0xef, 0x93,
	0x5e, 0x55, 0xde, 0xde, 0xfb, 0x3b, 0xc1, 0xb6, 0x59, 0x91, 0x78, 0x9f, 0x3c, 0xd2, 0x2b, 0x90,
	0x5a, 0xa1, 0x71, 0x20, 0x8f, 0x25, 0xd9, 0x3d, 0x67, 0x25, 0x74, 0xff, 0x37, 0x26, 0xdb, 0x19,
	0x36, 0x42, 0x99, 0x38, 0xa3, 0xc4, 0x1e, 0xfe, 0xff, 0x77, 0x5a, 0x27, 0x0b, 0xe7, 0x64, 0xc1,
	0x4a, 0x90, 0xf1, 0x2b, 0x9d, 0xf8, 0xf3, 0xf7, 0x43, 0xef, 0x23, 0x2d, 0xb8, 0xba, 0x6c, 0xe6,
	0x51, 0x86, 0x95, 0x1d, 0x2d, 0xc7, 0x6e, 0xc4, 0xf4, 0xb7, 0xf5, 0x6e, 0x57, 0x7b, 0xde, 0x37,
	0x73, 0x7f, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x33, 0x14, 0x83, 0x7b, 0xfc, 0x02, 0x00, 0x00,
}

func (this *MockResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource)
	if !ok {
		that2, ok := that.(MockResource)
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
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	if this.SomeDumbField != that1.SomeDumbField {
		return false
	}
	if that1.TestOneofFields == nil {
		if this.TestOneofFields != nil {
			return false
		}
	} else if this.TestOneofFields == nil {
		return false
	} else if !this.TestOneofFields.Equal(that1.TestOneofFields) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MockResource_OneofOne) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_OneofOne)
	if !ok {
		that2, ok := that.(MockResource_OneofOne)
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
	if this.OneofOne != that1.OneofOne {
		return false
	}
	return true
}
func (this *MockResource_OneofTwo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_OneofTwo)
	if !ok {
		that2, ok := that.(MockResource_OneofTwo)
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
	if this.OneofTwo != that1.OneofTwo {
		return false
	}
	return true
}
func (this *FakeResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FakeResource)
	if !ok {
		that2, ok := that.(FakeResource)
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
	if this.Count != that1.Count {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
