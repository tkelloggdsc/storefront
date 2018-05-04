// Code generated by protoc-gen-micro. DO NOT EDIT.
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

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Accounts service

type AccountsService interface {
	GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*Account, error)
	CreateAccount(ctx context.Context, in *Account, opts ...client.CallOption) (*Account, error)
}

type accountsService struct {
	c           client.Client
	serviceName string
}

func AccountsServiceClient(serviceName string, c client.Client) AccountsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.accounts"
	}
	return &accountsService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountsService) GetAccount(ctx context.Context, in *GetAccountRequest, opts ...client.CallOption) (*Account, error) {
	req := c.c.NewRequest(c.serviceName, "Accounts.GetAccount", in)
	out := new(Account)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsService) CreateAccount(ctx context.Context, in *Account, opts ...client.CallOption) (*Account, error) {
	req := c.c.NewRequest(c.serviceName, "Accounts.CreateAccount", in)
	out := new(Account)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Accounts service

type AccountsHandler interface {
	GetAccount(context.Context, *GetAccountRequest, *Account) error
	CreateAccount(context.Context, *Account, *Account) error
}

func RegisterAccountsHandler(s server.Server, hdlr AccountsHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Accounts{hdlr}, opts...))
}

type Accounts struct {
	AccountsHandler
}

func (h *Accounts) GetAccount(ctx context.Context, in *GetAccountRequest, out *Account) error {
	return h.AccountsHandler.GetAccount(ctx, in, out)
}

func (h *Accounts) CreateAccount(ctx context.Context, in *Account, out *Account) error {
	return h.AccountsHandler.CreateAccount(ctx, in, out)
}