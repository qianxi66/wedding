// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/policy_approval_status.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

// The possible policy approval statuses. When there are several approval
// statuses available the most severe one will be used. The order of severity
// is DISAPPROVED, AREA_OF_INTEREST_ONLY, APPROVED_LIMITED and APPROVED.
type PolicyApprovalStatusEnum_PolicyApprovalStatus int32

const (
	// No value has been specified.
	PolicyApprovalStatusEnum_UNSPECIFIED PolicyApprovalStatusEnum_PolicyApprovalStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyApprovalStatusEnum_UNKNOWN PolicyApprovalStatusEnum_PolicyApprovalStatus = 1
	// Will not serve.
	PolicyApprovalStatusEnum_DISAPPROVED PolicyApprovalStatusEnum_PolicyApprovalStatus = 2
	// Serves with restrictions.
	PolicyApprovalStatusEnum_APPROVED_LIMITED PolicyApprovalStatusEnum_PolicyApprovalStatus = 3
	// Serves without restrictions.
	PolicyApprovalStatusEnum_APPROVED PolicyApprovalStatusEnum_PolicyApprovalStatus = 4
	// Will not serve in targeted countries, but may serve for users who are
	// searching for information about the targeted countries.
	PolicyApprovalStatusEnum_AREA_OF_INTEREST_ONLY PolicyApprovalStatusEnum_PolicyApprovalStatus = 5
)

var PolicyApprovalStatusEnum_PolicyApprovalStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "DISAPPROVED",
	3: "APPROVED_LIMITED",
	4: "APPROVED",
	5: "AREA_OF_INTEREST_ONLY",
}

var PolicyApprovalStatusEnum_PolicyApprovalStatus_value = map[string]int32{
	"UNSPECIFIED":           0,
	"UNKNOWN":               1,
	"DISAPPROVED":           2,
	"APPROVED_LIMITED":      3,
	"APPROVED":              4,
	"AREA_OF_INTEREST_ONLY": 5,
}

func (x PolicyApprovalStatusEnum_PolicyApprovalStatus) String() string {
	return proto.EnumName(PolicyApprovalStatusEnum_PolicyApprovalStatus_name, int32(x))
}

func (PolicyApprovalStatusEnum_PolicyApprovalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ecf4cfc6f98f313e, []int{0, 0}
}

// Container for enum describing possible policy approval statuses.
type PolicyApprovalStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyApprovalStatusEnum) Reset()         { *m = PolicyApprovalStatusEnum{} }
func (m *PolicyApprovalStatusEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyApprovalStatusEnum) ProtoMessage()    {}
func (*PolicyApprovalStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecf4cfc6f98f313e, []int{0}
}

func (m *PolicyApprovalStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyApprovalStatusEnum.Unmarshal(m, b)
}
func (m *PolicyApprovalStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyApprovalStatusEnum.Marshal(b, m, deterministic)
}
func (m *PolicyApprovalStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyApprovalStatusEnum.Merge(m, src)
}
func (m *PolicyApprovalStatusEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyApprovalStatusEnum.Size(m)
}
func (m *PolicyApprovalStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyApprovalStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyApprovalStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.PolicyApprovalStatusEnum_PolicyApprovalStatus", PolicyApprovalStatusEnum_PolicyApprovalStatus_name, PolicyApprovalStatusEnum_PolicyApprovalStatus_value)
	proto.RegisterType((*PolicyApprovalStatusEnum)(nil), "google.ads.googleads.v2.enums.PolicyApprovalStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/policy_approval_status.proto", fileDescriptor_ecf4cfc6f98f313e)
}

