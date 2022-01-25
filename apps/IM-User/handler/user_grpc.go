package handler

import (
	"context"
	"github.com/mangenotwork/extras/apps/IM-User/proto"
	"github.com/mangenotwork/extras/apps/IM-User/service"
)

type GRPCService struct {

}

// Authentication 用户token验证
// resp.State=1 验证成功
// resp.State=2 验证失败
func (*GRPCService) Authentication(ctx context.Context, req *proto.AuthReq) (*proto.AuthResp, error) {
	var err error
	resp := &proto.AuthResp{
		State: 0,
	}
	uParams := &service.UserParam{
		TokenStr: req.Token,
	}
	resp.State, err = uParams.AuthToken()
	return resp, err
}