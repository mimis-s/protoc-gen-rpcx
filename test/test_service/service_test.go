package test_service

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/mimis/golang-tool/rpcx/service"
	"gitee.com/mimis/protoc-gen-rpcx/proto"
)

var (
	listenAddr string   = ":8848"
	exposeAddr string   = ":8919"
	etcdAddrs  []string = []string{":8379"}
	isLocal    bool     = false
)

// 测试类
type Server struct {
	S *service.ServerManage
}

func (s *Server) Signin(context.Context, *proto.SigninReq, *proto.SigninRes) error {
	fmt.Printf("签到成功\n")
	return nil
}

func Init() *Server {
	s := new(Server)
	sManager, err := service.NewPackHandlerAndRun(listenAddr, exposeAddr, etcdAddrs, s, isLocal)
	if err != nil {
		panic(err)
	}
	s.S = sManager
	return s
}

func TestService(t *testing.T) {
	s := Init()
	fmt.Print(s)
	select {}
}
