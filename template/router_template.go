package template

var RouterTemplate = `
package router

import (
	"context"
	"{{.Module}}/controller"
	"{{.Module}}/pb"
)

type Router struct {
	
}

{{range .Rpcs}}
func (*Router){{.Name}}(ctx context.Context,param *pb.{{.RequestType}}) (*pb.{{.ReturnsType}}, error){
	var err error
	var response *pb.{{.Rpc.ReturnsType}}
	response, c := new(controller.{{.Name}}Controller)
	if err = c.Validate(ctx, param); err != nil {
		return nil, nil
	}

	return c.Serve(ctx, param)
}
{{end}}

`
