package main

var controllerTemplate = `
package controller

import (
	"context"
	""
)

type {{.Service.Name}} struct {}

{{range .Rpcs}}
func (s *{{$.Service.Name}}) {{.Name}}(
	ctx context.Context, request *pb.{{.RequestType}}) (*pb.ReturnsType, error) {

    return
}
{{end}}
`
