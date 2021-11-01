package utils

import (
	"context"
	"github.com/mangenotwork/extras/common/conf"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"log"
	"runtime"
	"strings"
	"time"
)

func runFuncName()string{
	pc := make([]uintptr,1)
	runtime.Callers(3,pc)
	f := runtime.FuncForPC(pc[0])
	fName := f.Name()
	fList := strings.Split(fName,".")
	return fList[len(fList)-1]
}

func RpcLog(start time.Time,ctx context.Context) {
	if conf.Arg.GrpcServer.Log {
		pr, _ := peer.FromContext(ctx)
		md, _ := metadata.FromIncomingContext(ctx)
		log.Printf("[GRPC] %v %s %s %v", pr.Addr.String(), md, runFuncName(), time.Since(start))
	}

}