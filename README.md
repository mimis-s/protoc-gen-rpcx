# protobuf插件之rpcx

#### 介绍
基于protoc-gen-go扩展proto使用rpcx进行通信

#### 使用

1：编译代码：go build -o protoc-gen-rpcx main.go
2：找到protobuf的环境变量路径$PBPATH,把protoc-gen-rpcx可执行文件放到路径下，
  这样在protoc编译proto文件的时候自定义的--xxx_out会去$PBPATH路径下寻找执行文件
3：编译.proto文件生成我们自己的代码 protoc -I./ --rpcx_out=. test.proto

