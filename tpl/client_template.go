package tpl

var TmpClient string = `{{$root := .}}

type {{.ServiceName}}ClientInterface interface {
	{{- range $_, $m := .MethodList}}
	{{$m.MethodName}}(context.Context, *{{$m.InputTypeName}}) (*{{$m.OutputTypeName}}, error)
	{{- end}}
}

// rpcx客户端
type {{.ServiceName}}RpcxClient struct {
	c *client.ClientManager
}

{{range $_, $m := .MethodList}}
func (c *{{$root.ServiceName}}RpcxClient) {{$m.MethodName}}(ctx context.Context, 
	in *{{$m.InputTypeName}}) (*{{$m.OutputTypeName}}, error) {
    out := new({{$m.OutputTypeName}})
	err := c.c.Call(ctx, "{{$m.MethodName}}", in, out)
    return out, err
}
{{end}}

// 本地调用客户端
type {{.ServiceName}}LocalClient struct {
}

{{range $_, $m := .MethodList}}
func (c *{{$root.ServiceName}}LocalClient) {{$m.MethodName}}(ctx context.Context, 
	in *{{$m.InputTypeName}}) (*{{$m.OutputTypeName}}, error) {
    out := new({{$m.OutputTypeName}})
	err := {{$root.ServiceName}}ServiceLocal.{{$m.MethodName}}(ctx, in, out)
    return out, err
}
{{end}}

`
