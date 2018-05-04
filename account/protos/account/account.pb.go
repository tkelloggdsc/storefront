// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package go_micro_srv_accounts is a generated protocol buffer package.

It is generated from these files:
	account.proto

It has these top-level messages:
	GetAccountRequest
	Account
*/
package go_micro_srv_accounts

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

type Account_Type int32

const (
	Account_ADMIN    Account_Type = 0
	Account_CUSTOMER Account_Type = 1
	Account_GUEST    Account_Type = 2
)

var Account_Type_name = map[int32]string{
	0: "ADMIN",
	1: "CUSTOMER",
	2: "GUEST",
}
var Account_Type_value = map[string]int32{
	"ADMIN":    0,
	"CUSTOMER": 1,
	"GUEST":    2,
}

func (x Account_Type) String() string {
	return proto.EnumName(Account_Type_name, int32(x))
}
func (Account_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type GetAccountRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetAccountRequest) Reset()                    { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()               {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetAccountRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Account struct {
	Guid     string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Account) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccountRequest)(nil), "go.micro.srv.accounts.GetAccountRequest")
	proto.RegisterType((*Account)(nil), "go.micro.srv.accounts.Account")
	proto.RegisterEnum("go.micro.srv.accounts.Account_Type", Account_Type_name, Account_Type_value)
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0xcf, 0xd7, 0xcb, 0xcd,
	0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x83, 0xca, 0x15, 0x2b, 0x29, 0x73, 0x09, 0xba,
	0xa7, 0x96, 0x38, 0x42, 0xb8, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c,
	0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0xfd, 0x8c, 0x5c,
	0xec, 0x50, 0x25, 0x42, 0x42, 0x5c, 0x2c, 0xe9, 0xa5, 0x70, 0x59, 0x30, 0x1b, 0x24, 0x96, 0x97,
	0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x11, 0x03, 0xb1, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x73, 0x13, 0x33,
	0x73, 0x24, 0x98, 0xc1, 0x82, 0x10, 0x8e, 0x90, 0x14, 0x17, 0x47, 0x41, 0x62, 0x71, 0x71, 0x79,
	0x7e, 0x51, 0x8a, 0x04, 0x0b, 0x58, 0x02, 0xce, 0x57, 0xd2, 0xe2, 0x62, 0x09, 0xa9, 0x2c, 0x48,
	0x15, 0xe2, 0xe4, 0x62, 0x75, 0x74, 0xf1, 0xf5, 0xf4, 0x13, 0x60, 0x10, 0xe2, 0xe1, 0xe2, 0x70,
	0x0e, 0x0d, 0x0e, 0xf1, 0xf7, 0x75, 0x0d, 0x12, 0x60, 0x04, 0x49, 0xb8, 0x87, 0xba, 0x06, 0x87,
	0x08, 0x30, 0x19, 0x6d, 0x67, 0xe4, 0xe2, 0x80, 0xba, 0xa8, 0x58, 0x28, 0x82, 0x8b, 0x0b, 0xe1,
	0x07, 0x21, 0x0d, 0x3d, 0xac, 0x3e, 0xd5, 0xc3, 0xf0, 0xa6, 0x94, 0x1c, 0x0e, 0x95, 0x50, 0x65,
	0x4a, 0x0c, 0x42, 0x81, 0x5c, 0xbc, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x30, 0xc3, 0x09, 0x68,
	0x21, 0x6c, 0x64, 0x12, 0x1b, 0x38, 0x3a, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0xc4,
	0x0d, 0x2a, 0x9f, 0x01, 0x00, 0x00,
}