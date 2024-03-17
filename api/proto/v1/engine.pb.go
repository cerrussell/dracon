// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/v1/engine.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Scan related information, unique and immutable per scan run
// This message is copied from LaunchToolRequest to LaunchToolResponse
// by each producer wrapper
type ScanInfo struct {
	// scan unique identifier
	ScanUuid string `protobuf:"bytes,1,opt,name=scan_uuid,json=scanUuid,proto3" json:"scan_uuid,omitempty"`
	// timestamp of when the scan was triggered (passed to LaunchToolResponse)
	ScanStartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=scan_start_time,json=scanStartTime,proto3" json:"scan_start_time,omitempty"`
	// [Optional] scan tags is a user defined list of tags for this scan
	ScanTags             map[string]string `protobuf:"bytes,3,rep,name=scan_tags,json=scanTags,proto3" json:"scan_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ScanInfo) Reset()         { *m = ScanInfo{} }
func (m *ScanInfo) String() string { return proto.CompactTextString(m) }
func (*ScanInfo) ProtoMessage()    {}
func (*ScanInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ebf65bf48ad2fb, []int{0}
}

func (m *ScanInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanInfo.Unmarshal(m, b)
}
func (m *ScanInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanInfo.Marshal(b, m, deterministic)
}
func (m *ScanInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanInfo.Merge(m, src)
}
func (m *ScanInfo) XXX_Size() int {
	return xxx_messageInfo_ScanInfo.Size(m)
}
func (m *ScanInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ScanInfo proto.InternalMessageInfo

func (m *ScanInfo) GetScanUuid() string {
	if m != nil {
		return m.ScanUuid
	}
	return ""
}

func (m *ScanInfo) GetScanStartTime() *timestamppb.Timestamp {
	if m != nil {
		return m.ScanStartTime
	}
	return nil
}

func (m *ScanInfo) GetScanTags() map[string]string {
	if m != nil {
		return m.ScanTags
	}
	return nil
}

// LaunchToolReponse consists of a response built by a producer,
// to be interpreted by a consumer
type LaunchToolResponse struct {
	// The scan information, see above for details
	ScanInfo *ScanInfo `protobuf:"bytes,1,opt,name=scan_info,json=scanInfo,proto3" json:"scan_info,omitempty"`
	// The name of the tool that ran the scan
	ToolName string `protobuf:"bytes,2,opt,name=tool_name,json=toolName,proto3" json:"tool_name,omitempty"`
	// Issues discovered during the scan
	Issues               []*Issue `protobuf:"bytes,3,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LaunchToolResponse) Reset()         { *m = LaunchToolResponse{} }
func (m *LaunchToolResponse) String() string { return proto.CompactTextString(m) }
func (*LaunchToolResponse) ProtoMessage()    {}
func (*LaunchToolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ebf65bf48ad2fb, []int{1}
}

func (m *LaunchToolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchToolResponse.Unmarshal(m, b)
}
func (m *LaunchToolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchToolResponse.Marshal(b, m, deterministic)
}
func (m *LaunchToolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchToolResponse.Merge(m, src)
}
func (m *LaunchToolResponse) XXX_Size() int {
	return xxx_messageInfo_LaunchToolResponse.Size(m)
}
func (m *LaunchToolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchToolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchToolResponse proto.InternalMessageInfo

func (m *LaunchToolResponse) GetScanInfo() *ScanInfo {
	if m != nil {
		return m.ScanInfo
	}
	return nil
}

func (m *LaunchToolResponse) GetToolName() string {
	if m != nil {
		return m.ToolName
	}
	return ""
}

func (m *LaunchToolResponse) GetIssues() []*Issue {
	if m != nil {
		return m.Issues
	}
	return nil
}

// An EnrichedLaunchToolResponse consists of deduplicated vulnerability
// information, with added metadata for consumers
type EnrichedLaunchToolResponse struct {
	// The results of the original scan prior to enrichment
	OriginalResults *LaunchToolResponse `protobuf:"bytes,1,opt,name=original_results,json=originalResults,proto3" json:"original_results,omitempty"`
	// Enriched, deduplicated issues
	Issues               []*EnrichedIssue `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *EnrichedLaunchToolResponse) Reset()         { *m = EnrichedLaunchToolResponse{} }
func (m *EnrichedLaunchToolResponse) String() string { return proto.CompactTextString(m) }
func (*EnrichedLaunchToolResponse) ProtoMessage()    {}
func (*EnrichedLaunchToolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_46ebf65bf48ad2fb, []int{2}
}

func (m *EnrichedLaunchToolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnrichedLaunchToolResponse.Unmarshal(m, b)
}
func (m *EnrichedLaunchToolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnrichedLaunchToolResponse.Marshal(b, m, deterministic)
}
func (m *EnrichedLaunchToolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnrichedLaunchToolResponse.Merge(m, src)
}
func (m *EnrichedLaunchToolResponse) XXX_Size() int {
	return xxx_messageInfo_EnrichedLaunchToolResponse.Size(m)
}
func (m *EnrichedLaunchToolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnrichedLaunchToolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnrichedLaunchToolResponse proto.InternalMessageInfo

func (m *EnrichedLaunchToolResponse) GetOriginalResults() *LaunchToolResponse {
	if m != nil {
		return m.OriginalResults
	}
	return nil
}

func (m *EnrichedLaunchToolResponse) GetIssues() []*EnrichedIssue {
	if m != nil {
		return m.Issues
	}
	return nil
}

func init() {
	proto.RegisterType((*ScanInfo)(nil), "ocurity.dracon.v1.ScanInfo")
	proto.RegisterMapType((map[string]string)(nil), "ocurity.dracon.v1.ScanInfo.ScanTagsEntry")
	proto.RegisterType((*LaunchToolResponse)(nil), "ocurity.dracon.v1.LaunchToolResponse")
	proto.RegisterType((*EnrichedLaunchToolResponse)(nil), "ocurity.dracon.v1.EnrichedLaunchToolResponse")
}

func init() {
	proto.RegisterFile("api/proto/v1/engine.proto", fileDescriptor_46ebf65bf48ad2fb)
}

var fileDescriptor_46ebf65bf48ad2fb = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4d, 0x8b, 0xda, 0x40,
	0x1c, 0xc6, 0x89, 0x52, 0xd1, 0x11, 0xd1, 0x0e, 0x3d, 0xa4, 0xf1, 0xd0, 0x20, 0xb4, 0xa4, 0x97,
	0x49, 0xb5, 0x17, 0x69, 0x6f, 0x82, 0x05, 0xa1, 0x94, 0x12, 0xed, 0xa5, 0x97, 0x30, 0x26, 0x63,
	0x1c, 0x9a, 0xcc, 0x84, 0x79, 0x11, 0xfc, 0x30, 0xbd, 0xef, 0xc7, 0xdb, 0x8f, 0xb0, 0x4c, 0x26,
	0xb3, 0xac, 0x18, 0xf6, 0xf6, 0x7f, 0x7b, 0x9e, 0xe7, 0x37, 0x1a, 0xf0, 0x1e, 0xd7, 0x34, 0xae,
	0x05, 0x57, 0x3c, 0xbe, 0x2c, 0x63, 0xc2, 0x0a, 0xca, 0x08, 0x6a, 0x7a, 0xf8, 0x96, 0x67, 0x5a,
	0x50, 0x75, 0x45, 0xb9, 0xc0, 0x19, 0x67, 0xe8, 0xb2, 0x0c, 0xfc, 0x9b, 0x6b, 0x2a, 0xa5, 0x6e,
	0x8f, 0x83, 0x0f, 0x05, 0xe7, 0x45, 0x49, 0xec, 0xf2, 0xa8, 0x4f, 0xb1, 0xa2, 0x15, 0x91, 0x0a,
	0x57, 0xb5, 0x3d, 0x58, 0x3c, 0x7a, 0x60, 0xb8, 0xcf, 0x30, 0xdb, 0xb1, 0x13, 0x87, 0x73, 0x30,
	0x92, 0x19, 0x66, 0xa9, 0xd6, 0x34, 0xf7, 0xbd, 0xd0, 0x8b, 0x46, 0xc9, 0xd0, 0x0c, 0xfe, 0x68,
	0x9a, 0xc3, 0x0d, 0x98, 0x36, 0x4b, 0xa9, 0xb0, 0x50, 0xa9, 0xf1, 0xf1, 0x7b, 0xa1, 0x17, 0x8d,
	0x57, 0x01, 0xb2, 0x21, 0xc8, 0x85, 0xa0, 0x83, 0x0b, 0x49, 0x26, 0x46, 0xb2, 0x37, 0x0a, 0x33,
	0x83, 0x3f, 0xda, 0x00, 0x85, 0x0b, 0xe9, 0xf7, 0xc3, 0x7e, 0x34, 0x5e, 0x7d, 0x46, 0x77, 0xef,
	0x41, 0x0e, 0xa8, 0x29, 0x0e, 0xb8, 0x90, 0x5b, 0xa6, 0xc4, 0xd5, 0xb2, 0x98, 0x36, 0xf8, 0x0e,
	0x26, 0x37, 0x2b, 0x38, 0x03, 0xfd, 0x7f, 0xe4, 0xda, 0x32, 0x9b, 0x12, 0xbe, 0x03, 0x6f, 0x2e,
	0xb8, 0xd4, 0x16, 0x72, 0x94, 0xd8, 0xe6, 0x5b, 0x6f, 0xed, 0x2d, 0xfe, 0x7b, 0x00, 0xfe, 0xc4,
	0x9a, 0x65, 0xe7, 0x03, 0xe7, 0x65, 0x42, 0x64, 0xcd, 0x99, 0x24, 0x70, 0xdd, 0xb2, 0x51, 0x76,
	0xe2, 0x8d, 0xd1, 0x78, 0x35, 0x7f, 0x85, 0xcd, 0xd2, 0xb8, 0x9f, 0x4d, 0x71, 0x5e, 0xa6, 0x0c,
	0x57, 0x2e, 0x6e, 0x68, 0x06, 0xbf, 0x70, 0x45, 0xe0, 0x17, 0x30, 0x68, 0xfe, 0x10, 0xf7, 0x5e,
	0xbf, 0xc3, 0x73, 0x67, 0x0e, 0x92, 0xf6, 0x6e, 0xf1, 0xe0, 0x81, 0x60, 0xcb, 0x04, 0xcd, 0xce,
	0x24, 0xef, 0xe0, 0xfc, 0x0d, 0x66, 0x5c, 0xd0, 0x82, 0x32, 0x5c, 0xa6, 0x82, 0x48, 0x5d, 0x2a,
	0xd9, 0xe2, 0x7e, 0xec, 0xb0, 0xbe, 0x37, 0x48, 0xa6, 0x4e, 0x9e, 0x58, 0x35, 0x5c, 0x3f, 0x23,
	0xf6, 0x1a, 0xc4, 0xb0, 0xc3, 0xc7, 0x01, 0xdd, 0xa0, 0x6e, 0xa2, 0xbf, 0x9f, 0x0a, 0xaa, 0xce,
	0xfa, 0x88, 0x32, 0x5e, 0xc5, 0xad, 0x2a, 0xb6, 0xaa, 0xf8, 0xe5, 0x47, 0x79, 0x1c, 0x34, 0xd5,
	0xd7, 0xa7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xe7, 0x02, 0x3e, 0xd9, 0x02, 0x00, 0x00,
}
