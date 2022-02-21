package test_client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"gitee.com/mimis/golang-tool/rpcx/service"
	"gitee.com/mimis/protoc-gen-rpcx/proto"
)

var (
	etcdAddrs []string = []string{":8789"}
)

type ServiceClient struct {
	Pack service.PackServiceInterface
}

func Init() (*ServiceClient, error) {
	s := new(ServiceClient)
	pack := service.NewPackService(etcdAddrs, time.Minute, false, true)
	s.Pack = pack

	signinReq := &proto.SigninReq{
		ID: 1,
	}
	signinRes, err := s.Pack.Signin(context.Background(), signinReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("客户端返回%v\n", signinRes)
	return s, nil
}

func TestClient(t *testing.T) {

}
