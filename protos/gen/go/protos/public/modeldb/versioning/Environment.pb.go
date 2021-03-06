// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/public/modeldb/versioning/Environment.proto

package versioning

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

type EnvironmentBlob struct {
	// Types that are valid to be assigned to Content:
	//	*EnvironmentBlob_Python
	//	*EnvironmentBlob_Docker
	Content              isEnvironmentBlob_Content   `protobuf_oneof:"content"`
	EnvironmentVariables []*EnvironmentVariablesBlob `protobuf:"bytes,3,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty"`
	CommandLine          []string                    `protobuf:"bytes,4,rep,name=command_line,json=commandLine,proto3" json:"command_line,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EnvironmentBlob) Reset()         { *m = EnvironmentBlob{} }
func (m *EnvironmentBlob) String() string { return proto.CompactTextString(m) }
func (*EnvironmentBlob) ProtoMessage()    {}
func (*EnvironmentBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e422fc6faa4b486, []int{0}
}

func (m *EnvironmentBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentBlob.Unmarshal(m, b)
}
func (m *EnvironmentBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentBlob.Marshal(b, m, deterministic)
}
func (m *EnvironmentBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentBlob.Merge(m, src)
}
func (m *EnvironmentBlob) XXX_Size() int {
	return xxx_messageInfo_EnvironmentBlob.Size(m)
}
func (m *EnvironmentBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentBlob.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentBlob proto.InternalMessageInfo

type isEnvironmentBlob_Content interface {
	isEnvironmentBlob_Content()
}

type EnvironmentBlob_Python struct {
	Python *PythonEnvironmentBlob `protobuf:"bytes,1,opt,name=python,proto3,oneof"`
}

type EnvironmentBlob_Docker struct {
	Docker *DockerEnvironmentBlob `protobuf:"bytes,2,opt,name=docker,proto3,oneof"`
}

func (*EnvironmentBlob_Python) isEnvironmentBlob_Content() {}

func (*EnvironmentBlob_Docker) isEnvironmentBlob_Content() {}

func (m *EnvironmentBlob) GetContent() isEnvironmentBlob_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *EnvironmentBlob) GetPython() *PythonEnvironmentBlob {
	if x, ok := m.GetContent().(*EnvironmentBlob_Python); ok {
		return x.Python
	}
	return nil
}

func (m *EnvironmentBlob) GetDocker() *DockerEnvironmentBlob {
	if x, ok := m.GetContent().(*EnvironmentBlob_Docker); ok {
		return x.Docker
	}
	return nil
}

func (m *EnvironmentBlob) GetEnvironmentVariables() []*EnvironmentVariablesBlob {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

func (m *EnvironmentBlob) GetCommandLine() []string {
	if m != nil {
		return m.CommandLine
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EnvironmentBlob) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EnvironmentBlob_Python)(nil),
		(*EnvironmentBlob_Docker)(nil),
	}
}

type EnvironmentVariablesBlob struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironmentVariablesBlob) Reset()         { *m = EnvironmentVariablesBlob{} }
func (m *EnvironmentVariablesBlob) String() string { return proto.CompactTextString(m) }
func (*EnvironmentVariablesBlob) ProtoMessage()    {}
func (*EnvironmentVariablesBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e422fc6faa4b486, []int{1}
}

func (m *EnvironmentVariablesBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentVariablesBlob.Unmarshal(m, b)
}
func (m *EnvironmentVariablesBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentVariablesBlob.Marshal(b, m, deterministic)
}
func (m *EnvironmentVariablesBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentVariablesBlob.Merge(m, src)
}
func (m *EnvironmentVariablesBlob) XXX_Size() int {
	return xxx_messageInfo_EnvironmentVariablesBlob.Size(m)
}
func (m *EnvironmentVariablesBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentVariablesBlob.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentVariablesBlob proto.InternalMessageInfo

func (m *EnvironmentVariablesBlob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EnvironmentVariablesBlob) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VersionEnvironmentBlob struct {
	Major                int32    `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                int32    `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                int32    `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	Suffix               string   `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionEnvironmentBlob) Reset()         { *m = VersionEnvironmentBlob{} }
func (m *VersionEnvironmentBlob) String() string { return proto.CompactTextString(m) }
func (*VersionEnvironmentBlob) ProtoMessage()    {}
func (*VersionEnvironmentBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e422fc6faa4b486, []int{2}
}

func (m *VersionEnvironmentBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionEnvironmentBlob.Unmarshal(m, b)
}
func (m *VersionEnvironmentBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionEnvironmentBlob.Marshal(b, m, deterministic)
}
func (m *VersionEnvironmentBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionEnvironmentBlob.Merge(m, src)
}
func (m *VersionEnvironmentBlob) XXX_Size() int {
	return xxx_messageInfo_VersionEnvironmentBlob.Size(m)
}
func (m *VersionEnvironmentBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionEnvironmentBlob.DiscardUnknown(m)
}

var xxx_messageInfo_VersionEnvironmentBlob proto.InternalMessageInfo

func (m *VersionEnvironmentBlob) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *VersionEnvironmentBlob) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *VersionEnvironmentBlob) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *VersionEnvironmentBlob) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

type PythonRequirementEnvironmentBlob struct {
	Library              string                  `protobuf:"bytes,1,opt,name=library,proto3" json:"library,omitempty"`
	Constraint           string                  `protobuf:"bytes,2,opt,name=constraint,proto3" json:"constraint,omitempty"`
	Version              *VersionEnvironmentBlob `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PythonRequirementEnvironmentBlob) Reset()         { *m = PythonRequirementEnvironmentBlob{} }
