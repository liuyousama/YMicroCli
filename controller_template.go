package main

var controllerTemplate = `
package controller

import (
	"context"
	"{{.module}}/pb"
)

type {{.rpc.Name}}Controller struct {
	
}

func (*{{.rpc.Name}}Controller)Validate(
	ctx context.Context, param *pb.{{.rpc.RequestType}}) error {

	return nil
}

func (*{{.rpc.Name}}Controller)Serve(
	ctx context.Context, param *pb.{{.rpc.RequestType}}) (*pb.{{rpc.ReturnsType}}, error) {

	return nil, nil
}
`