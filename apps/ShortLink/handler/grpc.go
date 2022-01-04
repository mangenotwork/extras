package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/mangenotwork/extras/apps/ShortLink/model"
	"github.com/mangenotwork/extras/apps/ShortLink/proto"
	"github.com/mangenotwork/extras/apps/ShortLink/service"
)

type GRPCService struct {}


func (*GRPCService) Add (ctx context.Context, req *proto.AddReq) (*proto.AddResp, error){
	resp := new(proto.AddResp)
	if req.IsPrivacy && len(req.Password) < 1 {
		return resp, fmt.Errorf("设置了隐私但是password为空")
	}

	exp := req.Aging
	if exp == 0 {
		exp = req.Deadline
	}
	shortLink := &model.ShortLink{
		Short: "/"+service.MustGenerate(),
		Url: req.Url,
		Expiration: exp,
		IsPrivacy : req.IsPrivacy,
		Password : req.Password,
		Creation : time.Now().Unix(),
		View : 0,
		OpenBlockList : req.OpenBlockList,
		OpenWhiteList : req.OpenWhiteList,
		BlockList: req.BlockList,
		WhiteList: req.WhiteList,
	}

	err := shortLink.Save()
	if err != nil {
		return resp, err
	}

	resp.Url = shortLink.Short
	resp.Password = req.Password
	resp.Expire = time.Unix(exp, 0).Format("2006-01-02 15:04:05")
	return resp, nil
}

func (*GRPCService) Get (ctx context.Context, req *proto.GetReq) (*proto.GetResp, error){
	resp := new(proto.GetResp)
	link := new(model.ShortLink)
	err := link.Get(req.ShortLink)
	if err != nil {
		return resp, err
	}

	resp.Short = link.Short
	resp.Url = link.Url
	resp.Expiration = link.Expiration
	resp.IsPrivacy = link.IsPrivacy
	resp.Password = link.Password
	resp.Creation = link.Creation
	resp.View = link.View
	resp.OpenBlockList = link.OpenBlockList
	resp.OpenWhiteList = link.OpenWhiteList
	resp.BlockList = link.BlockList
	resp.WhiteList = link.WhiteList
	return resp, nil
}

func (*GRPCService) Modify (ctx context.Context, req *proto.ModifyReq) (*proto.ModifyResp, error){
	resp := new(proto.ModifyResp)

	link := new(model.ShortLink)
	err := link.Get(req.ShortLink)
	if err != nil {
		return resp, err
	}

	if link.IsPrivacy && req.Password != link.Password {
		return resp, fmt.Errorf("链接访问密码错误")
	}

	if len(req.Url) > 0 {
		link.Url = req.Url
	}
	link.IsPrivacy = req.IsPrivacy
	link.Password = req.Password

	err = link.Save()
	if err != nil {
		return resp, err
	}
	resp.Data = "成功"

	return resp, nil
}

func (*GRPCService) Del (ctx context.Context, req *proto.DelReq) (*proto.DelResp, error){
	resp := new(proto.DelResp)

	link := new(model.ShortLink)
	err := link.Get(req.ShortLink)
	if err != nil {
		return resp, err
	}

	if link.IsPrivacy && req.Password != link.Password {
		return resp, fmt.Errorf("链接访问密码错误")
	}

	err = link.Del()
	if err != nil {
		return resp, err
	}
	resp.Data = "成功"

	return resp, nil
}

