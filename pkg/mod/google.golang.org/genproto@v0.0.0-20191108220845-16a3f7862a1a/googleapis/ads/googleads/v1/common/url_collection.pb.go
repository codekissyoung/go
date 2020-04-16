// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/common/url_collection.proto

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

// Collection of urls that is tagged with a unique identifier.
type UrlCollection struct {
	// Unique identifier for this UrlCollection instance.
	UrlCollectionId *wrappers.StringValue `protobuf:"bytes,1,opt,name=url_collection_id,json=urlCollectionId,proto3" json:"url_collection_id,omitempty"`
	// A list of possible final URLs.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,2,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// A list of possible final mobile URLs.
	FinalMobileUrls []*wrappers.StringValue `protobuf:"bytes,3,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// URL template for constructing a tracking URL.
	TrackingUrlTemplate  *wrappers.StringValue `protobuf:"bytes,4,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UrlCollection) Reset()         { *m = UrlCollection{} }
func (m *UrlCollection) String() string { return proto.CompactTextString(m) }
func (*UrlCollection) ProtoMessage()    {}
func (*UrlCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_82f492aa5cd99b44, []int{0}
}

func (m *UrlCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlCollection.Unmarshal(m, b)
}
func (m *UrlCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlCollection.Marshal(b, m, deterministic)
}
func (m *UrlCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlCollection.Merge(m, src)
}
func (m *UrlCollection) XXX_Size() int {
	return xxx_messageInfo_UrlCollection.Size(m)
}
func (m *UrlCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlCollection.DiscardUnknown(m)
}

var xxx_messageInfo_UrlCollection proto.InternalMessageInfo

func (m *UrlCollection) GetUrlCollectionId() *wrappers.StringValue {
	if m != nil {
		return m.UrlCollectionId
	}
	return nil
}

func (m *UrlCollection) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *UrlCollection) GetFinalMobileUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalMobileUrls
	}
	return nil
}

func (m *UrlCollection) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func init() {
	proto.RegisterType((*UrlCollection)(nil), "google.ads.googleads.v1.common.UrlCollection")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/common/url_collection.proto", fileDescriptor_82f492aa5cd99b44)
}

var fileDescriptor_82f492aa5cd99b44 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x6a, 0xe3, 0x30,
	0x1c, 0xc6, 0xb1, 0x73, 0x1c, 0x9c, 0x8f, 0x23, 0x9c, 0x4b, 0x21, 0x84, 0x10, 0x42, 0xa6, 0x4e,
	0x12, 0x6e, 0x36, 0x65, 0x72, 0x52, 0x48, 0x3b, 0x14, 0x42, 0xdb, 0x78, 0x28, 0x06, 0xa3, 0xd8,
	0x8a, 0x10, 0x95, 0x25, 0x23, 0xd9, 0xe9, 0xfb, 0x74, 0xec, 0xd8, 0xc7, 0xe8, 0xa3, 0x14, 0xfa,
	0x0e, 0xc5, 0x92, 0x6d, 0x9a, 0xa1, 0x6d, 0x26, 0xff, 0x6d, 0x7f, 0xbf, 0x9f, 0x3e, 0x5b, 0xf2,
	0x66, 0x54, 0x4a, 0xca, 0x09, 0xc4, 0x99, 0x86, 0x76, 0xac, 0xa7, 0x7d, 0x00, 0x53, 0x99, 0xe7,
	0x52, 0xc0, 0x4a, 0xf1, 0x24, 0x95, 0x9c, 0x93, 0xb4, 0x64, 0x52, 0x80, 0x42, 0xc9, 0x52, 0xfa,
	0x63, 0x9b, 0x04, 0x38, 0xd3, 0xa0, 0x83, 0xc0, 0x3e, 0x00, 0x16, 0x1a, 0x8e, 0x5a, 0x69, 0xc1,
	0x20, 0x16, 0x42, 0x96, 0xb8, 0x86, 0xb5, 0xa5, 0x87, 0x0d, 0x0d, 0xcd, 0xdd, 0xb6, 0xda, 0xc1,
	0x47, 0x85, 0x8b, 0x82, 0xa8, 0xe6, 0xfd, 0xf4, 0xc5, 0xf5, 0xfe, 0x6d, 0x14, 0x5f, 0x76, 0xab,
	0xfa, 0x97, 0xde, 0xff, 0xc3, 0x1e, 0x09, 0xcb, 0x06, 0xce, 0xc4, 0x39, 0xfb, 0x7b, 0x3e, 0x6a,
	0x0a, 0x80, 0xd6, 0x06, 0x6e, 0x4b, 0xc5, 0x04, 0x8d, 0x30, 0xaf, 0xc8, 0x4d, 0xbf, 0xfa, 0xec,
	0xb9, 0xca, 0xfc, 0xb9, 0xe7, 0xed, 0x98, 0xc0, 0x3c, 0xa9, 0x14, 0xd7, 0x03, 0x77, 0xd2, 0xfb,
	0x51, 0xf1, 0xc7, 0xe4, 0x37, 0x8a, 0xeb, 0xba, 0x86, 0x85, 0x73, 0xb9, 0x65, 0x9c, 0x58, 0x47,
	0xef, 0x08, 0x47, 0xdf, 0x60, 0xd7, 0x86, 0x32, 0xa6, 0xb5, 0x77, 0x5a, 0x2a, 0x9c, 0x3e, 0x30,
	0x41, 0x6b, 0x4b, 0x52, 0x92, 0xbc, 0xe0, 0xb8, 0x24, 0x83, 0x5f, 0x47, 0x7c, 0xd4, 0x49, 0x8b,
	0x6e, 0x14, 0xbf, 0x6b, 0xc0, 0xc5, 0xbb, 0xe3, 0x4d, 0x53, 0x99, 0x83, 0xef, 0x77, 0x66, 0xe1,
	0x1f, 0xfc, 0xd8, 0x75, 0xad, 0x5f, 0x3b, 0xf7, 0x17, 0x0d, 0x45, 0x25, 0xc7, 0x82, 0x02, 0xa9,
	0x28, 0xa4, 0x44, 0x98, 0xc5, 0xdb, 0x43, 0x51, 0x30, 0xfd, 0xd5, 0x19, 0x99, 0xdb, 0xcb, 0x93,
	0xdb, 0x5b, 0x85, 0xe1, 0xb3, 0x3b, 0x5e, 0x59, 0x59, 0x98, 0x69, 0x60, 0xc7, 0x7a, 0x8a, 0x02,
	0xb0, 0x34, 0xb1, 0xd7, 0x36, 0x10, 0x87, 0x99, 0x8e, 0xbb, 0x40, 0x1c, 0x05, 0xb1, 0x0d, 0xbc,
	0xb9, 0x53, 0xfb, 0x14, 0xa1, 0x30, 0xd3, 0x08, 0x75, 0x11, 0x84, 0xa2, 0x00, 0x21, 0x1b, 0xda,
	0xfe, 0x36, 0xed, 0x66, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x53, 0x0d, 0xd6, 0xc2, 0xc0, 0x02,
	0x00, 0x00,
}
