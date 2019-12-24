// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/external/envoy/type/percent.proto

package _type

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

// Fraction percentages support several fixed denominator values.
type FractionalPercent_DenominatorType int32

const (
	// 100.
	//
	// **Example**: 1/100 = 1%.
	FractionalPercent_HUNDRED FractionalPercent_DenominatorType = 0
	// 10,000.
	//
	// **Example**: 1/10000 = 0.01%.
	FractionalPercent_TEN_THOUSAND FractionalPercent_DenominatorType = 1
	// 1,000,000.
	//
	// **Example**: 1/1000000 = 0.0001%.
	FractionalPercent_MILLION FractionalPercent_DenominatorType = 2
)

var FractionalPercent_DenominatorType_name = map[int32]string{
	0: "HUNDRED",
	1: "TEN_THOUSAND",
	2: "MILLION",
}

var FractionalPercent_DenominatorType_value = map[string]int32{
	"HUNDRED":      0,
	"TEN_THOUSAND": 1,
	"MILLION":      2,
}

func (x FractionalPercent_DenominatorType) String() string {
	return proto.EnumName(FractionalPercent_DenominatorType_name, int32(x))
}

func (FractionalPercent_DenominatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65e1ccc85d552d66, []int{1, 0}
}

// Identifies a percentage, in the range [0.0, 100.0].
type Percent struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Percent) Reset()         { *m = Percent{} }
func (m *Percent) String() string { return proto.CompactTextString(m) }
func (*Percent) ProtoMessage()    {}
func (*Percent) Descriptor() ([]byte, []int) {
	return fileDescriptor_65e1ccc85d552d66, []int{0}
}
func (m *Percent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Percent.Unmarshal(m, b)
}
func (m *Percent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Percent.Marshal(b, m, deterministic)
}
func (m *Percent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Percent.Merge(m, src)
}
func (m *Percent) XXX_Size() int {
	return xxx_messageInfo_Percent.Size(m)
}
func (m *Percent) XXX_DiscardUnknown() {
	xxx_messageInfo_Percent.DiscardUnknown(m)
}

var xxx_messageInfo_Percent proto.InternalMessageInfo

