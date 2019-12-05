package main

import (
	"context"
	"github.com/jian/grpc-go-demo/src/test"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	// 建立连接导gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := test.NewWaiterClient(conn)

	// 模拟请求数据
	res := "test123"
	// os.Args[1] 为用户执行输入的参数 如：go run ***.go 123
	if len(os.Args) > 1 {
		res = os.Args[1]
	}

	// 调用gRPC接口
	tr, err := t.DoMD5(context.Background(), &test.Req{JsonStr: res})
	if err != nil {
		log.Fatalf("could not great: %v", err)
	}
	log.Fatalf("服务端响应： %s", tr.BackJson)
}