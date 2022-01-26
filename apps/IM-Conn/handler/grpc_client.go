package handler

import (
	"fmt"
	"github.com/mangenotwork/extras/apps/IM-Conn/proto"
	"github.com/mangenotwork/extras/common/grpc"
	"github.com/mangenotwork/extras/common/logger"
)

func AuthToken(token string) (string, error) {
	client, err := grpc.NewClient(grpc.ClientArg{
		ServiceAddr: "192.168.0.9:29128",
		ServiceName: "IM-User",
	})
	if err != nil {
		logger.Error(err)
		return "", err
	}
	conn, ctx, err := client.Conn()
	logger.Info("conn = ", conn)
	defer conn.Close()
	if err != nil {
		return "", err
	}

	c := proto.NewIMUserRPCClient(conn)
	r, err := c.Authentication(ctx, &proto.AuthReq{
		Token: token,
	})
	logger.Info("AuthToken --> ", r, err)
	if err != nil {
		return "", err
	}
	if r.State != 1 {
		return "", fmt.Errorf("身份验证失败")
	}
	return r.Uid, nil
}

func Login(account, password string) (string, string, error) {
	client, err := grpc.NewClient(grpc.ClientArg{
		ServiceAddr: "192.168.0.9:29128",
		ServiceName: "IM-User",
	})
	if err != nil {
		logger.Error(err)
		return "", "", err
	}
	conn, ctx, err := client.Conn()
	logger.Info("conn = ", conn)
	defer conn.Close()
	if err != nil {
		return "", "", err
	}

	c := proto.NewIMUserRPCClient(conn)
	r, err := c.Login(ctx, &proto.LoginReq{
		Account: account,
		Password: password,
	})
	logger.Info("Login --> ", r, err)
	if err != nil {
		return "", "", err
	}
	if r.State != 1 {
		return "", "", fmt.Errorf("身份验证失败")
	}
	return r.Uid, r.Token, nil
}