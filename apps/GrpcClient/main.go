package main

import (
	blockword "github.com/mangenotwork/extras/apps/GrpcClient/blockword/proto"
	wordhelper "github.com/mangenotwork/extras/apps/GrpcClient/wordhelper/proto"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/grpc"
	"github.com/mangenotwork/extras/common/logger"
)

func main() {
	conf.InitConf()
	logger.InitLogger()

	//BlockWord()
	WordHelper()

}


func BlockWord() {
	client, err := grpc.NewClient(grpc.ClientArg{
		ServiceAddr: "192.168.0.9:11232",
		ServiceName: "BlockWord",
	})
	if err != nil {
		logger.Error(err)
		return
	}
	conn, ctx, err := client.Conn()
	logger.Info("conn = ", conn)
	defer conn.Close()
	if err != nil {
		logger.Error(err)
		return
	}

	c := blockword.NewBlockwordRPCClient(conn)
	r, err := c.WhiteWordAdd(ctx, &blockword.WhiteWordAddReq{
		Word: "飞机",
	})
	logger.Info(r, err)
	for i:=0; i<2; i++ {
		//r1, err := c.Do(context.Background(), &proto.DoReq{Str: "你是个废品你个狗日的你知道吗", Sub: "*"})
		r1, err := c.Do(ctx, &blockword.DoReq{Str: "打飞机,我在路口交通进爱圣诞节在欧帕斯卡分速度发完全二维卡；〔〕看好小卡卡是爬山分明；；发送爱啥啥的加拉收到啦行口交就在这个dsf收到了几个了路口交接", Sub: "*"})
		logger.Info(r1, err)
	}

	r2, err := c.IsHaveList(ctx, &blockword.IsHaveListReq{
		Str: "我在口交通进爱操好了你妈圣诞节在欧帕斯卡分废速度发完全二维卡",
	})
	logger.Info(r2, r2.IsHave, r2.List, err)
}

func WordHelper() {
	client, err := grpc.NewClient(grpc.ClientArg{
		ServiceAddr: "192.168.0.9:11252",
		ServiceName: "WordHelper",
	})
	if err != nil {
		logger.Error(err)
		return
	}
	conn, ctx, err := client.Conn()
	logger.Info("conn = ", conn)
	defer conn.Close()
	if err != nil {
		logger.Error(err)
		return
	}

	c := wordhelper.NewWordHelperRPCClient(conn)
	r, err := c.FenciJieba(ctx, &wordhelper.FenciJiebaReq{
		Str: "我想开飞机",
		Type: 1,
	})
	if err != nil {
		logger.Error(err)
	}
	for _, v := range r.Data {
		logger.Info(v)
	}


}