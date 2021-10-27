package main

import (
	"context"
	"github.com/mangenotwork/extras/apps/GrpcClient/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("192.168.0.9:1232", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewMessageRPCClient(conn)
	r, err := c.WhiteWordAdd(context.Background(), &proto.WhiteWordAddReq{
		Word: "飞机",
	})
	log.Println(r, err)
	for i:=0; i<2; i++ {
		//r1, err := c.Do(context.Background(), &proto.DoReq{Str: "你是个废品你个狗日的你知道吗", Sub: "*"})
		r1, err := c.Do(context.Background(), &proto.DoReq{Str: "打飞机,我在路口交通进爱圣诞节在欧帕斯卡分速度发完全二维卡；〔〕看好小卡卡是爬山分明；；发送爱啥啥的加拉收到啦行口交就在这个dsf收到了几个了路口交接", Sub: "*"})
		log.Println(r1, err)
	}

	r2, err := c.IsHaveList(context.Background(), &proto.IsHaveListReq{
		Str: "我在口交通进爱操好了你妈圣诞节在欧帕斯卡分废速度发完全二维卡",
	})
	log.Println(r2, r2.IsHave, r2.List, err)
}
