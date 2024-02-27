// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: api/proto/v1/issue.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Confidence represents the tool's confidence that an issue exists
type Confidence int32

const (
	// Represents an empty confidence field from a tool
	Confidence_CONFIDENCE_UNSPECIFIED Confidence = 0
	// Represents Confidence "Informational" or generic Warnings
	Confidence_CONFIDENCE_INFO Confidence = 1
	// Represents Confidence "Low"
	Confidence_CONFIDENCE_LOW Confidence = 2
	// Represents Confidence "Medium"
	Confidence_CONFIDENCE_MEDIUM Confidence = 3
	// Represents Confidence "High"
	Confidence_CONFIDENCE_HIGH Confidence = 4
	// Highest Confidence applicable
	Confidence_CONFIDENCE_CRITICAL Confidence = 5
)

// Enum value maps for Confidence.
var (
	Confidence_name = map[int32]string{
		0: "CONFIDENCE_UNSPECIFIED",
		1: "CONFIDENCE_INFO",
		2: "CONFIDENCE_LOW",
		3: "CONFIDENCE_MEDIUM",
		4: "CONFIDENCE_HIGH",
		5: "CONFIDENCE_CRITICAL",
	}
	Confidence_value = map[string]int32{
		"CONFIDENCE_UNSPECIFIED": 0,
		"CONFIDENCE_INFO":        1,
		"CONFIDENCE_LOW":         2,
		"CONFIDENCE_MEDIUM":      3,
		"CONFIDENCE_HIGH":        4,
		"CONFIDENCE_CRITICAL":    5,
	}
)

func (x Confidence) Enum() *Confidence {
	p := new(Confidence)
	*p = x
	return p
}

func (x Confidence) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Confidence) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_issue_proto_enumTypes[0].Descriptor()
}

func (Confidence) Type() protoreflect.EnumType {
	return &file_api_proto_v1_issue_proto_enumTypes[0]
}

func (x Confidence) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Confidence.Descriptor instead.
func (Confidence) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_issue_proto_rawDescGZIP(), []int{0}
}

// Severity represents the reported issue severity
type Severity int32

const (
	// Reserved in case a tool does not report severity
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	// Informational priority findings
	Severity_SEVERITY_INFO Severity = 1
	// Low priority findings
	Severity_SEVERITY_LOW Severity = 2
	// Medium priority findings
	Severity_SEVERITY_MEDIUM Severity = 3
	// High priority findings
	Severity_SEVERITY_HIGH Severity = 4
	// Critical priority findings
	Severity_SEVERITY_CRITICAL Severity = 5
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_INFO",
		2: "SEVERITY_LOW",
		3: "SEVERITY_MEDIUM",
		4: "SEVERITY_HIGH",
		5: "SEVERITY_CRITICAL",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_INFO":        1,
		"SEVERITY_LOW":         2,
		"SEVERITY_MEDIUM":      3,
		"SEVERITY_HIGH":        4,
		"SEVERITY_CRITICAL":    5,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_v1_issue_proto_enumTypes[1].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_api_proto_v1_issue_proto_enumTypes[1]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_v1_issue_proto_rawDescGZIP(), []int{1}
}