func (m *PythonRequirementEnvironmentBlob) String() string { return proto.CompactTextString(m) }
func (*PythonRequirementEnvironmentBlob) ProtoMessage()    {}
func (*PythonRequirementEnvironmentBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e422fc6faa4b486, []int{3}
}

func (m *PythonRequirementEnvironmentBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PythonRequirementEnvironmentBlob.Unmarshal(m, b)
}
func (m *PythonRequirementEnvironmentBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PythonRequirementEnvironmentBlob.Marshal(b, m, deterministic)
}
func (m *PythonRequirementEnvironmentBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PythonRequirementEnvironmentBlob.Merge(m, src)
}
func (m *PythonRequirementEnvironmentBlob) XXX_Size() int {
	return xxx_messageInfo_PythonRequirementEnvironmentBlob.Size(m)
}
func (m *PythonRequirementEnvironmentBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_PythonRequirementEnvironmentBlob.DiscardUnknown(m)
}

var xxx_messageInfo_PythonRequirementEnvironmentBlob proto.InternalMessageInfo

func (m *PythonRequirementEnvironmentBlob) GetLibrary() string {
	if m != nil {
		return m.Library
	}
	return ""
}

func (m *PythonRequirementEnvironmentBlob) GetConstraint() string {
	if m != nil {
		return m.Constraint
	}
	return ""
}

func (m *PythonRequirementEnvironmentBlob) GetVersion() *VersionEnvironmentBlob {
	if m != nil {
		return m.Version
	}
	return nil
}

