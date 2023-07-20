package tpl

var TmpService string = `{{$root := .}}
type {{.ServiceName}}ServiceInterface interface {
	{{- range $_, $m := .MethodList}}
	{{$m.MethodName}}(context.Context, *{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
	{{- end}}
}

var {{.ServiceName}}ServiceLocal {{.ServiceName}}ServiceInterface

func Register{{.ServiceName}}Service(s *service.ServerManage, hdlr {{.ServiceName}}ServiceInterface) error {
	// 本地调用的时候使用(rpc本地客户端对应调用本地服务器)
	{{.ServiceName}}ServiceLocal = hdlr
	return s.RegisterOneService(serverName, hdlr)
}

func New{{.ServiceName}}ServiceAndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler {{.ServiceName}}ServiceInterface, etcdBasePath string, isLocal bool) (*service.ServerManage, error) {
    if !isLocal {
		s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath, listenAddr)
		if err != nil {
			return nil, err
		}
	
		err = Register{{.ServiceName}}Service(s, handler)
		if err != nil {
			return nil, err
		}
	
		return s, nil
	}

	// 本地调用的时候使用
	{{.ServiceName}}ServiceLocal = handler
	return nil, nil
}`
