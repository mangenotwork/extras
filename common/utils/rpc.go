package utils

import (
	"context"
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
	pr, _ := peer.FromContext(ctx)
	log.Printf("[%s] %s %s %v", pr.Addr.String(), "GRPC", runFuncName(), time.Since(start))
}