type PythonEnvironmentBlob struct {
	Version              *VersionEnvironmentBlob             `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Requirements         []*PythonRequirementEnvironmentBlob `protobuf:"bytes,2,rep,name=requirements,proto3" json:"requirements,omitempty"`
	Constraints          []*PythonRequirementEnvironmentBlob `protobuf:"bytes,3,rep,name=constraints,proto3" json:"constraints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *PythonEnvironmentBlob) Reset()         { *m = PythonEnvironmentBlob{} }
func (m *PythonEnvironmentBlob) String() string { return proto.CompactTextString(m) }
func (*PythonEnvironmentBlob) ProtoMessage()    {}
func (*PythonEnvironmentBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e422fc6faa4b486, []int{4}
}

func (m *PythonEnvironmentBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PythonEnvironmentBlob.Unmarshal(m, b)
}
func (m *PythonEnvironmentBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PythonEnvironmentBlob.Marshal(b, m, deterministic)
}
func (m *PythonEnvironmentBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PythonEnvironmentBlob.Merge(m, src)
}
func (m *PythonEnvironmentBlob) XXX_Size() int {
	return xxx_messageInfo_PythonEnvironmentBlob.Size(m)
}
func (m *PythonEnvironmentBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_PythonEnvironmentBlob.DiscardUnknown(m)
}

var xxx_messageInfo_PythonEnvironmentBlob proto.InternalMessageInfo

func (m *PythonEnvironmentBlob) GetVersion() *VersionEnvironmentBlob {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *PythonEnvironmentBlob) GetRequirements() []*PythonRequirementEnvironmentBlob {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *PythonEnvironmentBlob) GetConstraints() []*PythonRequirementEnvironmentBlob {
	if m != nil {
		return m.Constraints
	}
	return nil
}

type DockerEnvironmentBlob struct {
	Repository           string   `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	Sha                  string   `protobuf:"bytes,3,opt,name=sha,proto3" json:"sha,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerEnvironmentBlob) Reset()         { *m = DockerEnvironmentBlob{} }
func (m *DockerEnvironmentBlob) String() string { return proto.CompactTextString(m) }
func (*DockerEnvironmentBlob) ProtoMessage()    {}
func (*DockerEnvironmentBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e422fc6faa4b486, []int{5}
}

func (m *DockerEnvironmentBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerEnvironmentBlob.Unmarshal(m, b)
}
func (m *DockerEnvironmentBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerEnvironmentBlob.Marshal(b, m, deterministic)
}
func (m *DockerEnvironmentBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerEnvironmentBlob.Merge(m, src)
}
func (m *DockerEnvironmentBlob) XXX_Size() int {
	return xxx_messageInfo_DockerEnvironmentBlob.Size(m)
}
func (m *DockerEnvironmentBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerEnvironmentBlob.DiscardUnknown(m)
}

var xxx_messageInfo_DockerEnvironmentBlob proto.InternalMessageInfo

func (m *DockerEnvironmentBlob) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

func (m *DockerEnvironmentBlob) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *DockerEnvironmentBlob) GetSha() string {
	if m != nil {
		return m.Sha
	}
	return ""
}

func init() {
	proto.RegisterType((*EnvironmentBlob)(nil), "ai.verta.modeldb.versioning.EnvironmentBlob")
	proto.RegisterType((*EnvironmentVariablesBlob)(nil), "ai.verta.modeldb.versioning.EnvironmentVariablesBlob")
	proto.RegisterType((*VersionEnvironmentBlob)(nil), "ai.verta.modeldb.versioning.VersionEnvironmentBlob")
	proto.RegisterType((*PythonRequirementEnvironmentBlob)(nil), "ai.verta.modeldb.versioning.PythonRequirementEnvironmentBlob")
	proto.RegisterType((*PythonEnvironmentBlob)(nil), "ai.verta.modeldb.versioning.PythonEnvironmentBlob")
	proto.RegisterType((*DockerEnvironmentBlob)(nil), "ai.verta.modeldb.versioning.DockerEnvironmentBlob")
}

func init() {
	proto.RegisterFile("protos/public/modeldb/versioning/Environment.proto", fileDescriptor_3e422fc6faa4b486)
}

var fileDescriptor_3e422fc6faa4b486 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0x97, 0x38, 0x3f, 0xf0, 0x73, 0x61, 0x43, 0x24, 0xc5, 0x30, 0x18, 0x99, 0x4f, 0x39,
	0xd9, 0xe0, 0xb2, 0xe3, 0x0e, 0x0b, 0x1d, 0xac, 0xd0, 0x42, 0xd1, 0x21, 0x87, 0xed, 0x10, 0x64,
	0x47, 0x75, 0xd4, 0xd9, 0x92, 0x27, 0xcb, 0x66, 0xfd, 0x97, 0xb6, 0x3f, 0x64, 0xff, 0xd6, 0x90,
	0xac, 0x34, 0x5e, 0x70, 0xb3, 0x42, 0x6f, 0x7a, 0x5f, 0x3d, 0x7d, 0x9e, 0xf4, 0xbe, 0x92, 0x20,
	0x2e, 0xa5, 0x50, 0xa2, 0x8a, 0xca, 0x3a, 0xc9, 0x59, 0x1a, 0x15, 0x62, 0x4b, 0xf3, 0x6d, 0x12,
	0x35, 0x54, 0x56, 0x4c, 0x70, 0xc6, 0xb3, 0xe8, 0x33, 0x6f, 0x98, 0x14, 0xbc, 0xa0, 0x5c, 0x85,
	0x26, 0x19, 0xbd, 0x25, 0x2c, 0x6c, 0xa8, 0x54, 0x24, 0xb4, 0xe9, 0xe1, 0x21, 0x3d, 0xf8, 0x33,
	0x84, 0xd7, 0x9d, 0x25, 0xab, 0x5c, 0x24, 0xe8, 0x1a, 0x26, 0xe5, 0x83, 0xda, 0x09, 0xee, 0x0f,
	0x16, 0x83, 0xa5, 0x17, 0xc7, 0xe1, 0x09, 0x42, 0x78, 0x6b, 0x52, 0x8f, 0x18, 0x5f, 0x5e, 0x61,
	0xcb, 0xd0, 0xb4, 0xad, 0x48, 0xbf, 0x53, 0xe9, 0x0f, 0x9f, 0x41, 0xbb, 0x34, 0xa9, 0x3d, 0xb4,
	0x96, 0x81, 0xee, 0x61, 0x4e, 0x0f, 0x93, 0x9b, 0x86, 0x48, 0x46, 0x92, 0x9c, 0x56, 0xbe, 0xb3,
	0x70, 0x96, 0x5e, 0xfc, 0xe1, 0x24, 0xbc, 0x83, 0x5d, 0xef, 0x17, 0x6a, 0x3e, 0x9e, 0xd1, 0x9e,
	0x19, 0xf4, 0x1e, 0xce, 0x52, 0x51, 0x14, 0x84, 0x6f, 0x37, 0x39, 0xe3, 0xd4, 0x1f, 0x2d, 0x9c,
	0xa5, 0x8b, 0x3d, 0xab, 0x5d, 0x33, 0x4e, 0x57, 0x2e, 0x4c, 0x53, 0xc1, 0x15, 0xe5, 0x2a, 0xb8,
	0x04, 0xff, 0x29, 0x3e, 0x42, 0x30, 0xe2, 0xa4, 0xa0, 0xa6, 0x9f, 0x2e, 0x36, 0x63, 0x34, 0x83,
	0x71, 0x43, 0xf2, 0x9a, 0x9a, 0xb6, 0xb8, 0xb8, 0x0d, 0x02, 0x09, 0xe7, 0xeb, 0x76, 0xc3, 0xc7,
	0xae, 0xcc, 0x60, 0x5c, 0x90, 0x7b, 0x21, 0x0d, 0x64, 0x8c, 0xdb, 0xc0, 0xa8, 0x8c, 0x8b, 0xb6,
	0xb9, 0x5a, 0xd5, 0x81, 0x56, 0x4b, 0xa2, 0xd2, 0x9d, 0xef, 0xb4, 0xaa, 0x09, 0xd0, 0x39, 0x4c,
	0xaa, 0xfa, 0xee, 0x8e, 0xfd, 0xf4, 0x47, 0xa6, 0xa4, 0x8d, 0x82, 0xdf, 0x03, 0x58, 0xb4, 0x2e,
	0x62, 0xfa, 0xa3, 0x66, 0x92, 0xea, 0x9a, 0xc7, 0xe5, 0x7d, 0x98, 0xe6, 0x2c, 0x91, 0x44, 0x3e,
	0xd8, 0x53, 0xec, 0x43, 0xf4, 0x0e, 0x20, 0x15, 0xbc, 0x52, 0x92, 0x30, 0xae, 0xec, 0x69, 0x3a,
	0x0a, 0xba, 0x81, 0xa9, 0xf5, 0xc0, 0x6c, 0xc7, 0x8b, 0x2f, 0x4e, 0x9a, 0xd4, 0x7f, 0x7c, 0xbc,
	0x67, 0x04, 0xbf, 0x86, 0x30, 0xef, 0xbd, 0x73, 0xdd, 0x42, 0x83, 0x97, 0x17, 0x42, 0x04, 0xce,
	0xe4, 0xa1, 0x1f, 0x95, 0x3f, 0x34, 0x37, 0xec, 0xe3, 0x33, 0x1e, 0xc3, 0xd3, 0x6d, 0xc4, 0xff,
	0x20, 0xd1, 0x06, 0xbc, 0x43, 0xa3, 0xf6, 0x77, 0xf8, 0x85, 0x15, 0xba, 0xc4, 0xe0, 0x1b, 0xcc,
	0x7b, 0x5f, 0x94, 0x36, 0x4d, 0xd2, 0x52, 0x54, 0x4c, 0x89, 0x47, 0x47, 0x3b, 0x0a, 0x7a, 0x03,
	0x8e, 0x22, 0x99, 0x75, 0x53, 0x0f, 0xb5, 0x52, 0xed, 0x88, 0xb1, 0xd0, 0xc5, 0x7a, 0xb8, 0xba,
	0xb9, 0x1d, 0x7c, 0xbd, 0xca, 0x98, 0xda, 0xd5, 0x49, 0x98, 0x8a, 0x22, 0x5a, 0xeb, 0x3d, 0x7f,
	0xba, 0x7a, 0xfc, 0x95, 0xec, 0x5f, 0x95, 0x51, 0x1e, 0x65, 0x22, 0xfa, 0xdf, 0xcf, 0x95, 0x4c,
	0x4c, 0xc6, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbe, 0xad, 0x7f, 0x31, 0xe4, 0x04, 0x00,
	0x00,
}
