// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: api/api.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	trending "github.com/slory7/angulargo/src/proto/trending"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for TrendingSrv service

func NewTrendingSrvEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TrendingSrv service

type TrendingSrvService interface {
	GetGithubTrending(ctx context.Context, in *trending.Request, opts ...client.CallOption) (*trending.GithubTrendingInfo, error)
}

type trendingSrvService struct {
	c    client.Client
	name string
}

func NewTrendingSrvService(name string, c client.Client) TrendingSrvService {
	return &trendingSrvService{
		c:    c,
		name: name,
	}
}

func (c *trendingSrvService) GetGithubTrending(ctx context.Context, in *trending.Request, opts ...client.CallOption) (*trending.GithubTrendingInfo, error) {
	req := c.c.NewRequest(c.name, "TrendingSrv.GetGithubTrending", in)
	out := new(trending.GithubTrendingInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TrendingSrv service

type TrendingSrvHandler interface {
	GetGithubTrending(context.Context, *trending.Request, *trending.GithubTrendingInfo) error
}

func RegisterTrendingSrvHandler(s server.Server, hdlr TrendingSrvHandler, opts ...server.HandlerOption) error {
	type trendingSrv interface {
		GetGithubTrending(ctx context.Context, in *trending.Request, out *trending.GithubTrendingInfo) error
	}
	type TrendingSrv struct {
		trendingSrv
	}
	h := &trendingSrvHandler{hdlr}
	return s.Handle(s.NewHandler(&TrendingSrv{h}, opts...))
}

type trendingSrvHandler struct {
	TrendingSrvHandler
}

func (h *trendingSrvHandler) GetGithubTrending(ctx context.Context, in *trending.Request, out *trending.GithubTrendingInfo) error {
	return h.TrendingSrvHandler.GetGithubTrending(ctx, in, out)
}
