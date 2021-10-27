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
	resp.Sub = req.Sub
	resp.Str, resp.Time = service.BlockWorkTrie.BlockWord(req.Str, req.Sub)
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) Add(ctx context.Context, req *proto.AddReq) (*proto.AddResp, error) {
	start := time.Now()
	resp := &proto.AddResp{
		Code: 0,
		Msg: "succeed",
	}
	service.AddWord(req.Word)
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) Del(ctx context.Context, req *proto.DelReq) (*proto.DelResp, error) {
	start := time.Now()
	resp := &proto.DelResp{
		Code: 0,
		Msg: "succeed",
	}
	service.DelWord(req.Word)
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

func (*GRPCService) WhiteWordAdd (ctx context.Context, req *proto.WhiteWordAddReq) (*proto.WhiteWordAddResp, error) {
	start := time.Now()
	resp := &proto.WhiteWordAddResp{
		Code: 0,
		Msg: "succeed",
	}
	service.WhiteAddWord(req.Word)
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) WhiteWordDel(ctx context.Context, req *proto.WhiteWordDelReq) (*proto.WhiteWordDelResp, error) {
	start := time.Now()
	resp := &proto.WhiteWordDelResp{
		Code: 0,
		Msg: "succeed",
	}
	service.WhiteDelWord(req.Word)
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) WhiteWordGet(ctx context.Context, req *proto.WhiteWordGetReq) (*proto.WhiteWordGetResp, error) {
	start := time.Now()
	resp := new(proto.WhiteWordGetResp)
	resp.List = service.WhiteGetWord()
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) IsHave(ctx context.Context, req *proto.IsHaveReq) (*proto.IsHaveResp, error) {
	start := time.Now()
	resp := new(proto.IsHaveResp)
	resp.IsHave = 0
	if !service.BlockWorkTrie.IsHave(req.Str) {
		resp.IsHave = 1
	}
	utils.RpcLog(start, ctx)
	return resp, nil
}

func (*GRPCService) IsHaveList(ctx context.Context, req *proto.IsHaveListReq) (*proto.IsHaveListResp, error) {
	start := time.Now()
	resp := new(proto.IsHaveListResp)
	resp.List = service.BlockWorkTrie.BlockHaveList(req.Str)
	resp.IsHave = 0
	if len(resp.List) > 0 {
		resp.IsHave = 1
	}
	utils.RpcLog(start, ctx)
	return resp, nil
}