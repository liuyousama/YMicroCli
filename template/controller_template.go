package template

var ControllerTemplate = `
package controller

import (
	"context"
	"{{.Module}}/pb"
)

type {{.Rpc.Name}}Controller struct {
	
}

func (*{{.Rpc.Name}}Controller)Validate(
	ctx context.Context, param *pb.{{.Rpc.RequestType}}) (*pb.{{.Rpc.ReturnsType}}, error) {
	res := new(pb.{{.Rpc.ReturnsType}})	

	return res, nil
}

func (*{{.Rpc.Name}}Controller)Serve(
	ctx context.Context, param *pb.{{.Rpc.RequestType}}) (*pb.{{.Rpc.ReturnsType}}, error) {
	res := new(pb.{{.Rpc.ReturnsType}})

	return res, nil
}
`