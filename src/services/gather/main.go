package main

import (
	"context"
	"errors"
	"strconv"

	"github.com/slory7/angulargo/src/infrastructure/app"
	"github.com/slory7/angulargo/src/infrastructure/framework/net/httpclient"
	"github.com/slory7/angulargo/src/proto/gather"
	"github.com/slory7/angulargo/src/services"

	"github.com/nuveo/log"

	_ "github.com/asim/go-micro/plugins/registry/etcd/v3"
	micro "github.com/asim/go-micro/v3"
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
	services.PrintTrace(ctx, "GetHttpContent")

	traceID, _ := services.GetTrace(ctx)
	httpClient := app.GetIoCInstanceMust[httpclient.IHttpClient]()

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