// Issue represents a vulnerability to be processed by consumers
type Issue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Scan target can be host:port, //vault/foo/bar:34-67 or some URL that is semantically a target
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// The finding ID from the tool if applicable, otherwise a vulnerability ID such as CWE-ID, etc for XSS, CSRF, etc.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The finding title from the tool
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// A severity indication, defaults to unspecified
	Severity Severity `protobuf:"varint,4,opt,name=severity,proto3,enum=ocurity.dracon.v1.Severity" json:"severity,omitempty"`
	// An optional cvss if the tool reports it
	Cvss float64 `protobuf:"fixed64,5,opt,name=cvss,proto3" json:"cvss,omitempty"`
	// Confidence indication, defaults to Unspecified
	Confidence Confidence `protobuf:"varint,6,opt,name=confidence,proto3,enum=ocurity.dracon.v1.Confidence" json:"confidence,omitempty"`
	// human readable description of the issue
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// Source of the scan e.g. https://github.com/ocurity/dracon.git?ref=<revision>, github.com:tektoncd/pipeline.git?ref=<revision>, local?ref=local
	Source string `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
	// [Optional] the CVE causing this vulnerability
	Cve string `protobuf:"bytes,9,opt,name=cve,proto3" json:"cve,omitempty"`
	// internal field reserved for the enrichment aggregator
	Uuid string `protobuf:"bytes,10,opt,name=uuid,proto3" json:"uuid,omitempty"`
	// optional field that allows us to also encode a bill of materials in an issue
	CycloneDXSBOM *string `protobuf:"bytes,11,opt,name=cyclone_d_x_s_b_o_m,json=cycloneDXSBOM,proto3,oneof" json:"cyclone_d_x_s_b_o_m,omitempty"`
	// optional string that allows producers to communicate relevant code/request segments
	ContextSegment *string `protobuf:"bytes,12,opt,name=context_segment,json=contextSegment,proto3,oneof" json:"context_segment,omitempty"`
	// optionally the related CWE
	Cwe *string `protobuf:"bytes,13,opt,name=cwe,proto3,oneof" json:"cwe,omitempty"`
}

func (x *Issue) Reset() {
	*x = Issue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_issue_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Issue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Issue) ProtoMessage() {}

func (x *Issue) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_issue_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Issue.ProtoReflect.Descriptor instead.
func (*Issue) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_issue_proto_rawDescGZIP(), []int{0}
}

func (x *Issue) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Issue) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Issue) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Issue) GetSeverity() Severity {
	if x != nil {
		return x.Severity
	}
	return Severity_SEVERITY_UNSPECIFIED
}

func (x *Issue) GetCvss() float64 {
	if x != nil {
		return x.Cvss
	}
	return 0
}

func (x *Issue) GetConfidence() Confidence {
	if x != nil {
		return x.Confidence
	}
	return Confidence_CONFIDENCE_UNSPECIFIED
}

func (x *Issue) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Issue) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Issue) GetCve() string {
	if x != nil {
		return x.Cve
	}
	return ""
}

func (x *Issue) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Issue) GetCycloneDXSBOM() string {
	if x != nil && x.CycloneDXSBOM != nil {
		return *x.CycloneDXSBOM
	}
	return ""
}

func (x *Issue) GetContextSegment() string {
	if x != nil && x.ContextSegment != nil {
		return *x.ContextSegment
	}
	return ""
}

func (x *Issue) GetCwe() string {
	if x != nil && x.Cwe != nil {
		return *x.Cwe
	}
	return ""
}

// Represents an issue that has been enriched with metadata from the enrichment service
type EnrichedIssue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the original finding
	RawIssue *Issue `protobuf:"bytes,1,opt,name=raw_issue,json=rawIssue,proto3" json:"raw_issue,omitempty"`
	// The first time this issue was seen by the enrichment service
	FirstSeen *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=first_seen,json=firstSeen,proto3" json:"first_seen,omitempty"`
	// The number of times this issue was seen
	Count uint64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	// Whether this issue has been previously marked as a false positive
	FalsePositive bool `protobuf:"varint,4,opt,name=false_positive,json=falsePositive,proto3" json:"false_positive,omitempty"`
	// The last time this issue was updated
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// hash
	Hash string `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	// an arbitrary list of extra annotations, reserved for use by the enrichers
	Annotations map[string]string `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EnrichedIssue) Reset() {
	*x = EnrichedIssue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_v1_issue_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrichedIssue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrichedIssue) ProtoMessage() {}

func (x *EnrichedIssue) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_v1_issue_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrichedIssue.ProtoReflect.Descriptor instead.
func (*EnrichedIssue) Descriptor() ([]byte, []int) {
	return file_api_proto_v1_issue_proto_rawDescGZIP(), []int{1}
}

func (x *EnrichedIssue) GetRawIssue() *Issue {
	if x != nil {
		return x.RawIssue
	}
	return nil
}

func (x *EnrichedIssue) GetFirstSeen() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstSeen
	}
	return nil
}

func (x *EnrichedIssue) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *EnrichedIssue) GetFalsePositive() bool {
	if x != nil {
		return x.FalsePositive
	}
	return false
}

func (x *EnrichedIssue) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *EnrichedIssue) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *EnrichedIssue) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

var File_api_proto_v1_issue_proto protoreflect.FileDescriptor

