package engine

import (
	"github.com/mangenotwork/extras/apps/ShortLink/handler"
	"github.com/mangenotwork/extras/apps/ShortLink/proto"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/grpc"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func StartRpcServer(){
	go func() {
		logger.Info("StartRpcServer...")
		g, err := grpc.NewServer(grpc.ServerArg{
			IP: "",
			Port: utils.Str2Int(conf.Arg.GrpcServer.Prod),
			Name: "WordHelper",
		})
		if err != nil {
			panic(err)
		}
		proto.RegisterShortLinkRPCServer(g.Server, &handler.GRPCService{})
		g.Run()
		logger.Info("[RPC] Listening and serving TCP on %d", g.Port)
	}()
}