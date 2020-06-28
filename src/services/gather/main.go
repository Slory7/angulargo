package main

import (
	"context"
	"errors"
	"strconv"

	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"
	gather "github.com/slory7/angulargo/src/services/gather/proto"

	"github.com/nuveo/log"
	"golang.org/x/net/trace"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/metadata"
)

func main() {
	service := micro.NewService(
		micro.Name("angulargo.micro.srv.gather"),
	)

	service.Init()

	gather.RegisterGatherHandler(service.Server(), &GatherSrv{})

	service.Run()
}

type GatherSrv struct {
}

func (s *GatherSrv) GetHttpContent(ctx context.Context, req *gather.Request, rsp *gather.Result) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["Traceid"]

	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("fromName %s", md["Fromname"])
		tr.LazyPrintf("traceID %s", traceID)
	}
	log.Printf("fromName %s\n", md["Fromname"])
	log.Printf("traceID %s\n", traceID)

	result, err := httpclient.HttpSend(req.BaseUrl, req.RelativeUrl, req.UrlParams, req.Headers, req.ContentType, req.Method, req.PostData, httpclient.TokenEmpty, traceID, false, req.TimeOut)

	if err != nil {
		return err
	}
	if !result.IsSuccess {
		return errors.New(result.Message + ":" + strconv.Itoa(result.StatusCode))
	}

	rsp.Content = result.Content

	return nil
}
