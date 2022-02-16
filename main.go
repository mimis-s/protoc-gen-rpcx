package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gitee.com/mimis/protoc-gen-rpcx/generator"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// 定义模块
const tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
	{{- range $_, $m := .MethodList}}
	{{$m.MethodName}}(*{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
	{{- end}}
}

func Register{{.ServiceName}}(
	srv *rpc.Server, x {{.ServiceName}}Interface,
) error {
	if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
		return err
	}
	return nil
}

type {{.ServiceName}}Client struct {
	*rpc.Client
}

var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)

func Dial{{.ServiceName}}(network, address string) (
	*{{.ServiceName}}Client, error,
) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &{{.ServiceName}}Client{Client: c}, nil
}

{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(
	in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}},
) error {
	return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}
`

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

	log.Printf("genServiceCode[%v]\n", spec)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec) // 把spec传入模板，返回初始化好的模板buf
	if err != nil {
		log.Fatal(err)
	}
	p.P(buf.String()) // 把模板的内容写入生成的proto文件里面
}

// 定义netrpcPlugin类，generator 作为成员变量存在, 继承公有方法
type netrpcPlugin struct{ *generator.Generator }

// 返回插件名称
func (p *netrpcPlugin) Name() string {
	return "goPro"
}

// 通过g 进入初始化
func (p *netrpcPlugin) Init(g *generator.Generator) {
	log.Printf("通过g 进入初始化\n")
	p.Generator = g
}

// 生成导入包
func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

// 生成主体代码
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	log.Printf("生成主体代码\n")
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

// 自定义方法，生成导入包
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("// TODO: import code here")
	p.P(`import "net/rpc"`)
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

// 以下内容都来自protoc-gen-go/main.go
func main() {
	g := generator.New()
	log.Printf("proto goPro start\n")
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		g.Error(err, "reading input")
	}

	if err := proto.Unmarshal(data, g.Request); err != nil {
		g.Error(err, "parsing input proto")
	}

	if len(g.Request.FileToGenerate) == 0 {
		g.Fail("no files to generate")
	}

	log.Printf("g.CommandLineParameters[%v]\n", g.Request.GetParameter())
	g.CommandLineParameters(g.Request.GetParameter())

	// Create a wrapped version of the Descriptors and EnumDescriptors that
	// point to the file that defines them.
	g.WrapTypes()

	g.SetPackageNames()
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
