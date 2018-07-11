package main

import (
	"log"
	"net"
	"runtime"
	"strconv"

	"github.com/cjysmat/middlewares"
	"github.com/cjysmat/middlewares/example/inf"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = "41005"
)

type Data struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(middlewares.UnaryInterceptorChain(middlewares.Recover, middlewares.Logging)),
		grpc.StreamInterceptor(middlewares.StreamServerChain(middlewares.StreamRecover)))
	inf.RegisterDataServer(s, &Data{})
	s.Serve(lis)

	log.Printf("grpc server in: %s", port)
}

func (t *Data) GetUser(ctx context.Context, request *inf.UserRq) (response *inf.UserRp, err error) {

	response = &inf.UserRp{
		Name: strconv.Itoa(int(request.Id)) + ":12313test",
	}

	//panic(errors.New("panic test"))

	return response, err
}
