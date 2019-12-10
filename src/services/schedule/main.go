package main

import (
	"context"
	trending "services/trending/proto"

	"github.com/slory7/angulargo/src/services/infrastructure/config"
	"github.com/slory7/angulargo/src/services/infrastructure/framework/globals"
	"time"

	"github.com/nuveo/log"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/trace"
)

func main() {
	log.Println("Schedule is running...")

	glbConfig = config.GetConfig(globals.GetEnvironment(), &Config{}).(*Config)

	log.Println("Run GetherSchedule first")
	GetherSchedule()

	tick := time.NewTicker(time.Second * time.Duration(glbConfig.IntervalSeconds))
	for range tick.C {
		log.Println("Run GetherSchedule scheduled")
		go GetherSchedule()
	}
}

func GetherSchedule() {
	tr := trace.New("schedule.v1", "Gather")
	defer tr.Finish()

	ctx := context.TODO()

	md := metadata.Metadata{}

	if traceID, err := uuid.NewV4(); err == nil {
		md["Traceid"] = traceID.String()
		md["Fromname"] = "schedule.v1"
		ctx = metadata.NewContext(ctx, md)

		log.Printf("traceID %s\n", traceID)
	}

	client := client.DefaultClient

	trendingClient := trending.NewTrendingService("angulargo.micro.srv.trending", client)
	req := &trending.Request{}
	result, err := trendingClient.GetAndSaveGithubTrending(ctx, req)
	if err != nil {
		log.Errorf("get trending error:%v\n", err)
		return
	}
	log.Println(result.Title)
}
