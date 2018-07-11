package middlewares

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Logging interceptor for grpc
func Logging(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	start := time.Now()
	resp, err = handler(ctx, req)
	fmt.Printf("calling=%v, timne=%d, req=%s, resp=%v, err=%v\n", info.FullMethod, time.Since(start), marshal(req), marshal(resp), err)

	return resp, err
}
