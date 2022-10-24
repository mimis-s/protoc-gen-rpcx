// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/mimis-s/golang_tools/rpcx/client"
	service "github.com/mimis-s/golang_tools/rpcx/service"
	"sync"
	"time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

var serverName string = "pack"

var callSingleMethodFunc func()

var PackClientInstance PackClientInterface
var PackClientOnce = new(sync.Once)

func newPackClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) PackClientInterface {
	if !isLocal {
		c := client.New(serverName, etcdAddrs, timeout, etcdBasePath)
		return &PackRpcxClient{
			c: c,
		}
	} else {
		return &PackLocalClient{}
	}
}

func SingleNewPackClient(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) {
	callSingleMethodFunc = func() {
		c := newPackClient(etcdAddrs, timeout, etcdBasePath, isLocal)
		PackClientInstance = c
	}
}

// 外部调用函数

func Signin(ctx context.Context,
	in *SigninReq) (*SigninRes, error) {

	if callSingleMethodFunc != nil {
		PackClientOnce.Do(callSingleMethodFunc)
	}

	out := new(SigninRes)
	out, err := PackClientInstance.Signin(ctx, in)
	return out, err
}

type PackClientInterface interface {
	Signin(context.Context, *SigninReq) (*SigninRes, error)
}

// rpcx客户端
type PackRpcxClient struct {
	c *client.ClientManager
}

func (c *PackRpcxClient) Signin(ctx context.Context,
	in *SigninReq) (*SigninRes, error) {
	out := new(SigninRes)
	err := c.c.Call(ctx, "Signin", in, out)
	return out, err
}

// 本地调用客户端
type PackLocalClient struct {
}

func (c *PackLocalClient) Signin(ctx context.Context,
	in *SigninReq) (*SigninRes, error) {
	out := new(SigninRes)
	err := PackServiceLocal.Signin(ctx, in, out)
	return out, err
}

type PackServiceInterface interface {
	Signin(context.Context, *SigninReq, *SigninRes) error
}

var PackServiceLocal PackServiceInterface

func RegisterPackService(s *service.ServerManage, hdlr PackServiceInterface) error {
	return s.RegisterOneService(serverName, hdlr)
}

func NewPackServiceAndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler PackServiceInterface, etcdBasePath string, isLocal bool) (*service.ServerManage, error) {
	if !isLocal {
		s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath)
		if err != nil {
			return nil, err
		}

		err = RegisterPackService(s, handler)
		if err != nil {
			return nil, err
		}

		go func() {
			err = s.Run(listenAddr)
			if err != nil {
				panic(fmt.Errorf("listen(%v) error(%v)", listenAddr, err))
			}
		}()
		return s, nil
	}

	// 本地调用的时候使用
	PackServiceLocal = handler
	return nil, nil
}
