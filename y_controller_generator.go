package main

import (
	"fmt"
	"github.com/emicklei/proto"
	"os"
	"path/filepath"
	"text/template"
	code "ymicro-cli/template"
)

type ControllerGenerator struct {
}

func init() {
	controllerGenerator := &ControllerGenerator{}

	Register("controller_generator", controllerGenerator)
}

func (g *ControllerGenerator) Generate(opt *Option, service *ServiceInfo) (err error) {
	var file *os.File
	defer func() { _ = file.Close() }()

	for _, rpc := range service.Rpcs {
		//打开（创建）控制器文件
		file, err = os.OpenFile(
			filepath.Join(opt.OutputFilePath, "controller", rpc.Name+".go"),
			os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)

		if err != nil {
			err = fmt.Errorf("Create contoller file %s.go failed: %v. ", rpc.Name, err)
			return
		}
		//渲染到文件
		t := template.New("controller")
		t, err = t.Parse(code.ControllerTemplate)
		if err != nil {
			err = fmt.Errorf("Render controller %s failed: %v. ", rpc.Name, err)
			return
		}

		var templateVar = struct {
			Rpc    *proto.RPC
			Module string
		}{rpc, opt.ProjectModule}
		err = t.Execute(file, templateVar)

		if err != nil {
			err = fmt.Errorf("Write code to controller %s failed: %v. ", rpc.Name, err)
			return
		}

		_ = file.Close()
	}
	return
}
