// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/setting_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible setting errors.
type SettingErrorEnum_SettingError int32

const (
	// Enum unspecified.
	SettingErrorEnum_UNSPECIFIED SettingErrorEnum_SettingError = 0
	// The received error code is not known in this version.
	SettingErrorEnum_UNKNOWN SettingErrorEnum_SettingError = 1
	// The campaign setting is not available for this Google Ads account.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_AVAILABLE SettingErrorEnum_SettingError = 3
	// The setting is not compatible with the campaign.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN SettingErrorEnum_SettingError = 4
	// The supplied TargetingSetting contains an invalid CriterionTypeGroup. See
	// CriterionTypeGroup documentation for CriterionTypeGroups allowed
	// in Campaign or AdGroup TargetingSettings.
	SettingErrorEnum_TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 5
	// TargetingSetting must not explicitly
	// set any of the Demographic CriterionTypeGroups (AGE_RANGE, GENDER,
	// PARENT, INCOME_RANGE) to false (it's okay to not set them at all, in
	// which case the system will set them to true automatically).
	SettingErrorEnum_TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL SettingErrorEnum_SettingError = 6
	// TargetingSetting cannot change any of
	// the Demographic CriterionTypeGroups (AGE_RANGE, GENDER, PARENT,
	// INCOME_RANGE) from true to false.
	SettingErrorEnum_TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 7
	// At least one feed id should be present.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT SettingErrorEnum_SettingError = 8
	// The supplied DynamicSearchAdsSetting contains an invalid domain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME SettingErrorEnum_SettingError = 9
	// The supplied DynamicSearchAdsSetting contains a subdomain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME SettingErrorEnum_SettingError = 10
	// The supplied DynamicSearchAdsSetting contains an invalid language code.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE SettingErrorEnum_SettingError = 11
	// TargetingSettings in search campaigns should not have
	// CriterionTypeGroup.PLACEMENT set to targetAll.
	SettingErrorEnum_TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN SettingErrorEnum_SettingError = 12
	// Duplicate description in universal app setting description field.
	SettingErrorEnum_UNIVERSAL_APP_CAMPAIGN_SETTING_DUPLICATE_DESCRIPTION SettingErrorEnum_SettingError = 13
	// Description line width is too long in universal app setting description
	// field.
	SettingErrorEnum_UNIVERSAL_APP_CAMPAIGN_SETTING_DESCRIPTION_LINE_WIDTH_TOO_LONG SettingErrorEnum_SettingError = 14
	// Universal app setting appId field cannot be modified for COMPLETE
	// campaigns.
	SettingErrorEnum_UNIVERSAL_APP_CAMPAIGN_SETTING_APP_ID_CANNOT_BE_MODIFIED SettingErrorEnum_SettingError = 15
	// YoutubeVideoMediaIds in universal app setting cannot exceed size limit.
	SettingErrorEnum_TOO_MANY_YOUTUBE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN SettingErrorEnum_SettingError = 16
	// ImageMediaIds in universal app setting cannot exceed size limit.
	SettingErrorEnum_TOO_MANY_IMAGE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN SettingErrorEnum_SettingError = 17
	// Media is incompatible for universal app campaign.
	SettingErrorEnum_MEDIA_INCOMPATIBLE_FOR_UNIVERSAL_APP_CAMPAIGN SettingErrorEnum_SettingError = 18
	// Too many exclamation marks in universal app campaign ad text ideas.
	SettingErrorEnum_TOO_MANY_EXCLAMATION_MARKS SettingErrorEnum_SettingError = 19
)

