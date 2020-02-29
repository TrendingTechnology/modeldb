// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/public/modeldb/versioning/Config.proto

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

type ConfigBlob struct {
	HyperparameterSet    []*HyperparameterSetConfigBlob `protobuf:"bytes,1,rep,name=hyperparameter_set,json=hyperparameterSet,proto3" json:"hyperparameter_set,omitempty"`
	Hyperparameters      []*HyperparameterConfigBlob    `protobuf:"bytes,2,rep,name=hyperparameters,proto3" json:"hyperparameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ConfigBlob) Reset()         { *m = ConfigBlob{} }
func (m *ConfigBlob) String() string { return proto.CompactTextString(m) }
func (*ConfigBlob) ProtoMessage()    {}
func (*ConfigBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5245ebe5bb808c, []int{0}
}

func (m *ConfigBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigBlob.Unmarshal(m, b)
}
func (m *ConfigBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigBlob.Marshal(b, m, deterministic)
}
func (m *ConfigBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigBlob.Merge(m, src)
}
func (m *ConfigBlob) XXX_Size() int {
	return xxx_messageInfo_ConfigBlob.Size(m)
}
func (m *ConfigBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigBlob.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigBlob proto.InternalMessageInfo

func (m *ConfigBlob) GetHyperparameterSet() []*HyperparameterSetConfigBlob {
	if m != nil {
		return m.HyperparameterSet
	}
	return nil
}

func (m *ConfigBlob) GetHyperparameters() []*HyperparameterConfigBlob {
	if m != nil {
		return m.Hyperparameters
	}
	return nil
}

type HyperparameterConfigBlob struct {
	Name                 string                          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                *HyperparameterValuesConfigBlob `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *HyperparameterConfigBlob) Reset()         { *m = HyperparameterConfigBlob{} }
func (m *HyperparameterConfigBlob) String() string { return proto.CompactTextString(m) }
func (*HyperparameterConfigBlob) ProtoMessage()    {}
func (*HyperparameterConfigBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5245ebe5bb808c, []int{1}
}

func (m *HyperparameterConfigBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HyperparameterConfigBlob.Unmarshal(m, b)
}
func (m *HyperparameterConfigBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HyperparameterConfigBlob.Marshal(b, m, deterministic)
}
func (m *HyperparameterConfigBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HyperparameterConfigBlob.Merge(m, src)
}
func (m *HyperparameterConfigBlob) XXX_Size() int {
	return xxx_messageInfo_HyperparameterConfigBlob.Size(m)
}
func (m *HyperparameterConfigBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_HyperparameterConfigBlob.DiscardUnknown(m)
}

var xxx_messageInfo_HyperparameterConfigBlob proto.InternalMessageInfo

func (m *HyperparameterConfigBlob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HyperparameterConfigBlob) GetValue() *HyperparameterValuesConfigBlob {
	if m != nil {
		return m.Value
	}
	return nil
}

type HyperparameterValuesConfigBlob struct {
	// Types that are valid to be assigned to Value:
	//	*HyperparameterValuesConfigBlob_IntValue
	//	*HyperparameterValuesConfigBlob_FloatValue
	//	*HyperparameterValuesConfigBlob_StringValue
	Value                isHyperparameterValuesConfigBlob_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *HyperparameterValuesConfigBlob) Reset()         { *m = HyperparameterValuesConfigBlob{} }
func (m *HyperparameterValuesConfigBlob) String() string { return proto.CompactTextString(m) }
func (*HyperparameterValuesConfigBlob) ProtoMessage()    {}
func (*HyperparameterValuesConfigBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5245ebe5bb808c, []int{2}
}

func (m *HyperparameterValuesConfigBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HyperparameterValuesConfigBlob.Unmarshal(m, b)
}
func (m *HyperparameterValuesConfigBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HyperparameterValuesConfigBlob.Marshal(b, m, deterministic)
}
func (m *HyperparameterValuesConfigBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HyperparameterValuesConfigBlob.Merge(m, src)
}
func (m *HyperparameterValuesConfigBlob) XXX_Size() int {
	return xxx_messageInfo_HyperparameterValuesConfigBlob.Size(m)
}
func (m *HyperparameterValuesConfigBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_HyperparameterValuesConfigBlob.DiscardUnknown(m)
}