var fileDescriptor_ecf4cfc6f98f313e = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4d, 0x4a, 0xc3, 0x40,
	0x14, 0x36, 0xa9, 0x7f, 0x4c, 0x05, 0x43, 0xa8, 0x60, 0xc5, 0x2e, 0xda, 0x03, 0x4c, 0x20, 0xee,
	0xc6, 0xd5, 0xd4, 0x4c, 0x4b, 0xb0, 0x26, 0xa1, 0x69, 0x23, 0x4a, 0x20, 0x8c, 0x4d, 0x09, 0x81,
	0x74, 0x26, 0x74, 0xd2, 0x82, 0x7b, 0x2f, 0xa2, 0x4b, 0x8f, 0xe2, 0x51, 0x5c, 0x78, 0x06, 0xc9,
	0xa4, 0xc9, 0xaa, 0xba, 0x19, 0xbe, 0x79, 0xdf, 0x0f, 0xef, 0x7d, 0x00, 0x25, 0x9c, 0x27, 0xd9,
	0xd2, 0xa0, 0xb1, 0x30, 0x2a, 0x58, 0xa2, 0xad, 0x69, 0x2c, 0xd9, 0x66, 0x25, 0x8c, 0x9c, 0x67,
	0xe9, 0xe2, 0x35, 0xa2, 0x79, 0xbe, 0xe6, 0x5b, 0x9a, 0x45, 0xa2, 0xa0, 0xc5, 0x46, 0xc0, 0x7c,
	0xcd, 0x0b, 0xae, 0xf7, 0x2a, 0x03, 0xa4, 0xb1, 0x80, 0x8d, 0x17, 0x6e, 0x4d, 0x28, 0xbd, 0x57,
	0xd7, 0x75, 0x74, 0x9e, 0x1a, 0x94, 0x31, 0x5e, 0xd0, 0x22, 0xe5, 0x6c, 0x67, 0x1e, 0xbc, 0x2b,
	0xe0, 0xd2, 0x93, 0xe9, 0x78, 0x17, 0xee, 0xcb, 0x6c, 0xc2, 0x36, 0xab, 0xc1, 0x9b, 0x02, 0x3a,
	0xfb, 0x48, 0xfd, 0x1c, 0xb4, 0xe7, 0x8e, 0xef, 0x91, 0x3b, 0x7b, 0x64, 0x13, 0x4b, 0x3b, 0xd0,
	0xdb, 0xe0, 0x64, 0xee, 0xdc, 0x3b, 0xee, 0xa3, 0xa3, 0x29, 0x25, 0x6b, 0xd9, 0x3e, 0xf6, 0xbc,
	0xa9, 0x1b, 0x10, 0x4b, 0x53, 0xf5, 0x0e, 0xd0, 0xea, 0x5f, 0x34, 0xb1, 0x1f, 0xec, 0x19, 0xb1,
	0xb4, 0x96, 0x7e, 0x06, 0x4e, 0x1b, 0xcd, 0xa1, 0xde, 0x05, 0x17, 0x78, 0x4a, 0x70, 0xe4, 0x8e,
	0x22, 0xdb, 0x99, 0x91, 0x29, 0xf1, 0x67, 0x91, 0xeb, 0x4c, 0x9e, 0xb4, 0xa3, 0xe1, 0x8f, 0x02,
	0xfa, 0x0b, 0xbe, 0x82, 0xff, 0xde, 0x39, 0xec, 0xee, 0xdb, 0xd4, 0x2b, 0x8f, 0xf4, 0x94, 0xe7,
	0xe1, 0xce, 0x9b, 0xf0, 0x8c, 0xb2, 0x04, 0xf2, 0x75, 0x62, 0x24, 0x4b, 0x26, 0x2b, 0xa8, 0xfb,
	0xce, 0x53, 0xf1, 0x47, 0xfd, 0xb7, 0xf2, 0xfd, 0x50, 0x5b, 0x63, 0x8c, 0x3f, 0xd5, 0xde, 0xb8,
	0x8a, 0xc2, 0xb1, 0x80, 0x15, 0x2c, 0x51, 0x60, 0xc2, 0xb2, 0x32, 0xf1, 0x55, 0xf3, 0x21, 0x8e,
	0x45, 0xd8, 0xf0, 0x61, 0x60, 0x86, 0x92, 0xff, 0x56, 0xfb, 0xd5, 0x10, 0x21, 0x1c, 0x0b, 0x84,
	0x1a, 0x05, 0x42, 0x81, 0x89, 0x90, 0xd4, 0xbc, 0x1c, 0xcb, 0xc5, 0x6e, 0x7e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xa5, 0xf8, 0xc3, 0x8a, 0x16, 0x02, 0x00, 0x00,
}
