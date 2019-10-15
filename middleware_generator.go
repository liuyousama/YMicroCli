package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type MiddlewareTemplate struct {

}

func init()  {
	Register("middle_template", new(MiddlewareTemplate))
}

func (*MiddlewareTemplate) Generate(opt *Option, service *ServiceInfo) (err error) {
	var file *os.File
	defer func() {_ = file.Close()}()
	file, err = os.OpenFile(
		filepath.Join(opt.OutputFilePath,"middleware","middleware.go"),
		os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
	if err != nil {
		err = fmt.Errorf("Create middlerware file failed: %v. ", err)
		return
	}

	_, err = fmt.Fprint(file, middlewareCode)
	if err != nil {
		err = fmt.Errorf("Write code to middleware file failed: %v. ", err)
		return
	}

	return
}
