// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/retry/omit_canary_hosts/v2/omit_canary_hosts.proto

package envoy_config_retry_omit_canary_hosts_v2

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

type OmitCanaryHostsPredicate struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OmitCanaryHostsPredicate) Reset()         { *m = OmitCanaryHostsPredicate{} }
func (m *OmitCanaryHostsPredicate) String() string { return proto.CompactTextString(m) }
func (*OmitCanaryHostsPredicate) ProtoMessage()    {}
func (*OmitCanaryHostsPredicate) Descriptor() ([]byte, []int) {
	return fileDescriptor_08501cb61947b43e, []int{0}
}

func (m *OmitCanaryHostsPredicate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OmitCanaryHostsPredicate.Unmarshal(m, b)
}
func (m *OmitCanaryHostsPredicate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OmitCanaryHostsPredicate.Marshal(b, m, deterministic)
}
func (m *OmitCanaryHostsPredicate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OmitCanaryHostsPredicate.Merge(m, src)
}
func (m *OmitCanaryHostsPredicate) XXX_Size() int {
	return xxx_messageInfo_OmitCanaryHostsPredicate.Size(m)
}
func (m *OmitCanaryHostsPredicate) XXX_DiscardUnknown() {
	xxx_messageInfo_OmitCanaryHostsPredicate.DiscardUnknown(m)
}

var xxx_messageInfo_OmitCanaryHostsPredicate proto.InternalMessageInfo

func init() {
	proto.RegisterType((*OmitCanaryHostsPredicate)(nil), "envoy.config.retry.omit_canary_hosts.v2.OmitCanaryHostsPredicate")
}

func init() {
	proto.RegisterFile("envoy/config/retry/omit_canary_hosts/v2/omit_canary_hosts.proto", fileDescriptor_08501cb61947b43e)
}

var fileDescriptor_08501cb61947b43e = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2f, 0x4a, 0x2d, 0x29, 0xaa, 0xd4, 0xcf,
	0xcf, 0xcd, 0x2c, 0x89, 0x4f, 0x4e, 0xcc, 0x4b, 0x2c, 0xaa, 0x8c, 0xcf, 0xc8, 0x2f, 0x2e, 0x29,
	0xd6, 0x2f, 0x33, 0xc2, 0x14, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x07, 0x1b, 0xa0,
	0x07, 0x31, 0x40, 0x0f, 0x6c, 0x80, 0x1e, 0xa6, 0xda, 0x32, 0x23, 0x25, 0x29, 0x2e, 0x09, 0xff,
	0xdc, 0xcc, 0x12, 0x67, 0xb0, 0xb0, 0x07, 0x48, 0x34, 0xa0, 0x28, 0x35, 0x25, 0x33, 0x39, 0xb1,
	0x24, 0xd5, 0xc9, 0x9f, 0xcb, 0x34, 0x33, 0x5f, 0x0f, 0x6c, 0x52, 0x41, 0x51, 0x7e, 0x45, 0xa5,
	0x1e, 0x91, 0x86, 0x3a, 0x89, 0x60, 0x18, 0x99, 0x5f, 0x92, 0x1f, 0xc0, 0x98, 0xc4, 0x06, 0x76,
	0x9c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x80, 0xa6, 0x23, 0x38, 0xdf, 0x00, 0x00, 0x00,
}