var SettingErrorEnum_SettingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "SETTING_TYPE_IS_NOT_AVAILABLE",
	4:  "SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN",
	5:  "TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP",
	6:  "TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL",
	7:  "TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP",
	8:  "DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT",
	9:  "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME",
	10: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME",
	11: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE",
	12: "TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN",
	13: "UNIVERSAL_APP_CAMPAIGN_SETTING_DUPLICATE_DESCRIPTION",
	14: "UNIVERSAL_APP_CAMPAIGN_SETTING_DESCRIPTION_LINE_WIDTH_TOO_LONG",
	15: "UNIVERSAL_APP_CAMPAIGN_SETTING_APP_ID_CANNOT_BE_MODIFIED",
	16: "TOO_MANY_YOUTUBE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN",
	17: "TOO_MANY_IMAGE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN",
	18: "MEDIA_INCOMPATIBLE_FOR_UNIVERSAL_APP_CAMPAIGN",
	19: "TOO_MANY_EXCLAMATION_MARKS",
}
var SettingErrorEnum_SettingError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"SETTING_TYPE_IS_NOT_AVAILABLE": 3,
	"SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN":                                             4,
	"TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP":                                  5,
	"TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL":            6,
	"TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP": 7,
	"DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT":                          8,
	"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME":                                  9,
	"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME":                                       10,
	"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE":                                11,
	"TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN":                               12,
	"UNIVERSAL_APP_CAMPAIGN_SETTING_DUPLICATE_DESCRIPTION":                                     13,
	"UNIVERSAL_APP_CAMPAIGN_SETTING_DESCRIPTION_LINE_WIDTH_TOO_LONG":                           14,
	"UNIVERSAL_APP_CAMPAIGN_SETTING_APP_ID_CANNOT_BE_MODIFIED":                                 15,
	"TOO_MANY_YOUTUBE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN":                                     16,
	"TOO_MANY_IMAGE_MEDIA_IDS_IN_UNIVERSAL_APP_CAMPAIGN":                                       17,
	"MEDIA_INCOMPATIBLE_FOR_UNIVERSAL_APP_CAMPAIGN":                                            18,
	"TOO_MANY_EXCLAMATION_MARKS":                                                               19,
}

func (x SettingErrorEnum_SettingError) String() string {
	return proto.EnumName(SettingErrorEnum_SettingError_name, int32(x))
}
func (SettingErrorEnum_SettingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_setting_error_284274bcad413424, []int{0, 0}
}

// Container for enum describing possible setting errors.
type SettingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingErrorEnum) Reset()         { *m = SettingErrorEnum{} }
func (m *SettingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*SettingErrorEnum) ProtoMessage()    {}
func (*SettingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_setting_error_284274bcad413424, []int{0}
}
func (m *SettingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingErrorEnum.Unmarshal(m, b)
}
func (m *SettingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingErrorEnum.Marshal(b, m, deterministic)
}
func (dst *SettingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingErrorEnum.Merge(dst, src)
}
func (m *SettingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_SettingErrorEnum.Size(m)
}
func (m *SettingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_SettingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SettingErrorEnum)(nil), "google.ads.googleads.v0.errors.SettingErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.SettingErrorEnum_SettingError", SettingErrorEnum_SettingError_name, SettingErrorEnum_SettingError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/setting_error.proto", fileDescriptor_setting_error_284274bcad413424)
}