var file_api_proto_v1_issue_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x6f, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf,
	0x03, 0x0a, 0x05, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x76, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x63, 0x76, 0x73, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63,
	0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e,
	0x65, 0x5f, 0x64, 0x5f, 0x78, 0x5f, 0x73, 0x5f, 0x62, 0x5f, 0x6f, 0x5f, 0x6d, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x44, 0x58,
	0x53, 0x42, 0x4f, 0x4d, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x63, 0x77, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x63, 0x77, 0x65, 0x88, 0x01, 0x01, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x63, 0x79, 0x63, 0x6c, 0x6f, 0x6e, 0x65, 0x5f, 0x64, 0x5f, 0x78, 0x5f, 0x73, 0x5f, 0x62,
	0x5f, 0x6f, 0x5f, 0x6d, 0x42, 0x12, 0x0a, 0x10, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x63, 0x77, 0x65,
	0x22, 0xa2, 0x03, 0x0a, 0x0d, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68, 0x65, 0x64, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52,
	0x08, 0x72, 0x61, 0x77, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x53, 0x65, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61,
	0x6c, 0x73, 0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x53, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6f, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e,
	0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x69, 0x63, 0x68,
	0x65, 0x64, 0x49, 0x73, 0x73, 0x75, 0x65, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x96, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45, 0x4e,
	0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x49,
	0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x4f, 0x4e,
	0x46, 0x49, 0x44, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03,
	0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45, 0x4e, 0x43, 0x45, 0x5f, 0x48,
	0x49, 0x47, 0x48, 0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x44, 0x45,
	0x4e, 0x43, 0x45, 0x5f, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x2a, 0x88,
	0x01, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14, 0x53,
	0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54,
	0x59, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x56, 0x45,
	0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45,
	0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03, 0x12,
	0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48,
	0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x43,
	0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f,
	0x64, 0x72, 0x61, 0x63, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_v1_issue_proto_rawDescOnce sync.Once
	file_api_proto_v1_issue_proto_rawDescData = file_api_proto_v1_issue_proto_rawDesc
)

func file_api_proto_v1_issue_proto_rawDescGZIP() []byte {
	file_api_proto_v1_issue_proto_rawDescOnce.Do(func() {
		file_api_proto_v1_issue_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_v1_issue_proto_rawDescData)
	})
	return file_api_proto_v1_issue_proto_rawDescData
}

var file_api_proto_v1_issue_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_proto_v1_issue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_v1_issue_proto_goTypes = []interface{}{
	(Confidence)(0),               // 0: ocurity.dracon.v1.Confidence
	(Severity)(0),                 // 1: ocurity.dracon.v1.Severity
	(*Issue)(nil),                 // 2: ocurity.dracon.v1.Issue
	(*EnrichedIssue)(nil),         // 3: ocurity.dracon.v1.EnrichedIssue
	nil,                           // 4: ocurity.dracon.v1.EnrichedIssue.AnnotationsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_api_proto_v1_issue_proto_depIdxs = []int32{
	1, // 0: ocurity.dracon.v1.Issue.severity:type_name -> ocurity.dracon.v1.Severity
	0, // 1: ocurity.dracon.v1.Issue.confidence:type_name -> ocurity.dracon.v1.Confidence
	2, // 2: ocurity.dracon.v1.EnrichedIssue.raw_issue:type_name -> ocurity.dracon.v1.Issue
	5, // 3: ocurity.dracon.v1.EnrichedIssue.first_seen:type_name -> google.protobuf.Timestamp
	5, // 4: ocurity.dracon.v1.EnrichedIssue.updated_at:type_name -> google.protobuf.Timestamp
	4, // 5: ocurity.dracon.v1.EnrichedIssue.annotations:type_name -> ocurity.dracon.v1.EnrichedIssue.AnnotationsEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_proto_v1_issue_proto_init() }
func file_api_proto_v1_issue_proto_init() {
	if File_api_proto_v1_issue_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_v1_issue_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Issue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_v1_issue_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrichedIssue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_api_proto_v1_issue_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_issue_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_v1_issue_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_issue_proto_depIdxs,
		EnumInfos:         file_api_proto_v1_issue_proto_enumTypes,
		MessageInfos:      file_api_proto_v1_issue_proto_msgTypes,
	}.Build()
	File_api_proto_v1_issue_proto = out.File
	file_api_proto_v1_issue_proto_rawDesc = nil
	file_api_proto_v1_issue_proto_goTypes = nil
	file_api_proto_v1_issue_proto_depIdxs = nil
}
