package tpl

var TmpClientHandler string = `{{$root := .}}

var callSingleMethodFunc func()

var {{$root.ServiceName}}ClientInstance {{$root.ServiceName}}ClientInterface
var {{$root.ServiceName}}ClientOnce = new(sync.Once)


func new{{$root.ServiceName}}Client(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) {{.ServiceName}}ClientInterface {
	if !isLocal {
		c := client.New(serverName, etcdAddrs, timeout, etcdBasePath)
		return &{{$root.ServiceName}}RpcxClient{
			c: c,
		}
	}else{
		return &{{$root.ServiceName}}LocalClient{}
	}
}

func SingleNew{{$root.ServiceName}}Client(etcdAddrs []string, timeout time.Duration, etcdBasePath string, isLocal bool) {
	callSingleMethodFunc = func() {
		c := new{{$root.ServiceName}}Client(etcdAddrs, timeout, etcdBasePath, isLocal)
		{{$root.ServiceName}}ClientInstance = c
	}
}

// 外部调用函数
{{range $_, $m := .MethodList}}
func {{$m.MethodName}}(ctx context.Context, 
	in *{{$m.InputTypeName}}) (*{{$m.OutputTypeName}}, error) {

	if callSingleMethodFunc != nil {
		{{$root.ServiceName}}ClientOnce.Do(callSingleMethodFunc)
	}

	out := new({{$m.OutputTypeName}})
	out, err := {{$root.ServiceName}}ClientInstance.{{$m.MethodName}}(ctx, in)
    return out, err
}
{{end}}
`
