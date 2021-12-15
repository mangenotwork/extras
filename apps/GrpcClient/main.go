package main

import (
	"github.com/mangenotwork/extras/apps/GrpcClient/proto"
	"github.com/mangenotwork/extras/common/grpc"
	"log"
)

func main() {
	client, err := grpc.NewClient(grpc.ClientArg{
		ServiceAddr: "192.168.0.9:11232",
		ServiceName: "BlockWord",
	})
	if err != nil {
		log.Println(err)
		return
	}
	conn, ctx, err := client.Conn()
	log.Print("conn = ", conn)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	c := proto.NewMessageRPCClient(conn)
	r, err := c.WhiteWordAdd(ctx, &proto.WhiteWordAddReq{
		Word: "飞机",
	})
	log.Println(r, err)
	for i:=0; i<2; i++ {
		//r1, err := c.Do(context.Background(), &proto.DoReq{Str: "你是个废品你个狗日的你知道吗", Sub: "*"})
		r1, err := c.Do(ctx, &proto.DoReq{Str: "打飞机,我在路口交通进爱圣诞节在欧帕斯卡分速度发完全二维卡；〔〕看好小卡卡是爬山分明；；发送爱啥啥的加拉收到啦行口交就在这个dsf收到了几个了路口交接", Sub: "*"})
		log.Println(r1, err)
	}

	r2, err := c.IsHaveList(ctx, &proto.IsHaveListReq{
		Str: "我在口交通进爱操好了你妈圣诞节在欧帕斯卡分废速度发完全二维卡",
	})
	log.Println(r2, r2.IsHave, r2.List, err)
}
