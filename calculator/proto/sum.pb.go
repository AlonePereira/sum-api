// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sum.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CalculatorRequest struct {
	FirstNumber          int64    `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber         int64    `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorRequest) Reset()         { *m = CalculatorRequest{} }
func (m *CalculatorRequest) String() string { return proto.CompactTextString(m) }
func (*CalculatorRequest) ProtoMessage()    {}
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{0}
}

func (m *CalculatorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorRequest.Unmarshal(m, b)
}
func (m *CalculatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorRequest.Marshal(b, m, deterministic)
}
func (m *CalculatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorRequest.Merge(m, src)
}
func (m *CalculatorRequest) XXX_Size() int {
	return xxx_messageInfo_CalculatorRequest.Size(m)
}
func (m *CalculatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorRequest proto.InternalMessageInfo

func (m *CalculatorRequest) GetFirstNumber() int64 {
	if m != nil {
		return m.FirstNumber
	}
	return 0
}

func (m *CalculatorRequest) GetSecondNumber() int64 {
	if m != nil {
		return m.SecondNumber
	}
	return 0
}

type CalculatorResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalculatorResponse) Reset()         { *m = CalculatorResponse{} }
func (m *CalculatorResponse) String() string { return proto.CompactTextString(m) }
func (*CalculatorResponse) ProtoMessage()    {}
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62743f9cdc99b9fd, []int{1}
}

func (m *CalculatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalculatorResponse.Unmarshal(m, b)
}
func (m *CalculatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalculatorResponse.Marshal(b, m, deterministic)
}
func (m *CalculatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalculatorResponse.Merge(m, src)
}
func (m *CalculatorResponse) XXX_Size() int {
	return xxx_messageInfo_CalculatorResponse.Size(m)
}
func (m *CalculatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalculatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalculatorResponse proto.InternalMessageInfo

func (m *CalculatorResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalculatorRequest)(nil), "calculator.CalculatorRequest")
	proto.RegisterType((*CalculatorResponse)(nil), "calculator.CalculatorResponse")
}

func init() { proto.RegisterFile("sum.proto", fileDescriptor_62743f9cdc99b9fd) }

var fileDescriptor_62743f9cdc99b9fd = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2e, 0xcd, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x4e, 0xcc, 0x49, 0x2e, 0xcd, 0x49, 0x2c, 0xc9,
	0x2f, 0x52, 0x8a, 0xe6, 0x12, 0x74, 0x86, 0xf3, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84,
	0x14, 0xb9, 0x78, 0xd2, 0x32, 0x8b, 0x8a, 0x4b, 0xe2, 0xf3, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xb8, 0xc1, 0x62, 0x7e, 0x60, 0x21, 0x21, 0x65, 0x2e, 0xde,
	0xe2, 0xd4, 0xe4, 0xfc, 0xbc, 0x14, 0x98, 0x1a, 0x26, 0xb0, 0x1a, 0x1e, 0x88, 0x20, 0x44, 0x91,
	0x92, 0x0e, 0x97, 0x10, 0xb2, 0xe1, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x62, 0x5c, 0x6c,
	0x45, 0xa9, 0xc5, 0xa5, 0x39, 0x25, 0x50, 0x73, 0xa1, 0x3c, 0xa3, 0x78, 0x64, 0xa7, 0x04, 0xa7,
	0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x79, 0x71, 0x71, 0xc2, 0x04, 0x53, 0x85, 0x64, 0xf5, 0x10,
	0x2e, 0xd7, 0xc3, 0x70, 0xb6, 0x94, 0x1c, 0x2e, 0x69, 0x88, 0xc5, 0x4e, 0x4a, 0x51, 0x0a, 0xc5,
	0xa5, 0xb9, 0xba, 0x89, 0x05, 0x99, 0xfa, 0x08, 0x85, 0xfa, 0xe0, 0x30, 0xb1, 0x2e, 0x86, 0xd8,
	0x97, 0xc4, 0x06, 0xe6, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x97, 0x6b, 0x3c, 0x71, 0x2f,
	0x01, 0x00, 0x00,
}
