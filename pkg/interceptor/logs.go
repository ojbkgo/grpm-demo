package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

var LogsInterceptor1 grpc.UnaryServerInterceptor
var LogsInterceptor2 grpc.UnaryServerInterceptor

func init() {

	LogsInterceptor1 = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("LogsInterceptor1:", info.FullMethod, req)

		return nil, nil
	}

	LogsInterceptor2 = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("LogsInterceptor2:", req)
		return nil, nil
	}
}
