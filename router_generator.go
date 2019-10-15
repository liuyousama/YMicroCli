package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type RouterGenerator struct {

}

func init() {
	controllerGenerator := &RouterGenerator{}

	Register("router_generator", controllerGenerator)
}

func (*RouterGenerator) Generate(opt *Option, service *ServiceInfo) (err error) {
	var file *os.File
	file, err = os.OpenFile(
		filepath.Join(opt.OutputFilePath,"router","router.go"),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	defer func() {_ = file.Close()}()

	t := template.New("router")
	t, err = t.Parse(routerTemplate)
	if err != nil {
		err = fmt.Errorf("Render router failed: %v. ", err)
		return
	}
	err = t.Execute(file, service)
	if err != nil {
		err = fmt.Errorf("Write code to router failed: %v. ", err)
		return
	}

	return
}
