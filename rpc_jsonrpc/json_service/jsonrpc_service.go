package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Arith struct {
}

// 请求结构体
type ArithRequest struct {
	A int
	B int
}

// 响应结构体
type ArithResponse struct {
	Pro int //乘积
	Quo int //商
	Rem int //余数
}

// 乘积方法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main() {
	//注册rpc服务
	rpc.Register(new(Arith))
	//采用http协议作为rpc载体
	rpc.HandleHTTP()

	lis, err := net.Listen("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Fatalln("fatal error:", err)
	}

	fmt.Fprintf(os.Stdout, "%s", "start connection\n")

	//接收客户端请求 并发处理 jsonrpc
	for {
		conn, err := lis.Accept() //接收客户端连接请求
		if err != nil {
			continue
		}

		//并发处理客户端请求
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}

	//常规启动http服务
	//http.Serve(lis,nil)
}
