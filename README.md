# protobuf插件之rpcx

#### 介绍
基于protoc-gen-go扩展proto使用rpcx进行通信
简单修改generator文件中的goFileName函数的name，自定义工具生成的文件名字
修改CommandLineParameters函数里面的pluginList := ""为空

每个服务都要在使用的时候定义一个client, 这个操作有点麻烦, 所以这里把客户端写成单例模式,
然后加入对于本地调用和rpc调用的判断, 应用场景扩展到本地和远程都能进行调用

#### 关键函数使用说明
New{{.ServiceName}}ServiceAndRun: 创建一个本地或者rpcx服务器监听等待客户端连接
Register{{.ServiceName}}Service:  将已经创建好的rpcx服务器对象绑定etcd, 启动监听(本地调用时不会创建rpcx对象, 调用函数返回nil)

#### 使用

<p>1：编译代码：go build -o protoc-gen-rpcx main.go</p>
<p>2：找到protobuf的环境变量路径$PATH,把protoc-gen-rpcx可执行文件放到路径下，
  这样在protoc编译proto文件的时候自定义的--xxx_out会去$PBPATH路径下寻找执行文件</p>
<p>3：编译.proto文件生成我们自己的代码 protoc -I./ --rpcx_out=. test.proto</p>

<p>注意：仔细检查PATH的文件，千万不要乱放执行文件，不然怎么运行的都不晓得</p>

<p>4：测试生成proto文件：protoc -I./ --rpcx_out=. --go_out=. ./proto/test.proto</p>

