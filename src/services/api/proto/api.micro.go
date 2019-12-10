// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	proto/api.proto

It has these top-level messages:
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import trending "github.com/slory7/angulargo/src/services/trending/proto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = trending.GithubTrendingInfo{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for API service

type APIService interface {
	GetGithubTrending(ctx context.Context, in *trending.Request, opts ...client.CallOption) (*trending.GithubTrendingInfo, error)
}

type aPIService struct {
	c    client.Client
	name string
}

func NewAPIService(name string, c client.Client) APIService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "api"
	}
	return &aPIService{
		c:    c,
		name: name,
	}
}

func (c *aPIService) GetGithubTrending(ctx context.Context, in *trending.Request, opts ...client.CallOption) (*trending.GithubTrendingInfo, error) {
	req := c.c.NewRequest(c.name, "API.GetGithubTrending", in)
	out := new(trending.GithubTrendingInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIHandler interface {
	GetGithubTrending(context.Context, *trending.Request, *trending.GithubTrendingInfo) error
}

func RegisterAPIHandler(s server.Server, hdlr APIHandler, opts ...server.HandlerOption) error {
	type aPI interface {
		GetGithubTrending(ctx context.Context, in *trending.Request, out *trending.GithubTrendingInfo) error
	}
	type API struct {
		aPI
	}
	h := &aPIHandler{hdlr}
	return s.Handle(s.NewHandler(&API{h}, opts...))
}

type aPIHandler struct {
	APIHandler
}

func (h *aPIHandler) GetGithubTrending(ctx context.Context, in *trending.Request, out *trending.GithubTrendingInfo) error {
	return h.APIHandler.GetGithubTrending(ctx, in, out)
}
