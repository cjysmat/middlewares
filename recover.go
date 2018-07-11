package middlewares

import (
	"fmt"
	"runtime"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	MAXSTACKSIZE = 4096
)

// Recover interceptor to handle grpc panic
func Recover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// Recover func
	defer func() {
		if err := recover(); err != nil {
			// log stack
			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]
			fmt.Printf("panic grpc invoke: %s, err=%v, stack:\n%s", info.FullMethod, err, string(stack))

			// if panic, set custom error to 'err', in order that client and sense it.
			err = grpc.Errorf(codes.Internal, "panic error: %v", err)
		}
	}()

	return handler(ctx, req)
}

func StreamRecover(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// Recover func
	defer func() {
		if err := recover(); err != nil {
			// log stack
			stack := make([]byte, MAXSTACKSIZE)
			stack = stack[:runtime.Stack(stack, false)]
			fmt.Printf("panic grpc invoke: %s, err=%v, stack:\n%s", info.FullMethod, err, string(stack))

			// if panic, set custom error to 'err', in order that client and sense it.
			err = grpc.Errorf(codes.Internal, "panic error: %v", err)
		}
	}()

	return handler(srv, ss)
}
