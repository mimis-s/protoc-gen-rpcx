# protobuf插件之rpcx

#### 介绍
基于protoc-gen-go扩展proto使用rpcx进行通信
简单修改generator文件中的goFileName函数的name，自定义工具生成的文件名字
修改CommandLineParameters函数里面的pluginList := ""为空

#### 使用

<p>1：编译代码：go build -o protoc-gen-rpcx main.go</p>
<p>2：找到protobuf的环境变量路径$PATH,把protoc-gen-rpcx可执行文件放到路径下，
  这样在protoc编译proto文件的时候自定义的--xxx_out会去$PBPATH路径下寻找执行文件</p>
<p>3：编译.proto文件生成我们自己的代码 protoc -I./ --rpcx_out=. test.proto</p>

<p>注意：仔细检查PATH的文件，千万不要乱放执行文件，不然怎么运行的都不晓得</p>

<p>4：测试生成proto文件：protoc -I./ --rpcx_out=. --go_out=. ./proto/test.proto</p>

