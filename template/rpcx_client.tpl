{{$root := .}}
var serverName string = "pack"

type {{.ServiceName}}ClientInterface interface {
	{{- range $_, $m := .MethodList}}
	{{$m.MethodName}}(context.Context, *{{$m.InputTypeName}}) (*{{$m.OutputTypeName}}, error)
	{{- end}}
}

func New{{.ServiceName}}Client(etcdAddrs []string, timeout time.Duration, etcdBasePath string) {{.ServiceName}}ClientInterface {
	c := client.New(serverName, etcdAddrs, timeout, etcdBasePath)

	return &{{.ServiceName}}Client{
		c: c,
	}
}

type {{.ServiceName}}Client struct {
	c *client.ClientManager
}

{{range $_, $m := .MethodList}}
func (c *{{$root.ServiceName}}Client) {{$m.MethodName}}(ctx context.Context, 
	in *{{$m.InputTypeName}}) (*{{$m.OutputTypeName}}, error) {
    out := new({{$m.OutputTypeName}})
	err := c.c.Call(ctx, "{{$m.MethodName}}", in, out)
    return out, err
}
{{end}}
