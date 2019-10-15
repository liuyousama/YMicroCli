package template

var MiddlewareCode = `
package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func GrpcMiddlewareOption () grpc.ServerOption {
	return grpc.UnaryInterceptor(interceptor)
}

func interceptor(ctx context.Context, req interface{},
info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	//在此处理你想要在执行rpc之前处理的逻辑

	return handler(ctx, req)
}
`
