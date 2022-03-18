package main

import (
	"context"
	"log"

	"github.com/slory7/angulargo/src/proto/api"
	"github.com/slory7/angulargo/src/proto/trending"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("server.crt", "")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := grpc.Dial("docker.data:8088", grpc.WithTransportCredentials(creds))
	if err == nil {
		c := api.NewTrendingSrvClient(conn)
		info, err := c.GetGithubTrending(context.Background(), &trending.Request{Title: "tt"})
		log.Print(err)
		log.Println(info)
	} else {
		log.Print(err)
	}
	defer conn.Close()

}
