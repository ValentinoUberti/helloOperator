// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/location_placeholder_field.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Possible values for Location placeholder fields.
type LocationPlaceholderFieldEnum_LocationPlaceholderField int32

const (
	// Not specified.
	LocationPlaceholderFieldEnum_UNSPECIFIED LocationPlaceholderFieldEnum_LocationPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	LocationPlaceholderFieldEnum_UNKNOWN LocationPlaceholderFieldEnum_LocationPlaceholderField = 1
	// Data Type: STRING. The name of the business.
	LocationPlaceholderFieldEnum_BUSINESS_NAME LocationPlaceholderFieldEnum_LocationPlaceholderField = 2
	// Data Type: STRING. Line 1 of the business address.
	LocationPlaceholderFieldEnum_ADDRESS_LINE_1 LocationPlaceholderFieldEnum_LocationPlaceholderField = 3
	// Data Type: STRING. Line 2 of the business address.
	LocationPlaceholderFieldEnum_ADDRESS_LINE_2 LocationPlaceholderFieldEnum_LocationPlaceholderField = 4
	// Data Type: STRING. City of the business address.
	LocationPlaceholderFieldEnum_CITY LocationPlaceholderFieldEnum_LocationPlaceholderField = 5
	// Data Type: STRING. Province of the business address.
	LocationPlaceholderFieldEnum_PROVINCE LocationPlaceholderFieldEnum_LocationPlaceholderField = 6
	// Data Type: STRING. Postal code of the business address.
	LocationPlaceholderFieldEnum_POSTAL_CODE LocationPlaceholderFieldEnum_LocationPlaceholderField = 7
	// Data Type: STRING. Country code of the business address.
	LocationPlaceholderFieldEnum_COUNTRY_CODE LocationPlaceholderFieldEnum_LocationPlaceholderField = 8
	// Data Type: STRING. Phone number of the business.
	LocationPlaceholderFieldEnum_PHONE_NUMBER LocationPlaceholderFieldEnum_LocationPlaceholderField = 9
)

var LocationPlaceholderFieldEnum_LocationPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BUSINESS_NAME",
	3: "ADDRESS_LINE_1",
	4: "ADDRESS_LINE_2",
	5: "CITY",
	6: "PROVINCE",
	7: "POSTAL_CODE",
	8: "COUNTRY_CODE",
	9: "PHONE_NUMBER",
}
var LocationPlaceholderFieldEnum_LocationPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"BUSINESS_NAME":  2,
	"ADDRESS_LINE_1": 3,
	"ADDRESS_LINE_2": 4,
	"CITY":           5,
	"PROVINCE":       6,
	"POSTAL_CODE":    7,
	"COUNTRY_CODE":   8,
	"PHONE_NUMBER":   9,
}

func (x LocationPlaceholderFieldEnum_LocationPlaceholderField) String() string {
	return proto.EnumName(LocationPlaceholderFieldEnum_LocationPlaceholderField_name, int32(x))
}
func (LocationPlaceholderFieldEnum_LocationPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_location_placeholder_field_385b0c1f5b48bf69, []int{0, 0}
}

// Values for Location placeholder fields.
type LocationPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocationPlaceholderFieldEnum) Reset()         { *m = LocationPlaceholderFieldEnum{} }
func (m *LocationPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*LocationPlaceholderFieldEnum) ProtoMessage()    {}
func (*LocationPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_location_placeholder_field_385b0c1f5b48bf69, []int{0}
}
func (m *LocationPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocationPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *LocationPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocationPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (dst *LocationPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocationPlaceholderFieldEnum.Merge(dst, src)
}
func (m *LocationPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_LocationPlaceholderFieldEnum.Size(m)
}
func (m *LocationPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_LocationPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_LocationPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LocationPlaceholderFieldEnum)(nil), "google.ads.googleads.v1.enums.LocationPlaceholderFieldEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.LocationPlaceholderFieldEnum_LocationPlaceholderField", LocationPlaceholderFieldEnum_LocationPlaceholderField_name, LocationPlaceholderFieldEnum_LocationPlaceholderField_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/location_placeholder_field.proto", fileDescriptor_location_placeholder_field_385b0c1f5b48bf69)
}

