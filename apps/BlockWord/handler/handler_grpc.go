package handler

import (
	"context"
	"time"

	"github.com/mangenotwork/extras/apps/BlockWord/proto"
	"github.com/mangenotwork/extras/apps/BlockWord/service"
	"github.com/mangenotwork/extras/common/utils"

)

type GRPCService struct {}

func (*GRPCService) Do(ctx context.Context, req *proto.DoReq) (*proto.DoResp, error) {
	start := time.Now()
	resp := new(proto.DoResp)
	resp.Str = service.BlockWorkTrie.Replace(req.Str, req.Sub)
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) Add(ctx context.Context, req *proto.AddReq) (*proto.AddResp, error) {
	start := time.Now()
	resp := &proto.AddResp{
		Code: 0,
		Msg: "succeed",
	}
	err := service.AddWord(req.Word)
	if err != nil {
		resp.Code = 201
		resp.Msg = err.Error()
	}
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) Del(ctx context.Context, req *proto.DelReq) (*proto.DelResp, error) {
	start := time.Now()
	resp := &proto.DelResp{
		Code: 0,
		Msg: "succeed",
	}
	err := service.DelWord(req.Word)
	if err != nil {
		resp.Code = 201
		resp.Msg = err.Error()
	}
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) Get(ctx context.Context, req *proto.GetReq) (*proto.GetResp, error) {
	start := time.Now()
	resp := new(proto.GetResp)
	resp.List = service.GetWord()
	utils.RpcLog(start, ctx)
	return resp, nil
}
