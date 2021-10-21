package main

import (
	"context"
	"github.com/mangenotwork/extras/apps/GrpcClient/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("192.168.0.9:1211", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewMessageRPCClient(conn)
	r, err := c.Get(context.Background(), &proto.GetReq{})
	log.Println(r, err)
	for i:=0; i<100; i++ {
		r1, err := c.Do(context.Background(), &proto.DoReq{Str: "你是个废物你知道吗", Sub: "*"})
		log.Println(r1, err)
	}

}
