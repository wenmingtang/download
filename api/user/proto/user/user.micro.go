// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/user/proto/user/user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	api/user/proto/user/user.proto

It has these top-level messages:
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import account "github.com/wenmingtang/download/srv/account/proto/account"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = account.User{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for User service

type UserService interface {
	Create(ctx context.Context, in *account.User, opts ...client.CallOption) (*account.User, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "user"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *account.User, opts ...client.CallOption) (*account.User, error) {
	req := c.c.NewRequest(c.name, "User.Create", in)
	out := new(account.User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Create(context.Context, *account.User, *account.User) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Create(ctx context.Context, in *account.User, out *account.User) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Create(ctx context.Context, in *account.User, out *account.User) error {
	return h.UserHandler.Create(ctx, in, out)
}
