{{$root := .}}

type {{.ServiceName}}ServiceInterface interface {
	{{- range $_, $m := .MethodList}}
	{{$m.MethodName}}(context.Context, *{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
	{{- end}}
}

func Register{{.ServiceName}}Service(s *service.ServerManage, hdlr {{.ServiceName}}ServiceInterface) error {
    return s.RegisterOneService(serverName, hdlr)
}

func New{{.ServiceName}}ServiceAndRun(listenAddr, exposeAddr string, etcdAddrs []string, handler {{.ServiceName}}ServiceInterface, etcdBasePath string) (*service.ServerManage, error) {
    s, err := service.New(exposeAddr, etcdAddrs, etcdBasePath)
	if err != nil {
		return nil, err
	}

	err = Register{{.ServiceName}}Service(s, handler)
	if err != nil {
		return nil, err
	}

	go func() {
		err = s.Run(listenAddr)
		if err != nil {
			panic(fmt.Errorf("listen(%v) error(%v)", listenAddr, err))
		}
	}()

	return s, nil
}
