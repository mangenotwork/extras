package engine

import (
	"github.com/mangenotwork/extras/apps/BlockWord/handler"
	"github.com/mangenotwork/extras/apps/BlockWord/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartRpcSrc(){
	go func() {
		listen, err := net.Listen("tcp", ":1211")
		if err != nil {
			panic(err)
		}
		grpcServer := grpc.NewServer()
		proto.RegisterMessageRPCServer(grpcServer, &handler.GRPCService{})
		log.Println("Starting block word grpc server")
		err = grpcServer.Serve(listen)
		if err != nil {
			panic(err)
		}

	}()
}


