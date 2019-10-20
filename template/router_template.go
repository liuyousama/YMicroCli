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
	var response *pb.{{.ReturnsType}}
	c := new(controller.{{.Name}}Controller)
	if response, err = c.Validate(ctx, param); err != nil {
		return response, err
	}

	return c.Serve(ctx, param)
}
{{end}}

`
