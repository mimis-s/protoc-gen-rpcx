package test_client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/mimis-s/protoc-gen-rpcx/proto"
)

var (
	etcdAddrs []string = []string{"192.168.1.98:2379"}
)

func Init() {
	proto.SingleNewPackClient(etcdAddrs, time.Minute, "", false)

	signinReq := &proto.SigninReq{
		ID: 1,
	}
	signinRes, err := proto.Signin(context.Background(), signinReq)
	if err != nil {
		panic(err)
	}
	fmt.Printf("签到成功ID:%v\n", signinRes.ID)
}

func TestClient(t *testing.T) {
	Init()
}
