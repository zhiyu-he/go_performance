package main

import (
	"github.com/ThoseFlowers/thrift/lib/go/thrift"
	"github.com/zhiyu-he/go_performance/simple-thrift/echo"
	"syscall"
	"os/signal"
	"os"
)

type Handler struct {
}

func (p *Handler) Hi() error {
	return nil
}

func (p *Handler) Do(req *echo.EchoReq) (*echo.EchoRsp, error) {
	return &echo.EchoRsp{
		Status: 0,

	}, nil
}


func NewSimpleThriftServer() {
	h := &Handler{}
	processor := echo.NewEchoServiceProcessor(h)

	socket, err := thrift.NewTServerSocket("127.0.0.1:23333")
	if err != nil {
		panic(err)
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())

	server := thrift.NewTSimpleServer4(processor, socket, transportFactory, protocolFactory)
	err = server.Serve()
	if err != nil {
		panic(err)
	}
}


func main() {
	NewSimpleThriftServer()
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	<-termChan
}