package template

var RouterTemplate = `
package router

import (
	"context"
	"{{.module}}/controller"
	"{{.module}}/pb"
)

type Router struct {
	
}

{{range .Rpcs}}
func (*Router){{.Name}}(ctx context.Context,param *pb.{{.RequestType}}) (*pb.{{.ReturnsType}}, error){
	var err error
	c := new(controller.{{.Name}}Controller)
	if err = c.Validate(ctx, param); err != nil {
		return nil, nil
	}

	return c.Serve(ctx, param)
}
{{end}}

`
