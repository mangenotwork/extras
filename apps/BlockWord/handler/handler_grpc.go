package handler

import (
	"context"
	"github.com/mangenotwork/extras/apps/BlockWord/proto"
	"github.com/mangenotwork/extras/apps/BlockWord/service"
)

type GRPCService struct {}

func (*GRPCService) Do(ctx context.Context, req *proto.DoReq) (*proto.DoResp, error) {
	resp := new(proto.DoResp)
	resp.Sub = req.Sub
	resp.Str, resp.Time = service.BlockWorkTrie.BlockWord(req.Str, req.Sub)
	return resp, nil
}

func (*GRPCService) Add(ctx context.Context, req *proto.AddReq) (*proto.AddResp, error) {
	resp := &proto.AddResp{
		Code: 0,
		Msg: "succeed",
	}
	service.AddWord(req.Word)
	return resp, nil
}

func (*GRPCService) Del(ctx context.Context, req *proto.DelReq) (*proto.DelResp, error) {
	resp := &proto.DelResp{
		Code: 0,
		Msg: "succeed",
	}
	service.DelWord(req.Word)
	return resp, nil
}

func (*GRPCService) Get(ctx context.Context, req *proto.GetReq) (*proto.GetResp, error) {
	resp := new(proto.GetResp)
	resp.List = service.GetWord()
	return resp, nil
}

func (*GRPCService) WhiteWordAdd (ctx context.Context, req *proto.WhiteWordAddReq) (*proto.WhiteWordAddResp, error) {
	resp := &proto.WhiteWordAddResp{
		Code: 0,
		Msg: "succeed",
	}
	service.WhiteAddWord(req.Word)
	return resp, nil
}

func (*GRPCService) WhiteWordDel(ctx context.Context, req *proto.WhiteWordDelReq) (*proto.WhiteWordDelResp, error) {
	resp := &proto.WhiteWordDelResp{
		Code: 0,
		Msg: "succeed",
	}
	service.WhiteDelWord(req.Word)
	return resp, nil
}

func (*GRPCService) WhiteWordGet(ctx context.Context, req *proto.WhiteWordGetReq) (*proto.WhiteWordGetResp, error) {
	resp := new(proto.WhiteWordGetResp)
	resp.List = service.WhiteGetWord()
	return resp, nil
}

func (*GRPCService) IsHave(ctx context.Context, req *proto.IsHaveReq) (*proto.IsHaveResp, error) {
	resp := new(proto.IsHaveResp)
	resp.IsHave = 0
	if !service.BlockWorkTrie.IsHave(req.Str) {
		resp.IsHave = 1
	}
	return resp, nil
}

func (*GRPCService) IsHaveList(ctx context.Context, req *proto.IsHaveListReq) (*proto.IsHaveListResp, error) {
	resp := new(proto.IsHaveListResp)
	resp.List = service.BlockWorkTrie.BlockHaveList(req.Str)
	resp.IsHave = 0
	if len(resp.List) > 0 {
		resp.IsHave = 1
	}
	return resp, nil
}