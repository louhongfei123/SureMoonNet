// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smn_base/smn_base.proto

package smn_base

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

type Call struct {
	Dict                 int32    `protobuf:"varint,1,opt,name=dict,proto3" json:"dict,omitempty"`
	Msg                  []byte   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Call) Reset()         { *m = Call{} }
func (m *Call) String() string { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()    {}
func (*Call) Descriptor() ([]byte, []int) {
	return fileDescriptor_b00eef6af23f64b0, []int{0}
}

func (m *Call) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Call.Unmarshal(m, b)
}
func (m *Call) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Call.Marshal(b, m, deterministic)
}
func (m *Call) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Call.Merge(m, src)
}
func (m *Call) XXX_Size() int {
	return xxx_messageInfo_Call.Size(m)
}
func (m *Call) XXX_DiscardUnknown() {
	xxx_messageInfo_Call.DiscardUnknown(m)
}

var xxx_messageInfo_Call proto.InternalMessageInfo

func (m *Call) GetDict() int32 {
	if m != nil {
		return m.Dict
	}
	return 0
}

func (m *Call) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Ret struct {
	Dict                 int32    `protobuf:"varint,1,opt,name=dict,proto3" json:"dict,omitempty"`
	Err                  bool     `protobuf:"varint,2,opt,name=Err,proto3" json:"Err,omitempty"`
	Msg                  []byte   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ret) Reset()         { *m = Ret{} }
func (m *Ret) String() string { return proto.CompactTextString(m) }
func (*Ret) ProtoMessage()    {}
func (*Ret) Descriptor() ([]byte, []int) {
	return fileDescriptor_b00eef6af23f64b0, []int{1}
}

func (m *Ret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ret.Unmarshal(m, b)
}
func (m *Ret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ret.Marshal(b, m, deterministic)
}
func (m *Ret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ret.Merge(m, src)
}
func (m *Ret) XXX_Size() int {
	return xxx_messageInfo_Ret.Size(m)
}
func (m *Ret) XXX_DiscardUnknown() {
	xxx_messageInfo_Ret.DiscardUnknown(m)
}

var xxx_messageInfo_Ret proto.InternalMessageInfo

func (m *Ret) GetDict() int32 {
	if m != nil {
		return m.Dict
	}
	return 0
}

func (m *Ret) GetErr() bool {
	if m != nil {
		return m.Err
	}
	return false
}

func (m *Ret) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type FPkg struct {
	NO                   int64    `protobuf:"varint,1,opt,name=NO,proto3" json:"NO,omitempty"`
	Msg                  []byte   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Err                  bool     `protobuf:"varint,3,opt,name=Err,proto3" json:"Err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FPkg) Reset()         { *m = FPkg{} }
func (m *FPkg) String() string { return proto.CompactTextString(m) }
func (*FPkg) ProtoMessage()    {}
func (*FPkg) Descriptor() ([]byte, []int) {
	return fileDescriptor_b00eef6af23f64b0, []int{2}
}

func (m *FPkg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FPkg.Unmarshal(m, b)
}
func (m *FPkg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FPkg.Marshal(b, m, deterministic)
}
func (m *FPkg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FPkg.Merge(m, src)
}
func (m *FPkg) XXX_Size() int {
	return xxx_messageInfo_FPkg.Size(m)
}
func (m *FPkg) XXX_DiscardUnknown() {
	xxx_messageInfo_FPkg.DiscardUnknown(m)
}

var xxx_messageInfo_FPkg proto.InternalMessageInfo

func (m *FPkg) GetNO() int64 {
	if m != nil {
		return m.NO
	}
	return 0
}

func (m *FPkg) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *FPkg) GetErr() bool {
	if m != nil {
		return m.Err
	}
	return false
}

func init() {
	proto.RegisterType((*Call)(nil), "smn_base.Call")
	proto.RegisterType((*Ret)(nil), "smn_base.Ret")
	proto.RegisterType((*FPkg)(nil), "smn_base.FPkg")
}

func init() { proto.RegisterFile("smn_base/smn_base.proto", fileDescriptor_b00eef6af23f64b0) }

var fileDescriptor_b00eef6af23f64b0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xce, 0xcd, 0x8b,
	0x4f, 0x4a, 0x2c, 0x4e, 0xd5, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60,
	0x7c, 0x25, 0x1d, 0x2e, 0x16, 0xe7, 0xc4, 0x9c, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xcc, 0xe4,
	0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x30, 0x5b, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38,
	0x5d, 0x82, 0x49, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc4, 0x54, 0xb2, 0xe5, 0x62, 0x0e, 0x4a, 0x2d,
	0xc1, 0xa5, 0xd8, 0xb5, 0xa8, 0x08, 0xac, 0x98, 0x23, 0x08, 0xc4, 0x84, 0x69, 0x67, 0x46, 0x68,
	0xb7, 0xe2, 0x62, 0x71, 0x0b, 0xc8, 0x4e, 0x17, 0xe2, 0xe3, 0x62, 0xf2, 0xf3, 0x07, 0xeb, 0x66,
	0x0e, 0x62, 0xf2, 0xf3, 0xc7, 0xb4, 0x08, 0x66, 0x1a, 0x33, 0xdc, 0xb4, 0x24, 0x36, 0xb0, 0xcb,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xc5, 0x2c, 0x28, 0xd4, 0x00, 0x00, 0x00,
}
