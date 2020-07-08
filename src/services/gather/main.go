package main

import (
	"context"
	"errors"
	"strconv"

	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"
	"github.com/slory7/angulargo/src/services"
	gather "github.com/slory7/angulargo/src/services/gather/proto"

	"github.com/nuveo/log"
	"golang.org/x/net/trace"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
)

func main() {
	log.Printf("%s is running...\n", services.ServiceNameGather)

	app.InitAppInstance(nil)
	app.Instance.RegisterIoC(nil)

	service := micro.NewService(
		micro.Name(services.ServiceNameGather),
	)

	service.Init()

	err := gather.RegisterGatherHandler(service.Server(), &GatherSrv{})

	if err != nil {
		log.Fatal(err)
	}
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

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

	httpClient := app.Instance.GetIoCInstanceMust((*httpclient.IHttpClient)(nil)).(httpclient.IHttpClient)
	result, err := httpClient.HttpSend(req.BaseUrl, req.RelativeUrl, req.UrlParams, req.Headers, req.ContentType, req.Method, req.PostData, httpclient.TokenEmpty, traceID, false, req.TimeOut)

	if err != nil {
		return err
	}
	if !result.IsSuccess {
		return errors.New(result.Message + ":" + strconv.Itoa(result.StatusCode))
	}

	rsp.Content = result.Content

	return nil
}