func (m *Percent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

// A fractional percentage is used in cases in which for performance reasons performing floating
// point to integer conversions during randomness calculations is undesirable. The message includes
// both a numerator and denominator that together determine the final fractional value.
//
// * **Example**: 1/100 = 1%.
// * **Example**: 3/10000 = 0.03%.
type FractionalPercent struct {
	// Specifies the numerator. Defaults to 0.
	Numerator uint32 `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	// Specifies the denominator. If the denominator specified is less than the numerator, the final
	// fractional percentage is capped at 1 (100%).
	Denominator          FractionalPercent_DenominatorType `protobuf:"varint,2,opt,name=denominator,proto3,enum=envoy.type.FractionalPercent_DenominatorType" json:"denominator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FractionalPercent) Reset()         { *m = FractionalPercent{} }
func (m *FractionalPercent) String() string { return proto.CompactTextString(m) }
func (*FractionalPercent) ProtoMessage()    {}
func (*FractionalPercent) Descriptor() ([]byte, []int) {
	return fileDescriptor_65e1ccc85d552d66, []int{1}
}
func (m *FractionalPercent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FractionalPercent.Unmarshal(m, b)
}
func (m *FractionalPercent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FractionalPercent.Marshal(b, m, deterministic)
}
func (m *FractionalPercent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractionalPercent.Merge(m, src)
}
func (m *FractionalPercent) XXX_Size() int {
	return xxx_messageInfo_FractionalPercent.Size(m)
}
func (m *FractionalPercent) XXX_DiscardUnknown() {
	xxx_messageInfo_FractionalPercent.DiscardUnknown(m)
}

var xxx_messageInfo_FractionalPercent proto.InternalMessageInfo

func (m *FractionalPercent) GetNumerator() uint32 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *FractionalPercent) GetDenominator() FractionalPercent_DenominatorType {
	if m != nil {
		return m.Denominator
	}
	return FractionalPercent_HUNDRED
}

func init() {
	proto.RegisterEnum("envoy.type.FractionalPercent_DenominatorType", FractionalPercent_DenominatorType_name, FractionalPercent_DenominatorType_value)
	proto.RegisterType((*Percent)(nil), "envoy.type.Percent")
	proto.RegisterType((*FractionalPercent)(nil), "envoy.type.FractionalPercent")
}

func init() {
	proto.RegisterFile("api/external/envoy/type/percent.proto", fileDescriptor_65e1ccc85d552d66)
}

var fileDescriptor_65e1ccc85d552d66 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0xcd, 0x2b, 0xcb, 0xaf, 0xd4, 0x2f,
	0xa9, 0x2c, 0x48, 0xd5, 0x2f, 0x48, 0x2d, 0x4a, 0x4e, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x02, 0xcb, 0xe8, 0x81, 0x64, 0xa4, 0xc4, 0xcb, 0x12, 0x73, 0x32, 0x53, 0x12,
	0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x22, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f,
	0xc4, 0x82, 0x8a, 0x0a, 0xa5, 0x56, 0x94, 0x40, 0x04, 0x53, 0x2b, 0xa0, 0xc6, 0x29, 0x59, 0x71,
	0xb1, 0x07, 0x40, 0xcc, 0x17, 0xd2, 0xe7, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x74, 0x92, 0xdc, 0xf5, 0xf2, 0x00, 0xb3, 0x88, 0x90, 0x90, 0x24, 0x03, 0x18,
	0x44, 0x3a, 0x68, 0x32, 0x40, 0x41, 0x10, 0x44, 0x9d, 0xd2, 0x59, 0x46, 0x2e, 0x41, 0xb7, 0xa2,
	0xc4, 0xe4, 0x92, 0xcc, 0xfc, 0xbc, 0xc4, 0x1c, 0x98, 0x31, 0x32, 0x5c, 0x9c, 0x79, 0xa5, 0xb9,
	0xa9, 0x45, 0x89, 0x25, 0xf9, 0x45, 0x60, 0xa3, 0x78, 0x83, 0x10, 0x02, 0x42, 0xd1, 0x5c, 0xdc,
	0x29, 0xa9, 0x79, 0xf9, 0xb9, 0x99, 0x79, 0x60, 0x79, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23, 0x5d,
	0x3d, 0x84, 0xa7, 0xf4, 0x30, 0x4c, 0xd4, 0x73, 0x41, 0x68, 0x08, 0xa9, 0x2c, 0x48, 0x75, 0xe2,
	0x02, 0xb9, 0x8c, 0xb5, 0x89, 0x91, 0x49, 0x80, 0x31, 0x08, 0xd9, 0x34, 0x25, 0x5b, 0x2e, 0x7e,
	0x34, 0xb5, 0x42, 0xdc, 0x5c, 0xec, 0x1e, 0xa1, 0x7e, 0x2e, 0x41, 0xae, 0x2e, 0x02, 0x0c, 0x42,
	0x02, 0x5c, 0x3c, 0x21, 0xae, 0x7e, 0xf1, 0x21, 0x1e, 0xfe, 0xa1, 0xc1, 0x8e, 0x7e, 0x2e, 0x02,
	0x8c, 0x20, 0x69, 0x5f, 0x4f, 0x1f, 0x1f, 0x4f, 0x7f, 0x3f, 0x01, 0x26, 0xa7, 0xec, 0x1d, 0x5f,
	0x59, 0x18, 0x57, 0x3c, 0x92, 0x63, 0xe4, 0x92, 0xc8, 0xcc, 0x87, 0x38, 0xa9, 0xa0, 0x28, 0xbf,
	0xa2, 0x12, 0xc9, 0x75, 0x4e, 0x3c, 0x50, 0x47, 0x05, 0x80, 0x42, 0x2f, 0x80, 0x31, 0xca, 0x3c,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xbf, 0x38, 0x3f, 0x27, 0x5f, 0x37,
	0x33, 0x1f, 0x42, 0x67, 0x67, 0x96, 0xe8, 0x17, 0x64, 0xa7, 0xeb, 0xe3, 0x88, 0xd5, 0x24, 0x36,
	0x70, 0xf8, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x27, 0x9e, 0x2a, 0x0c, 0xf7, 0x01, 0x00,
	0x00,
}

func (this *Percent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Percent)
	if !ok {
		that2, ok := that.(Percent)
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
	if this.Value != that1.Value {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *FractionalPercent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FractionalPercent)
	if !ok {
		that2, ok := that.(FractionalPercent)
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
	if this.Numerator != that1.Numerator {
		return false
	}
	if this.Denominator != that1.Denominator {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
