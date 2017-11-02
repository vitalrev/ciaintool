// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interface.proto

/*
Package meta is a generated protocol buffer package.

It is generated from these files:
	interface.proto

It has these top-level messages:
	InterfaceDescriptor
	Interfaces
	GetInterfacesParams
	GetInterfaceParams
	GetFactsParams
	Facts
*/
package meta

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

type InterfaceDescriptor struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *InterfaceDescriptor) Reset()                    { *m = InterfaceDescriptor{} }
func (m *InterfaceDescriptor) String() string            { return proto.CompactTextString(m) }
func (*InterfaceDescriptor) ProtoMessage()               {}
func (*InterfaceDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InterfaceDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfaceDescriptor) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Interfaces struct {
	Descriptors []*InterfaceDescriptor `protobuf:"bytes,1,rep,name=descriptors" json:"descriptors,omitempty"`
}

func (m *Interfaces) Reset()                    { *m = Interfaces{} }
func (m *Interfaces) String() string            { return proto.CompactTextString(m) }
func (*Interfaces) ProtoMessage()               {}
func (*Interfaces) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Interfaces) GetDescriptors() []*InterfaceDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

type GetInterfacesParams struct {
	IncludeContent bool `protobuf:"varint,1,opt,name=IncludeContent" json:"IncludeContent,omitempty"`
}

func (m *GetInterfacesParams) Reset()                    { *m = GetInterfacesParams{} }
func (m *GetInterfacesParams) String() string            { return proto.CompactTextString(m) }
func (*GetInterfacesParams) ProtoMessage()               {}
func (*GetInterfacesParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetInterfacesParams) GetIncludeContent() bool {
	if m != nil {
		return m.IncludeContent
	}
	return false
}

type GetInterfaceParams struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GetInterfaceParams) Reset()                    { *m = GetInterfaceParams{} }
func (m *GetInterfaceParams) String() string            { return proto.CompactTextString(m) }
func (*GetInterfaceParams) ProtoMessage()               {}
func (*GetInterfaceParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetInterfaceParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetFactsParams struct {
}

func (m *GetFactsParams) Reset()                    { *m = GetFactsParams{} }
func (m *GetFactsParams) String() string            { return proto.CompactTextString(m) }
func (*GetFactsParams) ProtoMessage()               {}
func (*GetFactsParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Facts struct {
	Facts []*Facts_Fact `protobuf:"bytes,1,rep,name=facts" json:"facts,omitempty"`
}

func (m *Facts) Reset()                    { *m = Facts{} }
func (m *Facts) String() string            { return proto.CompactTextString(m) }
func (*Facts) ProtoMessage()               {}
func (*Facts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Facts) GetFacts() []*Facts_Fact {
	if m != nil {
		return m.Facts
	}
	return nil
}

type Facts_Fact struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Facts_Fact) Reset()                    { *m = Facts_Fact{} }
func (m *Facts_Fact) String() string            { return proto.CompactTextString(m) }
func (*Facts_Fact) ProtoMessage()               {}
func (*Facts_Fact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *Facts_Fact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Facts_Fact) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*InterfaceDescriptor)(nil), "meta.InterfaceDescriptor")
	proto.RegisterType((*Interfaces)(nil), "meta.Interfaces")
	proto.RegisterType((*GetInterfacesParams)(nil), "meta.GetInterfacesParams")
	proto.RegisterType((*GetInterfaceParams)(nil), "meta.GetInterfaceParams")
	proto.RegisterType((*GetFactsParams)(nil), "meta.GetFactsParams")
	proto.RegisterType((*Facts)(nil), "meta.Facts")
	proto.RegisterType((*Facts_Fact)(nil), "meta.Facts.Fact")
}

func init() { proto.RegisterFile("interface.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x50, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0x9a, 0x88, 0x9d, 0x4a, 0x2d, 0xab, 0x87, 0xe8, 0x29, 0xec, 0xa1, 0xe4, 0x14, 0x44,
	0x8f, 0xe2, 0x49, 0xb1, 0xe4, 0x26, 0xfb, 0x07, 0xe3, 0x66, 0x0a, 0x85, 0x66, 0xb7, 0xec, 0x4e,
	0xfd, 0x7e, 0xd9, 0x49, 0xac, 0xa5, 0xe4, 0x12, 0xe6, 0xcd, 0xcb, 0x9b, 0xf7, 0xf6, 0xc1, 0xed,
	0xd6, 0x31, 0x85, 0x0d, 0x5a, 0x6a, 0xf6, 0xc1, 0xb3, 0x57, 0x79, 0x4f, 0x8c, 0xfa, 0x0d, 0xee,
	0xda, 0x3f, 0xe2, 0x83, 0xa2, 0x0d, 0xdb, 0x3d, 0xfb, 0xa0, 0x14, 0xe4, 0x0e, 0x7b, 0x2a, 0xb3,
	0x2a, 0xab, 0x67, 0x46, 0xe6, 0xb4, 0xeb, 0x90, 0xb1, 0xbc, 0xa8, 0xb2, 0xfa, 0xc6, 0xc8, 0xac,
	0x5b, 0x80, 0xa3, 0x3c, 0xaa, 0x57, 0x98, 0x77, 0xc7, 0x1b, 0xb1, 0xcc, 0xaa, 0xcb, 0x7a, 0xfe,
	0xfc, 0xd0, 0x24, 0xa3, 0x66, 0xc2, 0xc5, 0x9c, 0xfe, 0x9d, 0x92, 0xac, 0x89, 0xff, 0xaf, 0x7d,
	0x61, 0xc0, 0x3e, 0xaa, 0x15, 0x2c, 0x5a, 0x67, 0x77, 0x87, 0x8e, 0xde, 0xbd, 0x63, 0x72, 0x2c,
	0x99, 0xae, 0xcd, 0xd9, 0x56, 0xd7, 0xa0, 0x4e, 0xe5, 0xa3, 0x7a, 0xe2, 0x1d, 0x7a, 0x09, 0x8b,
	0x35, 0xf1, 0x27, 0x5a, 0x1e, 0x3d, 0x34, 0x42, 0x21, 0x50, 0xad, 0xa0, 0xd8, 0xa4, 0x61, 0x8c,
	0xbe, 0x1c, 0xa2, 0x0b, 0x27, 0x5f, 0x33, 0xd0, 0x8f, 0x4f, 0x90, 0x27, 0x38, 0x59, 0xd3, 0x3d,
	0x14, 0x3f, 0xb8, 0x3b, 0x90, 0xf4, 0x34, 0x33, 0x03, 0xf8, 0xbe, 0x92, 0xd2, 0x5f, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x27, 0x6e, 0x7c, 0x4e, 0x87, 0x01, 0x00, 0x00,
}
