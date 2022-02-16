package rpcx

// import (
// 	"context"
// 	"flag"
// 	"log"
// 	"time"

// 	"zhangbin.com/study/study_rpcx/args"

// 	etcd_client "github.com/rpcxio/rpcx-etcd/client"
// 	"github.com/smallnest/rpcx/client"
// )

// var (
// 	etcdAddr = flag.String("etcdAddr", "192.168.1.19:2379", "etcd address")
// 	basePath = flag.String("base", "/rpcx_test", "prefix path")
// )

// func main() {
// 	flag.Parse()

// 	//NewEtcdV3Discovery指定basePath和etcd集群地址,方法路径(服务名称)
// 	d, _ := etcd_client.NewEtcdV3Discovery(*basePath, "Arith", []string{*etcdAddr}, false, nil)
// 	//获取服务器的ip端口,和一些其它信息
// 	xclient := client.NewXClient("Arith", client.Failover, client.RoundRobin, d, client.DefaultOption)
// 	defer xclient.Close()

// 	for {
// 		arg := &args.Args{A: "client"}
// 		reply := &args.Reply{}
// 		//连接服务器,远程调用
// 		err := xclient.Call(context.Background(), "PrintIP", arg, reply)
// 		if err != nil {
// 			log.Printf("failed to call: %v\n", err)
// 			time.Sleep(5 * time.Second)
// 			continue
// 		}

// 		log.Printf("%s , %s", arg.A, reply.B)

// 		time.Sleep(5 * time.Second)
// 	}
// }