var fileDescriptor_location_placeholder_field_385b0c1f5b48bf69 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0xae, 0x93, 0x40,
	0x14, 0x15, 0xde, 0xf3, 0xbd, 0x3a, 0xef, 0xa9, 0xe3, 0xac, 0x8c, 0x69, 0x17, 0xed, 0x07, 0x0c,
	0x41, 0x77, 0x63, 0x62, 0x32, 0xc0, 0xb4, 0x12, 0xdb, 0x81, 0x40, 0xc1, 0xd4, 0x90, 0x10, 0x2c,
	0x88, 0x24, 0x94, 0x21, 0x9d, 0xb6, 0x1f, 0xe4, 0xd2, 0xaf, 0x70, 0xed, 0x5f, 0xb8, 0x75, 0xe9,
	0x17, 0x98, 0x81, 0xb6, 0x2e, 0x4c, 0xdf, 0x86, 0x9c, 0x9c, 0x7b, 0xcf, 0x39, 0xcc, 0xb9, 0xe0,
	0x5d, 0x29, 0x44, 0x59, 0x17, 0x46, 0x96, 0x4b, 0xa3, 0x87, 0x0a, 0x1d, 0x4c, 0xa3, 0x68, 0xf6,
	0x1b, 0x69, 0xd4, 0x62, 0x9d, 0xed, 0x2a, 0xd1, 0xa4, 0x6d, 0x9d, 0xad, 0x8b, 0xaf, 0xa2, 0xce,
	0x8b, 0x6d, 0xfa, 0xa5, 0x2a, 0xea, 0x1c, 0xb7, 0x5b, 0xb1, 0x13, 0x68, 0xd4, 0x8b, 0x70, 0x96,
	0x4b, 0x7c, 0xd6, 0xe3, 0x83, 0x89, 0x3b, 0xfd, 0xab, 0xe1, 0xc9, 0xbe, 0xad, 0x8c, 0xac, 0x69,
	0xc4, 0xae, 0x73, 0x93, 0xbd, 0x78, 0xf2, 0x4b, 0x03, 0xc3, 0xf9, 0x31, 0xc1, 0xff, 0x17, 0x30,
	0x55, 0xfe, 0xac, 0xd9, 0x6f, 0x26, 0x3f, 0x34, 0xf0, 0xf2, 0xd2, 0x02, 0x7a, 0x0e, 0xee, 0x22,
	0x1e, 0xfa, 0xcc, 0x76, 0xa7, 0x2e, 0x73, 0xe0, 0x23, 0x74, 0x07, 0x6e, 0x23, 0xfe, 0x81, 0x7b,
	0x1f, 0x39, 0xd4, 0xd0, 0x0b, 0xf0, 0xd4, 0x8a, 0x42, 0x97, 0xb3, 0x30, 0x4c, 0x39, 0x5d, 0x30,
	0xa8, 0x23, 0x04, 0x9e, 0x51, 0xc7, 0x09, 0x14, 0x33, 0x77, 0x39, 0x4b, 0x4d, 0x78, 0xf5, 0x1f,
	0xf7, 0x1a, 0x5e, 0xa3, 0x01, 0xb8, 0xb6, 0xdd, 0xe5, 0x0a, 0x3e, 0x46, 0xf7, 0x60, 0xe0, 0x07,
	0x5e, 0xec, 0x72, 0x9b, 0xc1, 0x1b, 0x15, 0xe8, 0x7b, 0xe1, 0x92, 0xce, 0x53, 0xdb, 0x73, 0x18,
	0xbc, 0x45, 0x10, 0xdc, 0xdb, 0x5e, 0xc4, 0x97, 0xc1, 0xaa, 0x67, 0x06, 0x8a, 0xf1, 0xdf, 0x7b,
	0x9c, 0xa5, 0x3c, 0x5a, 0x58, 0x2c, 0x80, 0x4f, 0xac, 0x3f, 0x1a, 0x18, 0xaf, 0xc5, 0x06, 0x3f,
	0xd8, 0x93, 0x35, 0xba, 0xf4, 0x4a, 0x5f, 0x15, 0xe5, 0x6b, 0x9f, 0xac, 0xa3, 0xbe, 0x14, 0x75,
	0xd6, 0x94, 0x58, 0x6c, 0x4b, 0xa3, 0x2c, 0x9a, 0xae, 0xc6, 0xd3, 0xdd, 0xda, 0x4a, 0x5e, 0x38,
	0xe3, 0xdb, 0xee, 0xfb, 0x4d, 0xbf, 0x9a, 0x51, 0xfa, 0x5d, 0x1f, 0xcd, 0x7a, 0x2b, 0x9a, 0x4b,
	0xdc, 0x43, 0x85, 0x62, 0x13, 0xab, 0xca, 0xe5, 0xcf, 0xd3, 0x3c, 0xa1, 0xb9, 0x4c, 0xce, 0xf3,
	0x24, 0x36, 0x93, 0x6e, 0xfe, 0x5b, 0x1f, 0xf7, 0x24, 0x21, 0x34, 0x97, 0x84, 0x9c, 0x37, 0x08,
	0x89, 0x4d, 0x42, 0xba, 0x9d, 0xcf, 0x37, 0xdd, 0x8f, 0xbd, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0x63, 0x67, 0x22, 0x40, 0x5e, 0x02, 0x00, 0x00,
}
