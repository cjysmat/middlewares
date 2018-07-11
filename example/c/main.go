package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/cjysmat/middlewares/example/inf"

	"math/rand"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	wg sync.WaitGroup
)

const (
	networkType = "tcp"
	server      = "127.0.0.1"
	port        = "41005"
	parallel    = 1 //连接并行度
	times       = 1 //每连接请求次数
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	currTime := time.Now()

	//并行请求
	for i := 0; i < int(parallel); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			exe()
		}()
	}
	wg.Wait()

	log.Printf("time taken: %.8f ", time.Now().Sub(currTime).Seconds()/parallel/times)
}

func exe() {
	conn, _ := grpc.Dial(server+":"+port, grpc.WithInsecure())
	defer conn.Close()
	client := inf.NewDataClient(conn)

	for i := 0; i < int(times); i++ {
		getUser(client)
	}
}

func getUser(client inf.DataClient) {

	var request inf.UserRq
	r := rand.Intn(parallel)
	request.Id = int32(r)

	client.GetUser(context.Background(), &request)
}
