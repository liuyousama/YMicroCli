package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	code "ymicro-cli/template"
)

type MainGenerator struct {

}

func init()  {
	Register("main_generator", new(MainGenerator))
}

func (*MainGenerator)Generate(opt *Option, service *ServiceInfo) (err error) {
	var file *os.File
	defer func() {_ = file.Close()}()

	file, err = os.OpenFile(filepath.Join(opt.OutputFilePath,"main","main.go"),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		err = fmt.Errorf("Create main file failed: %v. ", err)
		return
	}

	//渲染到文件
	t := template.New("controller")
	t, err = t.Parse(code.MainTemplate)
	if err != nil {
		err = fmt.Errorf("Render main file failed: %v. ", err)
		return
	}
	err = t.Execute(file, service)
	if err != nil {
		err = fmt.Errorf("Write code to main file failed: %v. ", err)
		return
	}

	return
}