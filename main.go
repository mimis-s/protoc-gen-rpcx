package main

import (
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/mimis-s/protoc-gen-rpcx/generator"
	_ "github.com/mimis-s/protoc-gen-rpcx/plugin"
)

// 以下内容都来自protoc-gen-go/main.go
func main() {
	// protoc读取proto文件，以二进制传入插件，被反序列化为CodeGeneratorRequest
	// 做完插件逻辑之后构造CodeGeneratorResponse对象，并转为二进制输出，再由protoc转化为.go文件输出
	g := generator.New()
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	// 这里面主要存储proto文件，包，文件内部一些参数的信息，数据实在太多，不一一列举
	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	// 判断是否有填入待编译proto文件
	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	// --rpcx_out=plugins=rpcx:.，那么Parameter就是后面的plugins=rpcx
	// 这里以plugins举例，--rpcx_out后面正确的写法就应该是要加plugins
	// 旨在如果有很多插件，那么把插件列表清空，只留下我们选择的插件,但是我为了方便，CommandLineParameters函数判断插件的逻辑被改为了默认使用所有插件
	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	// 将g.Request里面的proto文件信息,和命令行输入的参数信息转化为g.allFiles,g.allFilesByName里面,
	// 用FileDescriptor, EnumDescriptors字段存储,便于后面计算
	g.WrapTypes()

	// 设置所有生成的文件包名
	g.SetPackageNames()

	// 把proto文件里面的方法名(.包名.方法名如.proto.SigninReq)和方法的数据信息用map关联起来
	g.BuildTypeNameMap()

	g.GenerateAllFiles() // 使用模板生成自定义pb文件

	// Send back the results.
	data, err = proto.Marshal(g.Response)
	if err != nil {
		g.Error(err, "failed to marshal output proto")
	}

	// log.Printf("goPro end%s\n", data)

	_, err = os.Stdout.Write(data)
	if err != nil {
		g.Error(err, "failed to write output proto")
	}
}
