package engine

import (
	"github.com/mangenotwork/extras/apps/IM-User/handler"
	"github.com/mangenotwork/extras/apps/IM-User/proto"
	"github.com/mangenotwork/extras/common/conf"
	"github.com/mangenotwork/extras/common/grpc"
	"github.com/mangenotwork/extras/common/logger"
	"github.com/mangenotwork/extras/common/utils"
)

func StartRPC(){
	go func() {
		logger.Info("StartRPC")
		g, err := grpc.NewServer(grpc.ServerArg{
			IP: "",
			Port: utils.Str2Int(conf.Arg.GrpcServer.Prod),
			Name: "BlockWord",
		})
		if err != nil {
			panic(err)
		}

		proto.RegisterIMUserRPCServer(g.Server, &handler.GRPCService{})
		g.Run()
		logger.Info("[RPC] Listening and serving TCP on %d", g.Port)

	}()
}