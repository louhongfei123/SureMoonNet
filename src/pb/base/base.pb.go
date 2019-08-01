// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

/*
Package base is a generated protocol buffer package.

It is generated from these files:
	base.proto

It has these top-level messages:
	Call
*/
package base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Call struct {
	Dict int32  `protobuf:"varint,1,opt,name=dict" json:"dict,omitempty"`
	Msg  []byte `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func init() {
	proto.RegisterType((*Call)(nil), "base.Call")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x14, 0xb9, 0x58, 0x9c, 0x13,
	0x73, 0x72, 0x84, 0x78, 0xb8, 0x58, 0x52, 0x32, 0x93, 0x4b, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58,
	0x85, 0xb8, 0xb9, 0x98, 0x73, 0x8b, 0xd3, 0x25, 0x98, 0x14, 0x18, 0x35, 0x78, 0x92, 0xd8, 0xc0,
	0xea, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0x0b, 0x99, 0xa1, 0x3d, 0x00, 0x00, 0x00,
}
