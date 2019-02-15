// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/trending.proto

/*
Package trending is a generated protocol buffer package.

It is generated from these files:
	proto/trending.proto

It has these top-level messages:
	Request
	GithubTrendingInfo
	GitRepo
*/
package trending

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Client API for Trending service

type TrendingService interface {
	GetGithubTrending(ctx context.Context, in *Request, opts ...client.CallOption) (*GithubTrendingInfo, error)
	GetAndSaveGithubTrending(ctx context.Context, in *Request, opts ...client.CallOption) (*GithubTrendingInfo, error)
}

type trendingService struct {
	c    client.Client
	name string
}

func NewTrendingService(name string, c client.Client) TrendingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "trending"
	}
	return &trendingService{
		c:    c,
		name: name,
	}
}

func (c *trendingService) GetGithubTrending(ctx context.Context, in *Request, opts ...client.CallOption) (*GithubTrendingInfo, error) {
	req := c.c.NewRequest(c.name, "Trending.GetGithubTrending", in)
	out := new(GithubTrendingInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *trendingService) GetAndSaveGithubTrending(ctx context.Context, in *Request, opts ...client.CallOption) (*GithubTrendingInfo, error) {
	req := c.c.NewRequest(c.name, "Trending.GetAndSaveGithubTrending", in)
	out := new(GithubTrendingInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Trending service

type TrendingHandler interface {
	GetGithubTrending(context.Context, *Request, *GithubTrendingInfo) error
	GetAndSaveGithubTrending(context.Context, *Request, *GithubTrendingInfo) error
}

func RegisterTrendingHandler(s server.Server, hdlr TrendingHandler, opts ...server.HandlerOption) error {
	type trending interface {
		GetGithubTrending(ctx context.Context, in *Request, out *GithubTrendingInfo) error
		GetAndSaveGithubTrending(ctx context.Context, in *Request, out *GithubTrendingInfo) error
	}
	type Trending struct {
		trending
	}
	h := &trendingHandler{hdlr}
	return s.Handle(s.NewHandler(&Trending{h}, opts...))
}

type trendingHandler struct {
	TrendingHandler
}

func (h *trendingHandler) GetGithubTrending(ctx context.Context, in *Request, out *GithubTrendingInfo) error {
	return h.TrendingHandler.GetGithubTrending(ctx, in, out)
}

func (h *trendingHandler) GetAndSaveGithubTrending(ctx context.Context, in *Request, out *GithubTrendingInfo) error {
	return h.TrendingHandler.GetAndSaveGithubTrending(ctx, in, out)
}
