package test_service

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/mimis/golang-tool/rpcx/service"
	"gitee.com/mimis/protoc-gen-rpcx/proto"
)

var (
	listenAddr string   = "localhost:8579"
	addr       string   = "localhost:8579"
	etcdAddrs  []string = []string{"192.168.1.165:2379"}
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
	sManager, err := proto.NewPackServiceAndRun(listenAddr, addr, etcdAddrs, s, "")
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
