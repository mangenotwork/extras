package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/handler"
	"github.com/mangenotwork/extras/apps/BlockWord/proto"
	"github.com/mangenotwork/extras/common/conf"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartRpcServer(){
	go func() {
		listen, err := net.Listen("tcp", ":"+conf.Arg.GrpcServer.Prod)
		if err != nil {
			panic(err)
		}
		grpcServer := grpc.NewServer()
		proto.RegisterMessageRPCServer(grpcServer, &handler.GRPCService{})
		log.Println("Starting block word grpc server -> ", conf.Arg.GrpcServer.Prod)
		err = grpcServer.Serve(listen)
		if err != nil {
			panic(err)
		}

	}()
}


