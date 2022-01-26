package handler

import (
	"context"
	"github.com/mangenotwork/extras/apps/IM-User/proto"
	"github.com/mangenotwork/extras/apps/IM-User/service"
	"github.com/mangenotwork/extras/common/logger"
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
	logger.Debug("token = ", req.Token)
	uParams := &service.UserParam{
		TokenStr: req.Token,
	}
	resp.State, resp.Uid, err = uParams.AuthToken()
	return resp, err
}

func (*GRPCService) Login(ctx context.Context, req *proto.LoginReq) (*proto.LoginResp, error) {
	var err error
	resp := &proto.LoginResp{
		State: 1,
	}
	uParams := &service.UserParam{
		Account: req.Account,
		Password: req.Password,
	}
	token, uid, err := uParams.Token()
	if err != nil {
		resp.State = 2
	}
	resp.Token = token
	resp.Uid = uid
	return resp, err
}