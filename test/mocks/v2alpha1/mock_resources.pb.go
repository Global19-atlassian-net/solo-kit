// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-kit/test/mocks/api/v2alpha1/mock_resources.proto

package v2alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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
//The best mock resource you ever done seen
type MockResource struct {
	Status   core.Status   `protobuf:"bytes,6,opt,name=status,proto3" json:"status"`
	Metadata core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	// Types that are valid to be assigned to WeStuckItInAOneof:
	//	*MockResource_SomeDumbField
	//	*MockResource_Data
	WeStuckItInAOneof isMockResource_WeStuckItInAOneof `protobuf_oneof:"we_stuck_it_in_a_oneof"`
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
	return fileDescriptor_bbc86c81bab68fcb, []int{0}
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

type isMockResource_WeStuckItInAOneof interface {
	isMockResource_WeStuckItInAOneof()
	Equal(interface{}) bool
}
type isMockResource_TestOneofFields interface {
	isMockResource_TestOneofFields()
	Equal(interface{}) bool
}

type MockResource_SomeDumbField struct {
	SomeDumbField string `protobuf:"bytes,100,opt,name=some_dumb_field,json=someDumbField,proto3,oneof" json:"some_dumb_field,omitempty"`
}
type MockResource_Data struct {
	Data string `protobuf:"bytes,1,opt,name=data,json=data.json,proto3,oneof" json:"data.json"`
}
type MockResource_OneofOne struct {
	OneofOne string `protobuf:"bytes,3,opt,name=oneof_one,json=oneofOne,proto3,oneof" json:"oneof_one,omitempty"`
}
type MockResource_OneofTwo struct {
	OneofTwo bool `protobuf:"varint,2,opt,name=oneof_two,json=oneofTwo,proto3,oneof" json:"oneof_two,omitempty"`
}

func (*MockResource_SomeDumbField) isMockResource_WeStuckItInAOneof() {}
func (*MockResource_Data) isMockResource_WeStuckItInAOneof()          {}
func (*MockResource_OneofOne) isMockResource_TestOneofFields()        {}
func (*MockResource_OneofTwo) isMockResource_TestOneofFields()        {}

func (m *MockResource) GetWeStuckItInAOneof() isMockResource_WeStuckItInAOneof {
	if m != nil {
		return m.WeStuckItInAOneof
	}
	return nil
}
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

func (m *MockResource) GetSomeDumbField() string {
	if x, ok := m.GetWeStuckItInAOneof().(*MockResource_SomeDumbField); ok {
		return x.SomeDumbField
	}
	return ""
}

func (m *MockResource) GetData() string {
	if x, ok := m.GetWeStuckItInAOneof().(*MockResource_Data); ok {
		return x.Data
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
		(*MockResource_SomeDumbField)(nil),
		(*MockResource_Data)(nil),
		(*MockResource_OneofOne)(nil),
		(*MockResource_OneofTwo)(nil),
	}
}

type FrequentlyChangingAnnotationsResource struct {
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	Blah                 string        `protobuf:"bytes,1,opt,name=blah,proto3" json:"blah,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FrequentlyChangingAnnotationsResource) Reset()         { *m = FrequentlyChangingAnnotationsResource{} }
func (m *FrequentlyChangingAnnotationsResource) String() string { return proto.CompactTextString(m) }
func (*FrequentlyChangingAnnotationsResource) ProtoMessage()    {}
func (*FrequentlyChangingAnnotationsResource) Descriptor() ([]byte, []int) {
	return fileDescriptor_bbc86c81bab68fcb, []int{1}
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FrequentlyChangingAnnotationsResource.Unmarshal(m, b)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FrequentlyChangingAnnotationsResource.Marshal(b, m, deterministic)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FrequentlyChangingAnnotationsResource.Merge(m, src)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_Size() int {
	return xxx_messageInfo_FrequentlyChangingAnnotationsResource.Size(m)
}
func (m *FrequentlyChangingAnnotationsResource) XXX_DiscardUnknown() {
	xxx_messageInfo_FrequentlyChangingAnnotationsResource.DiscardUnknown(m)
}

var xxx_messageInfo_FrequentlyChangingAnnotationsResource proto.InternalMessageInfo

func (m *FrequentlyChangingAnnotationsResource) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

func (m *FrequentlyChangingAnnotationsResource) GetBlah() string {
	if m != nil {
		return m.Blah
	}
	return ""
}

func init() {
	proto.RegisterType((*MockResource)(nil), "testing.solo.io.MockResource")
	proto.RegisterType((*FrequentlyChangingAnnotationsResource)(nil), "testing.solo.io.FrequentlyChangingAnnotationsResource")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-kit/test/mocks/api/v2alpha1/mock_resources.proto", fileDescriptor_bbc86c81bab68fcb)
}

var fileDescriptor_bbc86c81bab68fcb = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xc3, 0x92, 0x26, 0x0b, 0x55, 0xc5, 0xb6, 0xaa, 0xac, 0x22, 0x68, 0x54, 0x09, 0xc9,
	0x17, 0xbc, 0x4a, 0x91, 0x10, 0xf4, 0x86, 0x41, 0x55, 0x2e, 0x15, 0x92, 0xe1, 0xc4, 0x65, 0xb5,
	0x76, 0x36, 0xce, 0x62, 0x7b, 0x27, 0x78, 0xd7, 0xa4, 0x5c, 0xf3, 0x09, 0x7c, 0x05, 0x9f, 0xc0,
	0x27, 0xf0, 0x11, 0xa8, 0x07, 0x8e, 0xdc, 0x72, 0xe0, 0x8e, 0x76, 0xed, 0x24, 0xe2, 0x82, 0xa2,
	0x9e, 0x3c, 0x33, 0xef, 0xbd, 0x19, 0xcf, 0xd3, 0x2c, 0x1e, 0x67, 0xd2, 0xcc, 0xea, 0x24, 0x4c,
	0xa1, 0xa4, 0x1a, 0x0a, 0x78, 0x2a, 0xa1, 0xf9, 0xe6, 0xd2, 0x50, 0x23, 0xb4, 0xa1, 0x25, 0xa4,
	0xb9, 0xa6, 0x7c, 0x2e, 0xe9, 0xe7, 0x73, 0x5e, 0xcc, 0x67, 0x7c, 0xe4, 0x4a, 0xac, 0x12, 0x1a,
	0xea, 0x2a, 0x15, 0x3a, 0x9c, 0x57, 0x60, 0x80, 0x1c, 0x58, 0xb6, 0x54, 0x59, 0x68, 0xe5, 0xa1,
	0x84, 0x93, 0xd1, 0xff, 0x5a, 0xbb, 0x7e, 0x23, 0x5a, 0x0a, 0xc3, 0x27, 0xdc, 0xf0, 0xa6, 0xc7,
	0x09, 0xdd, 0x41, 0xa2, 0x0d, 0x37, 0x75, 0x3b, 0x74, 0xa7, 0x19, 0xeb, 0xbc, 0x95, 0x1c, 0x65,
	0x90, 0x81, 0x0b, 0xa9, 0x8d, 0xda, 0x2a, 0x11, 0xd7, 0xa6, 0x29, 0x8a, 0xeb, 0x96, 0x79, 0xf6,
	0xb3, 0x8b, 0xef, 0x5f, 0x41, 0x9a, 0xc7, 0xed, 0xa6, 0xe4, 0x39, 0xee, 0x35, 0xd3, 0xfd, 0xde,
	0xd0, 0x0b, 0xee, 0x9d, 0x1f, 0x85, 0x29, 0x54, 0x62, 0xbd, 0x70, 0xf8, 0xce, 0x61, 0x51, 0xff,
	0xfb, 0x1f, 0xe4, 0xfd, 0xb8, 0x39, 0xed, 0xc4, 0x2d, 0x9b, 0xbc, 0xc0, 0xfd, 0xf5, 0xa2, 0xfe,
	0x9e, 0x53, 0x1e, 0xff, 0xab, 0xbc, 0x6a, 0xd1, 0x08, 0x39, 0xdd, 0x86, 0x4d, 0x42, 0x7c, 0xa0,
	0xa1, 0x14, 0x6c, 0x52, 0x97, 0x09, 0x9b, 0x4a, 0x51, 0x4c, 0xfc, 0xc9, 0xd0, 0x0b, 0x06, 0x11,
	0xb2, 0x43, 0xc6, 0x9d, 0x78, 0xdf, 0xc2, 0x6f, 0xea, 0x32, 0xb9, 0xb4, 0x20, 0x09, 0x30, 0x72,
	0x53, 0x3c, 0x47, 0xda, 0xff, 0x7d, 0x73, 0x3a, 0x70, 0xf6, 0x7e, 0xd4, 0xa0, 0xc6, 0x9d, 0x78,
	0x9b, 0x90, 0x47, 0x78, 0x00, 0x4a, 0xc0, 0x94, 0x81, 0x12, 0xfe, 0x1d, 0x4b, 0x1f, 0x7b, 0x71,
	0xdf, 0x95, 0xde, 0x2a, 0xb1, 0x85, 0xcd, 0x02, 0xfc, 0xee, 0xd0, 0x0b, 0xfa, 0x1b, 0xf8, 0xfd,
	0x02, 0x2e, 0x0e, 0x97, 0x2b, 0x84, 0x70, 0xb7, 0xcc, 0x97, 0x2b, 0xb4, 0x47, 0xee, 0xba, 0x1b,
	0x89, 0x7c, 0x7c, 0xbc, 0x10, 0x4c, 0x9b, 0x3a, 0xcd, 0x99, 0x34, 0x4c, 0x2a, 0xc6, 0x99, 0x53,
	0x44, 0x87, 0xf8, 0x81, 0xbd, 0x8e, 0x26, 0x6b, 0xf6, 0xd0, 0x67, 0x5f, 0x3d, 0xfc, 0xe4, 0xb2,
	0x12, 0x9f, 0x6a, 0xa1, 0x4c, 0xf1, 0xe5, 0xf5, 0x8c, 0xab, 0x4c, 0xaa, 0xec, 0x95, 0x52, 0x60,
	0xb8, 0x91, 0xa0, 0xf4, 0xc6, 0xf7, 0xdb, 0xfb, 0x47, 0x30, 0x4a, 0x0a, 0x3e, 0x6b, 0xfc, 0x88,
	0x5d, 0x7c, 0xf1, 0x70, 0xb9, 0x42, 0x3d, 0x8c, 0xa6, 0x29, 0xaf, 0x9a, 0xbf, 0xb7, 0x91, 0x5e,
	0xae, 0x50, 0x37, 0xf0, 0xa2, 0x97, 0xd6, 0xd7, 0x6f, 0xbf, 0x1e, 0x7b, 0x1f, 0xe8, 0x8e, 0x2f,
	0x63, 0xfd, 0x2a, 0x92, 0x9e, 0xbb, 0x9a, 0x67, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x83,
	0xcb, 0x7c, 0x53, 0x03, 0x00, 0x00,
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
	if that1.WeStuckItInAOneof == nil {
		if this.WeStuckItInAOneof != nil {
			return false
		}
	} else if this.WeStuckItInAOneof == nil {
		return false
	} else if !this.WeStuckItInAOneof.Equal(that1.WeStuckItInAOneof) {
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
func (this *MockResource_SomeDumbField) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_SomeDumbField)
	if !ok {
		that2, ok := that.(MockResource_SomeDumbField)
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
	if this.SomeDumbField != that1.SomeDumbField {
		return false
	}
	return true
}
func (this *MockResource_Data) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MockResource_Data)
	if !ok {
		that2, ok := that.(MockResource_Data)
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
	if this.Data != that1.Data {
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
func (this *FrequentlyChangingAnnotationsResource) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FrequentlyChangingAnnotationsResource)
	if !ok {
		that2, ok := that.(FrequentlyChangingAnnotationsResource)
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
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if this.Blah != that1.Blah {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
