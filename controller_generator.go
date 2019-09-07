package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type ControllerGenerator struct {

}

func init() {
	controllerGenerator := &ControllerGenerator{}

	Register("controller_generator", controllerGenerator)
}

func (g *ControllerGenerator) Generate(opt *Option, services []*ServiceInfo) (err error) {
	for _, service := range services {
		//打开（创建）控制器文件
		file, err := os.OpenFile(
			filepath.Join(opt.OutputFilePath,"controller",service.Service.Name+".go"),
			os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
		defer file.Close()
		if err != nil {
			fmt.Printf("Create contoller file %s.go failed: %v", service.Service.Name, err)
			return
		}
		//渲染到文件
		t := template.New("main")
		t, err = t.Parse(mainTemplate)
		if err != nil {
			fmt.Printf("Render controller %s failed: %v", service.Service.Name, err)
			return
		}
		err = t.Execute(file, service)
		if err != nil {
			fmt.Printf("Write code to controller %s failed: %v", service.Service.Name, err)
			return
		}

		_ = file.Close()
	}
	return
}