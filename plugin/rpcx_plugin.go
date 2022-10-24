package plugin

import (
	"bytes"
	"fmt"
	"html/template"
	"path"
	"strconv"
	"strings"

	"github.com/mimis-s/protoc-gen-rpcx/generator"
	"github.com/mimis-s/protoc-gen-rpcx/tpl"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

const (
	contextPkgPath = "context"
	clientPkgPath  = "github.com/mimis-s/golang_tools/rpcx/client"
	serverPkgPath  = "github.com/mimis-s/golang_tools/rpcx/service"
)

var (
	contextPkg string
	clientPkg  string
	serverPkg  string
)

// 定义服务和接口描述结构
type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}

// 解析每个服务的ServiceSpec元信息
func (p *netrpcPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	spec := &ServiceSpec{ServiceName: generator.CamelCase(svc.GetName())}
	for _, m := range svc.Method {
		// TypeName旨在判断生成文件是否在包路径，如果在那么就返回方法名，不在就在方法名上面加一个包路径
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}

	return spec
}

// 自定义方法，生成导入代码
func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)

	serverName := `var serverName string = "` + strings.ToLower(spec.ServiceName) + `"`
	p.P(serverName)

	// 客户端实例
	{
		var buf bytes.Buffer
		// t, err := template.ParseFiles("template/rpcx_client.tpl")
		// if err != nil {
		// 	fmt.Printf("template.ParseFiles is err:%v\n", err)
		// 	return
		// }
		t := template.Must(template.New("").Parse(tpl.TmpClientHandler))

		err := t.Execute(&buf, spec) // 把spec传入模板，返回初始化好的模板buf
		if err != nil {
			fmt.Printf("Execute is err:%v\n", err)
			return
		}
		p.P(buf.String()) // 把模板的内容写入生成的proto文件里面
	}

	// 客户端
	{
		var buf bytes.Buffer
		// t, err := template.ParseFiles("template/rpcx_client.tpl")
		// if err != nil {
		// 	fmt.Printf("template.ParseFiles is err:%v\n", err)
		// 	return
		// }
		t := template.Must(template.New("").Parse(tpl.TmpClient))

		err := t.Execute(&buf, spec) // 把spec传入模板，返回初始化好的模板buf
		if err != nil {
			fmt.Printf("Execute is err:%v\n", err)
			return
		}
		p.P(buf.String()) // 把模板的内容写入生成的proto文件里面
	}

	// 服务器
	{
		var buf bytes.Buffer
		// t, err := template.ParseFiles("template/rpcx_service.tpl")
		// if err != nil {
		// 	fmt.Printf("template.ParseFiles is err:%v\n", err)
		// 	return
		// }
		t := template.Must(template.New("").Parse(tpl.TmpService))

		err := t.Execute(&buf, spec) // 把spec传入模板，返回初始化好的模板buf
		if err != nil {
			fmt.Printf("Execute is err:%v\n", err)
			return
		}
		p.P(buf.String()) // 把模板的内容写入生成的proto文件里面
	}

}

// 定义netrpcPlugin类，generator 作为成员变量存在, 继承公有方法
type netrpcPlugin struct{ *generator.Generator }

// 返回插件名称
func (p *netrpcPlugin) Name() string {
	return "rpcx"
}

// 通过g 进入初始化
func (p *netrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g

	contextPkg = generator.RegisterUniquePackageName("context", nil)
	clientPkg = generator.RegisterUniquePackageName("client", nil)
	serverPkg = generator.RegisterUniquePackageName("service", nil)
}

// 生成导入包
func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

// 生成主体代码
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

// 自定义方法，生成导入包
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("import (")
	//g.P(apiPkg, " ", strconv.Quote(path.Join(g.gen.ImportPrefix, apiPkgPath)))
	p.P("\"time\"")
	// p.P("\"reflect\"")
	// p.P("\"encoding/json\"")
	p.P("\"sync\"")
	p.P(contextPkg, " ", strconv.Quote(path.Join(p.Generator.ImportPrefix, contextPkgPath)))
	p.P(clientPkg, " ", strconv.Quote(path.Join(p.Generator.ImportPrefix, clientPkgPath)))
	p.P(serverPkg, " ", strconv.Quote(path.Join(p.Generator.ImportPrefix, serverPkgPath)))
	p.P(")")
	p.P()
}

// 自定义方法，生成导入代码
/*
func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
}
*/

// 注册插件
func init() {
	generator.RegisterPlugin(new(netrpcPlugin))
}
