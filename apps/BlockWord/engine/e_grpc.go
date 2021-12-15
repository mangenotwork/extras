package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/handler"
	"github.com/mangenotwork/extras/apps/BlockWord/proto"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/grpc"
	"github.com/mangenotwork/extras/common/utils"
	"log"
)

func StartRpcServer(){
	go func() {
		g, err := grpc.NewServer(grpc.ServerArg{
			IP: "",
			Port: utils.Str2Int(conf.Arg.GrpcServer.Prod),
			Name: "BlockWord",
		})
		if err != nil {
			panic(err)
		}
		proto.RegisterMessageRPCServer(g.Server, &handler.GRPCService{})
		g.Run()
		log.Print("[RPC] Listening and serving TCP on %d", g.Port)
	}()
}
