package test_service

import (
	"context"
	"fmt"
	"testing"

	"github.com/mimis-s/golang_tools/rpcx/service"
	"github.com/mimis-s/protoc-gen-rpcx/proto"
)

var (
	listenAddr string   = "localhost:8579"
	addr       string   = "localhost:8579"
	etcdAddrs  []string = []string{"192.168.1.98:2379"}
	isLocal    bool     = false
)

// 测试类
type Server struct {
	S *service.ServerManage
}

func (s *Server) Signin(ctx context.Context, req *proto.SigninReq, res *proto.SigninRes) error {
	fmt.Printf("签到成功ID:%v\n", req.ID)
	res.ID = req.ID + 1
	return nil
}

func Init() *Server {
	s := new(Server)
	sManager, err := proto.NewPackServiceAndRun(listenAddr, addr, etcdAddrs, s, "", false)
	if err != nil {
		panic(err)
	}
	go func() {
		err = sManager.Run()
		if err != nil {
			panic(err)
		}
	}()
	s.S = sManager
	return s
}

func TestService(t *testing.T) {
	s := Init()
	fmt.Print(s)
	select {}
}