var fileDescriptor_setting_error_284274bcad413424 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x6f, 0x8f, 0xdb, 0x44,
	0x10, 0xc6, 0xb9, 0x2b, 0xb4, 0xc7, 0x5e, 0xa1, 0xdb, 0xe5, 0x1d, 0x12, 0x27, 0x91, 0xd7, 0xe0,
	0x84, 0x82, 0xa0, 0xb8, 0x50, 0x34, 0xf6, 0x4e, 0x9c, 0x55, 0xd7, 0xbb, 0xc6, 0xbb, 0xce, 0x35,
	0x28, 0xd2, 0xe8, 0x20, 0x51, 0x54, 0xa9, 0x8d, 0xab, 0xf8, 0xe8, 0x07, 0xe2, 0x1d, 0x7c, 0x08,
	0x3e, 0x00, 0x1f, 0x05, 0xf1, 0x21, 0xd0, 0xda, 0x8d, 0x2f, 0xd2, 0xe5, 0xfe, 0xf4, 0x55, 0x36,
	0xe3, 0xe7, 0x37, 0x33, 0x7e, 0x3c, 0x3b, 0xec, 0xd1, 0xaa, 0xae, 0x57, 0x2f, 0x97, 0xc3, 0xb3,
	0x45, 0x33, 0xec, 0x8e, 0xe1, 0xf4, 0x66, 0x34, 0x5c, 0x6e, 0x36, 0xf5, 0xa6, 0x19, 0x36, 0xcb,
	0xf3, 0xf3, 0x17, 0xeb, 0x15, 0xb5, 0x7f, 0xa3, 0xd7, 0x9b, 0xfa, 0xbc, 0x16, 0x27, 0x9d, 0x30,
	0x3a, 0x5b, 0x34, 0x51, 0xcf, 0x44, 0x6f, 0x46, 0x51, 0xc7, 0x0c, 0xfe, 0x3e, 0x62, 0xdc, 0x75,
	0x1c, 0x86, 0x08, 0xae, 0x7f, 0x7f, 0x35, 0xf8, 0xf3, 0x88, 0xdd, 0xdf, 0x0d, 0x8a, 0x07, 0xec,
	0xb8, 0x32, 0xae, 0xc0, 0x54, 0x8d, 0x15, 0x4a, 0xfe, 0x9e, 0x38, 0x66, 0xf7, 0x2a, 0xf3, 0xcc,
	0xd8, 0x53, 0xc3, 0x0f, 0xc4, 0xe7, 0xec, 0x33, 0x87, 0xde, 0x2b, 0x93, 0x91, 0x9f, 0x15, 0x48,
	0xca, 0x91, 0xb1, 0x9e, 0x60, 0x0a, 0x4a, 0x43, 0xa2, 0x91, 0xdf, 0x11, 0x23, 0xf6, 0xc5, 0x3e,
	0x49, 0x6a, 0xf3, 0x02, 0xbc, 0x4a, 0x34, 0xd2, 0xa9, 0xf2, 0x13, 0x4a, 0x21, 0x2f, 0x40, 0x65,
	0x86, 0xbf, 0x2f, 0x9e, 0xb0, 0xef, 0x3c, 0x94, 0x19, 0xb6, 0xcc, 0x96, 0x4d, 0xad, 0xf1, 0xa0,
	0x8c, 0x23, 0x65, 0xa6, 0xa0, 0x95, 0xa4, 0xb4, 0x54, 0x1e, 0x4b, 0x65, 0x4d, 0x97, 0x36, 0x2b,
	0x6d, 0x55, 0xf0, 0x0f, 0xc4, 0xcf, 0x2c, 0xbf, 0x0c, 0x4b, 0xcc, 0x6d, 0x56, 0x42, 0x31, 0x51,
	0xe9, 0x5e, 0xce, 0x51, 0x5e, 0x39, 0x4f, 0x09, 0x06, 0x82, 0xbc, 0xa5, 0x2e, 0x05, 0x81, 0xd6,
	0xfc, 0xae, 0x98, 0xb3, 0xe7, 0x7b, 0xfa, 0x01, 0xd3, 0xbe, 0xc6, 0x04, 0x4c, 0x86, 0x3b, 0xfa,
	0x40, 0x8f, 0x41, 0x3b, 0xa4, 0xb1, 0x2d, 0x6f, 0x2c, 0xcc, 0xef, 0x89, 0x94, 0xfd, 0x24, 0x67,
	0x06, 0x72, 0x95, 0x92, 0x43, 0x28, 0xd3, 0x09, 0x81, 0x74, 0x7d, 0x19, 0xf0, 0xa4, 0x11, 0x9c,
	0x27, 0x6b, 0x90, 0xc6, 0x88, 0x92, 0x94, 0xec, 0x9b, 0x2d, 0x4a, 0x74, 0x68, 0x3c, 0x3f, 0x0a,
	0x96, 0x5d, 0x93, 0xe4, 0x92, 0x77, 0xd2, 0xe6, 0xa0, 0x0c, 0x19, 0xc8, 0x91, 0x7f, 0x28, 0xbe,
	0x65, 0x8f, 0x6e, 0x03, 0xbb, 0x2a, 0xd9, 0xe5, 0x98, 0xf8, 0x91, 0x7d, 0xff, 0x2e, 0x45, 0x35,
	0x98, 0xac, 0x82, 0x0c, 0x29, 0xb5, 0x12, 0xf9, 0xb1, 0x78, 0xca, 0xe2, 0x1d, 0xdb, 0xb6, 0x93,
	0xa3, 0xb5, 0x3d, 0x45, 0xd9, 0x9a, 0x57, 0x68, 0x48, 0x31, 0x47, 0xe3, 0x49, 0x99, 0x6d, 0x85,
	0x7e, 0x4c, 0xee, 0x8b, 0xc7, 0xec, 0x9b, 0xca, 0xa8, 0x29, 0x96, 0x0e, 0x34, 0x41, 0x51, 0xf4,
	0xcf, 0x2e, 0x3e, 0x7b, 0x55, 0x68, 0x95, 0x82, 0x47, 0x92, 0xe8, 0xd2, 0x52, 0x15, 0x5e, 0x59,
	0xc3, 0x3f, 0x12, 0x09, 0x7b, 0x7a, 0x13, 0x79, 0xa1, 0x27, 0xad, 0x4c, 0x98, 0x51, 0xe9, 0x27,
	0xe4, 0xad, 0x25, 0x6d, 0x4d, 0xc6, 0x3f, 0x16, 0x3f, 0xb0, 0xc7, 0x37, 0xe4, 0x08, 0xc1, 0x30,
	0xa7, 0xdd, 0xa0, 0x24, 0x48, 0xb9, 0x95, 0xdd, 0x25, 0x7a, 0x10, 0x7a, 0x0f, 0xb9, 0x72, 0x30,
	0x33, 0x9a, 0xd9, 0xca, 0x57, 0xe1, 0x31, 0x4a, 0x05, 0xa4, 0x64, 0x70, 0x8c, 0xf6, 0xe7, 0xe6,
	0x3c, 0x7c, 0xac, 0x9e, 0x54, 0x79, 0xb0, 0xf3, 0x36, 0xdc, 0x43, 0xf1, 0x15, 0xfb, 0xf2, 0xad,
	0xd0, 0xec, 0x5c, 0xbd, 0x60, 0xf3, 0x15, 0x88, 0x10, 0x27, 0xec, 0xd3, 0xbe, 0x14, 0x3e, 0x4f,
	0x35, 0xe4, 0xd0, 0x1a, 0x92, 0x43, 0xf9, 0xcc, 0xf1, 0x4f, 0x92, 0xff, 0x0e, 0xd8, 0xe0, 0xb7,
	0xfa, 0x55, 0x74, 0xfd, 0x9e, 0x49, 0x1e, 0xee, 0xee, 0x93, 0x22, 0xac, 0xa6, 0xe2, 0xe0, 0x17,
	0xf9, 0x16, 0x5a, 0xd5, 0x2f, 0xcf, 0xd6, 0xab, 0xa8, 0xde, 0xac, 0x86, 0xab, 0xe5, 0xba, 0x5d,
	0x5c, 0xdb, 0x05, 0xf7, 0xfa, 0x45, 0x73, 0xd5, 0xbe, 0x7b, 0xd2, 0xfd, 0xfc, 0x71, 0x78, 0x27,
	0x03, 0xf8, 0xeb, 0xf0, 0x24, 0xeb, 0x92, 0xc1, 0xa2, 0x89, 0xba, 0x63, 0x38, 0x4d, 0x47, 0x51,
	0x5b, 0xb2, 0xf9, 0x67, 0x2b, 0x98, 0xc3, 0xa2, 0x99, 0xf7, 0x82, 0xf9, 0x74, 0x34, 0xef, 0x04,
	0xff, 0x1e, 0x0e, 0xba, 0x68, 0x1c, 0xc3, 0xa2, 0x89, 0xe3, 0x5e, 0x12, 0xc7, 0xd3, 0x51, 0x1c,
	0x77, 0xa2, 0x5f, 0xef, 0xb6, 0xdd, 0x7d, 0xfd, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb2,
	0xb5, 0x7b, 0x8c, 0x05, 0x00, 0x00,
}
