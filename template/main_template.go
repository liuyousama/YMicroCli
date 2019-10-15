package template


var MainTemplate = `
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"{{.Module}}/middleware"
	"{{.Module}}/pb"
	"{{.Module}}/router"
)

const port = 28080

func main()  {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d"), port)
	if err != nil {
		fmt.Println(err)
		return
	}
	gServer := grpc.NewServer(middleware.GrpcMiddlewareOption())
	pb.Register{{.Service.Name}}Server(gServer, new(router.Router))
	err = gServer.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
`