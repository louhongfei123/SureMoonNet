// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dict.proto

/*
Package dict is a generated protocol buffer package.

It is generated from these files:
	dict.proto

It has these top-level messages:
*/
package dict

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

type EDict int32

const (
	EDict_None                        EDict = 0
	EDict_base_Call                   EDict = 1
	EDict_base_Qnm                    EDict = 2
	EDict_rpc_itf_Rpc_Itf_DoLogin_Prm EDict = 5
	EDict_rpc_itf_Rpc_Itf_DoLogin_Ret EDict = 6
)

var EDict_name = map[int32]string{
	0: "None",
	1: "base_Call",
	2: "base_Qnm",
	5: "rpc_itf_Rpc_Itf_DoLogin_Prm",
	6: "rpc_itf_Rpc_Itf_DoLogin_Ret",
}
var EDict_value = map[string]int32{
	"None":                        0,
	"base_Call":                   1,
	"base_Qnm":                    2,
	"rpc_itf_Rpc_Itf_DoLogin_Prm": 5,
	"rpc_itf_Rpc_Itf_DoLogin_Ret": 6,
}

func (x EDict) String() string {
	return proto.EnumName(EDict_name, int32(x))
}
func (EDict) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("dict.EDict", EDict_name, EDict_value)
}

func init() { proto.RegisterFile("dict.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xc9, 0x4c, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xb5, 0x0a, 0xb8, 0x58, 0x5d, 0x5d,
	0x32, 0x93, 0x4b, 0x84, 0x38, 0xb8, 0x58, 0xfc, 0xf2, 0xf3, 0x52, 0x05, 0x18, 0x84, 0x78, 0xb9,
	0x38, 0x93, 0x12, 0x8b, 0x53, 0xe3, 0x9d, 0x13, 0x73, 0x72, 0x04, 0x18, 0x85, 0x78, 0xb8, 0x38,
	0xc0, 0xdc, 0xc0, 0xbc, 0x5c, 0x01, 0x26, 0x21, 0x79, 0x2e, 0xe9, 0xa2, 0x82, 0xe4, 0xf8, 0xcc,
	0x92, 0xb4, 0xf8, 0xa0, 0x82, 0xe4, 0x78, 0xcf, 0x92, 0xb4, 0x78, 0x97, 0x7c, 0x9f, 0xfc, 0xf4,
	0xcc, 0xbc, 0xf8, 0x80, 0xa2, 0x5c, 0x01, 0x56, 0x7c, 0x0a, 0x82, 0x52, 0x4b, 0x04, 0xd8, 0x92,
	0xd8, 0xc0, 0xd6, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xdd, 0xa8, 0xfa, 0x8c, 0x00,
	0x00, 0x00,
}
