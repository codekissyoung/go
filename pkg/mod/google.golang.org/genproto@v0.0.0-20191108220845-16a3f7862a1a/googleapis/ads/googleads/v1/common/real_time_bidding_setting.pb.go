// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/real_time_bidding_setting.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Settings for Real-Time Bidding, a feature only available for campaigns
// targeting the Ad Exchange network.
type RealTimeBiddingSetting struct {
	// Whether the campaign is opted in to real-time bidding.
	OptIn                *wrappers.BoolValue `protobuf:"bytes,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RealTimeBiddingSetting) Reset()         { *m = RealTimeBiddingSetting{} }
func (m *RealTimeBiddingSetting) String() string { return proto.CompactTextString(m) }
func (*RealTimeBiddingSetting) ProtoMessage()    {}
func (*RealTimeBiddingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_7241f0db40dc1f31, []int{0}
}

func (m *RealTimeBiddingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeBiddingSetting.Unmarshal(m, b)
}
func (m *RealTimeBiddingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeBiddingSetting.Marshal(b, m, deterministic)
}
func (m *RealTimeBiddingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeBiddingSetting.Merge(m, src)
}
func (m *RealTimeBiddingSetting) XXX_Size() int {
	return xxx_messageInfo_RealTimeBiddingSetting.Size(m)
}
func (m *RealTimeBiddingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeBiddingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeBiddingSetting proto.InternalMessageInfo

func (m *RealTimeBiddingSetting) GetOptIn() *wrappers.BoolValue {
	if m != nil {
		return m.OptIn
	}
	return nil
}

func init() {
	proto.RegisterType((*RealTimeBiddingSetting)(nil), "google.ads.googleads.v1.common.RealTimeBiddingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/real_time_bidding_setting.proto", fileDescriptor_7241f0db40dc1f31)
}

var fileDescriptor_7241f0db40dc1f31 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0xfc, 0xfc, 0xb3, 0xa8, 0xbb, 0x59, 0x88, 0x8c, 0x32, 0xc8, 0xac, 0x5c, 0x25,
	0x44, 0x77, 0x11, 0x84, 0x56, 0x61, 0x10, 0x37, 0xc3, 0x28, 0x5d, 0x48, 0xa1, 0x64, 0x26, 0x31,
	0x04, 0xd2, 0xdc, 0xd0, 0x64, 0xc6, 0xf7, 0x71, 0xe9, 0xa3, 0xf8, 0x28, 0x3e, 0x80, 0x6b, 0x69,
	0x6f, 0xdb, 0x95, 0xba, 0xca, 0x49, 0x72, 0xbe, 0x73, 0x2e, 0x37, 0xbd, 0xd1, 0x00, 0xda, 0x2a,
	0x2a, 0x64, 0xa0, 0x28, 0x5b, 0x75, 0x60, 0x74, 0x07, 0x75, 0x0d, 0x8e, 0x36, 0x4a, 0xd8, 0x2a,
	0x9a, 0x5a, 0x55, 0x5b, 0x23, 0xa5, 0x71, 0xba, 0x0a, 0x2a, 0x46, 0xe3, 0x34, 0xf1, 0x0d, 0x44,
	0x98, 0x2d, 0x10, 0x22, 0x42, 0x06, 0x32, 0xf2, 0xe4, 0xc0, 0x08, 0xf2, 0xf3, 0xb3, 0x21, 0xdf,
	0x1b, 0x2a, 0x9c, 0x83, 0x28, 0xa2, 0x01, 0x17, 0x90, 0x9e, 0xf7, 0x34, 0xed, 0x6e, 0xdb, 0xfd,
	0x0b, 0x7d, 0x6d, 0x84, 0xf7, 0xaa, 0xe9, 0xff, 0x97, 0x0f, 0xe9, 0xf1, 0x46, 0x09, 0xfb, 0x64,
	0x6a, 0x95, 0x63, 0xfd, 0x23, 0xb6, 0xcf, 0x58, 0x3a, 0x05, 0x1f, 0x2b, 0xe3, 0x4e, 0x92, 0xf3,
	0xe4, 0xe2, 0xe8, 0x72, 0xde, 0xb7, 0x93, 0x21, 0x8a, 0xe4, 0x00, 0xb6, 0x10, 0x76, 0xaf, 0x36,
	0xff, 0xc1, 0xc7, 0x7b, 0x97, 0x7f, 0x25, 0xe9, 0x72, 0x07, 0x35, 0xf9, 0x7b, 0xe2, 0xfc, 0xf4,
	0xe7, 0xc6, 0x75, 0x9b, 0xbb, 0x4e, 0x9e, 0xef, 0x7a, 0x5c, 0x83, 0x15, 0x4e, 0x13, 0x68, 0x34,
	0xd5, 0xca, 0x75, 0xad, 0xc3, 0x02, 0xbd, 0x09, 0xbf, 0xed, 0xf3, 0x1a, 0x8f, 0xb7, 0xc9, 0xbf,
	0x55, 0x96, 0xbd, 0x4f, 0x16, 0x2b, 0x0c, 0xcb, 0x64, 0x20, 0x28, 0x5b, 0x55, 0x30, 0x72, 0xdb,
	0xd9, 0x3e, 0x06, 0x43, 0x99, 0xc9, 0x50, 0x8e, 0x86, 0xb2, 0x60, 0x25, 0x1a, 0x3e, 0x27, 0x4b,
	0x7c, 0xe5, 0x3c, 0x93, 0x81, 0xf3, 0xd1, 0xc2, 0x79, 0xc1, 0x38, 0x47, 0xd3, 0x76, 0xda, 0x4d,
	0x77, 0xf5, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x89, 0xa2, 0x9b, 0x0c, 0xec, 0x01, 0x00, 0x00,
}
