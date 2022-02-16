package rpcx

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/rpcxio/rpcx-etcd/serverplugin"
	"github.com/smallnest/rpcx/server"
)

type Args struct {
	A string
}

type Reply struct {
	B string
}

var (
	listenAddr = flag.String("listenAddr", "localhost:8972", "server listen address")
	addr       = flag.String("addr", "localhost:8972", "server address")
	etcdAddr   = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath   = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	go http.ListenAndServe(":9981", nil)

	//创建一个服务器实例
	//服务器实例中有成员变量(Plugins PluginContainer),这个便利包含了服务器上的所有插件
	//还有一个成员函数(AuthFunc),可以检查客户端是否被授权
	s := server.NewServer()

	//etcd
	addRegistryPlugin(s)

	//注册方法,这里面会去注册s的每一个Plugins插件,调用他们的注册函数
	s.RegisterName("Arith", new(Arith), "")
	err := s.Serve("tcp", *listenAddr)
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {

	//prcx使用的etcd插件EtcdV3RegisterPlugin
	//关于插件的参数解析
	//1:  ServiceAddress:本机监听地址，这个对外暴露的监听地址,格式为tcp@address:port
	//2:  EtcdServers:etcd集群服务器地址
	//3:  BasePath:服务前缀。如果有多个项目同时使用zookeeper，避免命名冲突，可以设置这个参数，为当前服务器设置命名空间
	//4:  Metrics:用来更新服务的TPS(系统吞吐量 = 并发数/平均响应时间)
	//5:  UpdateInterval:服务刷新间隔，如果在一定间隔内(下面设置为一分钟)没有刷新，服务就会从etcd中删除
	r := &serverplugin.EtcdV3RegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		EtcdServers:    []string{*etcdAddr},
		BasePath:       *basePath,
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}

type Arith int

func (t *Arith) PrintIP(ctx context.Context, args *Args, reply *Reply) error {
	fmt.Printf("tcp:%v \n", *addr)
	return nil
}