var xxx_messageInfo_HyperparameterValuesConfigBlob proto.InternalMessageInfo

type isHyperparameterValuesConfigBlob_Value interface {
	isHyperparameterValuesConfigBlob_Value()
}

type HyperparameterValuesConfigBlob_IntValue struct {
	IntValue int64 `protobuf:"varint,1,opt,name=int_value,json=intValue,proto3,oneof"`
}

type HyperparameterValuesConfigBlob_FloatValue struct {
	FloatValue float32 `protobuf:"fixed32,2,opt,name=float_value,json=floatValue,proto3,oneof"`
}

type HyperparameterValuesConfigBlob_StringValue struct {
	StringValue string `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*HyperparameterValuesConfigBlob_IntValue) isHyperparameterValuesConfigBlob_Value() {}

func (*HyperparameterValuesConfigBlob_FloatValue) isHyperparameterValuesConfigBlob_Value() {}

func (*HyperparameterValuesConfigBlob_StringValue) isHyperparameterValuesConfigBlob_Value() {}

func (m *HyperparameterValuesConfigBlob) GetValue() isHyperparameterValuesConfigBlob_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *HyperparameterValuesConfigBlob) GetIntValue() int64 {
	if x, ok := m.GetValue().(*HyperparameterValuesConfigBlob_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *HyperparameterValuesConfigBlob) GetFloatValue() float32 {
	if x, ok := m.GetValue().(*HyperparameterValuesConfigBlob_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

func (m *HyperparameterValuesConfigBlob) GetStringValue() string {
	if x, ok := m.GetValue().(*HyperparameterValuesConfigBlob_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HyperparameterValuesConfigBlob) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HyperparameterValuesConfigBlob_IntValue)(nil),
		(*HyperparameterValuesConfigBlob_FloatValue)(nil),
		(*HyperparameterValuesConfigBlob_StringValue)(nil),
	}
}

type HyperparameterSetConfigBlob struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//	*HyperparameterSetConfigBlob_Continuous
	//	*HyperparameterSetConfigBlob_Discrete
	Value                isHyperparameterSetConfigBlob_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *HyperparameterSetConfigBlob) Reset()         { *m = HyperparameterSetConfigBlob{} }
func (m *HyperparameterSetConfigBlob) String() string { return proto.CompactTextString(m) }
func (*HyperparameterSetConfigBlob) ProtoMessage()    {}
func (*HyperparameterSetConfigBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5245ebe5bb808c, []int{3}
}

func (m *HyperparameterSetConfigBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HyperparameterSetConfigBlob.Unmarshal(m, b)
}
func (m *HyperparameterSetConfigBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HyperparameterSetConfigBlob.Marshal(b, m, deterministic)
}
func (m *HyperparameterSetConfigBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HyperparameterSetConfigBlob.Merge(m, src)
}
func (m *HyperparameterSetConfigBlob) XXX_Size() int {
	return xxx_messageInfo_HyperparameterSetConfigBlob.Size(m)
}
func (m *HyperparameterSetConfigBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_HyperparameterSetConfigBlob.DiscardUnknown(m)
}

var xxx_messageInfo_HyperparameterSetConfigBlob proto.InternalMessageInfo

func (m *HyperparameterSetConfigBlob) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isHyperparameterSetConfigBlob_Value interface {
	isHyperparameterSetConfigBlob_Value()
}

type HyperparameterSetConfigBlob_Continuous struct {
	Continuous *ContinuousHyperparameterSetConfigBlob `protobuf:"bytes,2,opt,name=continuous,proto3,oneof"`
}

type HyperparameterSetConfigBlob_Discrete struct {
	Discrete *DiscreteHyperparameterSetConfigBlob `protobuf:"bytes,3,opt,name=discrete,proto3,oneof"`
}

func (*HyperparameterSetConfigBlob_Continuous) isHyperparameterSetConfigBlob_Value() {}

func (*HyperparameterSetConfigBlob_Discrete) isHyperparameterSetConfigBlob_Value() {}

func (m *HyperparameterSetConfigBlob) GetValue() isHyperparameterSetConfigBlob_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *HyperparameterSetConfigBlob) GetContinuous() *ContinuousHyperparameterSetConfigBlob {
	if x, ok := m.GetValue().(*HyperparameterSetConfigBlob_Continuous); ok {
		return x.Continuous
	}
	return nil
}

func (m *HyperparameterSetConfigBlob) GetDiscrete() *DiscreteHyperparameterSetConfigBlob {
	if x, ok := m.GetValue().(*HyperparameterSetConfigBlob_Discrete); ok {
		return x.Discrete
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HyperparameterSetConfigBlob) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HyperparameterSetConfigBlob_Continuous)(nil),
		(*HyperparameterSetConfigBlob_Discrete)(nil),
	}
}

type ContinuousHyperparameterSetConfigBlob struct {
	IntervalBegin        *HyperparameterValuesConfigBlob `protobuf:"bytes,2,opt,name=interval_begin,json=intervalBegin,proto3" json:"interval_begin,omitempty"`
	IntervalEnd          *HyperparameterValuesConfigBlob `protobuf:"bytes,3,opt,name=interval_end,json=intervalEnd,proto3" json:"interval_end,omitempty"`
	IntervalStep         *HyperparameterValuesConfigBlob `protobuf:"bytes,4,opt,name=interval_step,json=intervalStep,proto3" json:"interval_step,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ContinuousHyperparameterSetConfigBlob) Reset()         { *m = ContinuousHyperparameterSetConfigBlob{} }
func (m *ContinuousHyperparameterSetConfigBlob) String() string { return proto.CompactTextString(m) }
func (*ContinuousHyperparameterSetConfigBlob) ProtoMessage()    {}
func (*ContinuousHyperparameterSetConfigBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5245ebe5bb808c, []int{4}
}

func (m *ContinuousHyperparameterSetConfigBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContinuousHyperparameterSetConfigBlob.Unmarshal(m, b)
}
func (m *ContinuousHyperparameterSetConfigBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContinuousHyperparameterSetConfigBlob.Marshal(b, m, deterministic)
}
func (m *ContinuousHyperparameterSetConfigBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousHyperparameterSetConfigBlob.Merge(m, src)
}
func (m *ContinuousHyperparameterSetConfigBlob) XXX_Size() int {
	return xxx_messageInfo_ContinuousHyperparameterSetConfigBlob.Size(m)
}
func (m *ContinuousHyperparameterSetConfigBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousHyperparameterSetConfigBlob.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousHyperparameterSetConfigBlob proto.InternalMessageInfo

func (m *ContinuousHyperparameterSetConfigBlob) GetIntervalBegin() *HyperparameterValuesConfigBlob {
	if m != nil {
		return m.IntervalBegin
	}
	return nil
}

func (m *ContinuousHyperparameterSetConfigBlob) GetIntervalEnd() *HyperparameterValuesConfigBlob {
	if m != nil {
		return m.IntervalEnd
	}
	return nil
}

func (m *ContinuousHyperparameterSetConfigBlob) GetIntervalStep() *HyperparameterValuesConfigBlob {
	if m != nil {
		return m.IntervalStep
	}
	return nil
}

type DiscreteHyperparameterSetConfigBlob struct {
	Values               []*HyperparameterValuesConfigBlob `protobuf:"bytes,4,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *DiscreteHyperparameterSetConfigBlob) Reset()         { *m = DiscreteHyperparameterSetConfigBlob{} }
func (m *DiscreteHyperparameterSetConfigBlob) String() string { return proto.CompactTextString(m) }
func (*DiscreteHyperparameterSetConfigBlob) ProtoMessage()    {}
func (*DiscreteHyperparameterSetConfigBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a5245ebe5bb808c, []int{5}
}

func (m *DiscreteHyperparameterSetConfigBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscreteHyperparameterSetConfigBlob.Unmarshal(m, b)
}
func (m *DiscreteHyperparameterSetConfigBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscreteHyperparameterSetConfigBlob.Marshal(b, m, deterministic)
}
func (m *DiscreteHyperparameterSetConfigBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscreteHyperparameterSetConfigBlob.Merge(m, src)
}
func (m *DiscreteHyperparameterSetConfigBlob) XXX_Size() int {
	return xxx_messageInfo_DiscreteHyperparameterSetConfigBlob.Size(m)
}
func (m *DiscreteHyperparameterSetConfigBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscreteHyperparameterSetConfigBlob.DiscardUnknown(m)
}

var xxx_messageInfo_DiscreteHyperparameterSetConfigBlob proto.InternalMessageInfo

func (m *DiscreteHyperparameterSetConfigBlob) GetValues() []*HyperparameterValuesConfigBlob {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigBlob)(nil), "ai.verta.modeldb.versioning.ConfigBlob")
	proto.RegisterType((*HyperparameterConfigBlob)(nil), "ai.verta.modeldb.versioning.HyperparameterConfigBlob")
	proto.RegisterType((*HyperparameterValuesConfigBlob)(nil), "ai.verta.modeldb.versioning.HyperparameterValuesConfigBlob")
	proto.RegisterType((*HyperparameterSetConfigBlob)(nil), "ai.verta.modeldb.versioning.HyperparameterSetConfigBlob")
	proto.RegisterType((*ContinuousHyperparameterSetConfigBlob)(nil), "ai.verta.modeldb.versioning.ContinuousHyperparameterSetConfigBlob")
	proto.RegisterType((*DiscreteHyperparameterSetConfigBlob)(nil), "ai.verta.modeldb.versioning.DiscreteHyperparameterSetConfigBlob")
}

func init() {
	proto.RegisterFile("protos/public/modeldb/versioning/Config.proto", fileDescriptor_0a5245ebe5bb808c)
}

var fileDescriptor_0a5245ebe5bb808c = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0xd9, 0xb5, 0xb6, 0x2f, 0x55, 0x71, 0x4e, 0x81, 0xa2, 0xac, 0x29, 0xc2, 0x5e,
	0x4c, 0x60, 0x45, 0x10, 0xbc, 0x68, 0xaa, 0xb0, 0x3d, 0x08, 0x9a, 0x85, 0x1e, 0x3c, 0x34, 0x26,
	0x9b, 0xd7, 0xd9, 0x81, 0xec, 0x4c, 0x98, 0x99, 0x2c, 0xe8, 0x49, 0xbf, 0x81, 0x5f, 0xcc, 0xab,
	0x9f, 0x47, 0x32, 0x99, 0xcd, 0xb6, 0x45, 0x63, 0xcb, 0x7a, 0x0b, 0x2f, 0xbf, 0xf9, 0xff, 0xff,
	0xef, 0x4d, 0xf2, 0xe0, 0x59, 0x25, 0x85, 0x16, 0x2a, 0xaa, 0xea, 0xbc, 0x64, 0x8b, 0x68, 0x25,
	0x0a, 0x2c, 0x8b, 0x3c, 0x5a, 0xa3, 0x54, 0x4c, 0x70, 0xc6, 0x69, 0x74, 0x22, 0xf8, 0x05, 0xa3,
	0xa1, 0xe1, 0xc8, 0x51, 0xc6, 0xc2, 0x35, 0x4a, 0x9d, 0x85, 0x96, 0x0c, 0xb7, 0x64, 0xf0, 0xcb,
	0x01, 0x68, 0xe9, 0xb8, 0x14, 0x39, 0xa1, 0x40, 0x96, 0x5f, 0x2a, 0x94, 0x55, 0x26, 0xb3, 0x15,
	0x6a, 0x94, 0xa9, 0x42, 0xed, 0x3b, 0xe3, 0xe1, 0xc4, 0x9b, 0xbe, 0x0c, 0x7b, 0x84, 0xc2, 0xd9,
	0x95, 0x63, 0x73, 0xd4, 0x5b, 0xd5, 0xe4, 0xe1, 0xf2, 0xfa, 0x4b, 0x92, 0xc2, 0x83, 0xab, 0x45,
	0xe5, 0xbb, 0xc6, 0xe5, 0xc5, 0x2d, 0x5c, 0x2e, 0x59, 0x5c, 0x57, 0x0b, 0xbe, 0x3b, 0xe0, 0xff,
	0x8d, 0x26, 0x04, 0x46, 0x3c, 0x5b, 0xa1, 0xef, 0x8c, 0x9d, 0xc9, 0x41, 0x62, 0x9e, 0xc9, 0x47,
	0xb8, 0xb3, 0xce, 0xca, 0x1a, 0x7d, 0x77, 0xec, 0x4c, 0xbc, 0xe9, 0xab, 0x5b, 0xe4, 0x38, 0x6b,
	0xce, 0xa9, 0x4b, 0x69, 0x5a, 0xa5, 0xe0, 0x87, 0x03, 0x8f, 0xfb, 0x49, 0xf2, 0x08, 0x0e, 0x18,
	0xd7, 0x69, 0xeb, 0xdc, 0xc4, 0x19, 0xce, 0x06, 0xc9, 0x3e, 0xe3, 0xda, 0x90, 0xe4, 0x09, 0x78,
	0x17, 0xa5, 0xc8, 0x36, 0x40, 0x13, 0xcd, 0x9d, 0x0d, 0x12, 0x30, 0xc5, 0x16, 0x39, 0x86, 0x43,
	0xa5, 0x25, 0xe3, 0xd4, 0x32, 0xc3, 0xa6, 0xa7, 0xd9, 0x20, 0xf1, 0xda, 0xaa, 0x81, 0xe2, 0xbb,
	0xb6, 0xb9, 0xe0, 0x9b, 0x0b, 0x47, 0x3d, 0x57, 0xf5, 0xc7, 0xc9, 0x14, 0x00, 0x0b, 0xc1, 0x35,
	0xe3, 0xb5, 0xa8, 0x95, 0x1d, 0x4f, 0xdc, 0x3b, 0x9e, 0x93, 0x0e, 0xef, 0xf1, 0x6a, 0xfa, 0xd8,
	0xea, 0x92, 0x73, 0xd8, 0x2f, 0x98, 0x5a, 0x48, 0xd4, 0x6d, 0x0f, 0xde, 0xf4, 0x75, 0xaf, 0xc7,
	0x5b, 0x0b, 0xf7, 0x3b, 0x74, 0x9a, 0xdb, 0x11, 0xfc, 0x74, 0xe1, 0xe9, 0x8d, 0x02, 0x92, 0x1c,
	0xee, 0x33, 0xae, 0x51, 0xae, 0xb3, 0x32, 0xcd, 0x91, 0x32, 0xfe, 0x3f, 0xbe, 0x8d, 0x7b, 0x1b,
	0xc9, 0xb8, 0x51, 0x24, 0xe7, 0x70, 0xd8, 0x79, 0x20, 0x2f, 0x6c, 0xeb, 0x3b, 0x39, 0x78, 0x1b,
	0xc1, 0x77, 0xbc, 0x20, 0x9f, 0xa1, 0x33, 0x4c, 0x95, 0xc6, 0xca, 0x1f, 0xed, 0x6e, 0xd0, 0x25,
	0x9e, 0x6b, 0xac, 0x82, 0xaf, 0x70, 0x7c, 0x83, 0xbb, 0x20, 0x73, 0xd8, 0x33, 0xf3, 0x57, 0xfe,
	0xc8, 0xfc, 0xe8, 0x3b, 0x25, 0xb0, 0x52, 0xf1, 0xfb, 0x0f, 0xce, 0xa7, 0x53, 0xca, 0xf4, 0xb2,
	0xce, 0xc3, 0x85, 0x58, 0x45, 0x67, 0x8d, 0xde, 0x9b, 0xd3, 0x6e, 0x27, 0xda, 0x4d, 0x49, 0x91,
	0x47, 0x54, 0x44, 0xff, 0xda, 0x9b, 0xf9, 0x9e, 0x21, 0x9e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x19, 0x0b, 0xd3, 0x19, 0x62, 0x05, 0x00, 0x00,
}
