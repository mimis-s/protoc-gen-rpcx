package test_client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"gitee.com/mimis/protoc-gen-rpcx/proto"
)

var (
	etcdAddrs []string = []string{"192.168.1.165:2379"}
)

type ServiceClient struct {
	Pack proto.PackClientInterface
}

func Init() (*ServiceClient, error) {
	s := new(ServiceClient)
	pack := proto.NewPackClient(etcdAddrs, time.Minute, "")
	s.Pack = pack

	signinReq := &proto.SigninReq{
		ID: 1,
	}
	signinRes, err := s.Pack.Signin(context.Background(), signinReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("签到成功ID:%v\n", signinRes.ID)

	return s, nil
}

func TestClient(t *testing.T) {
	Init()
}
