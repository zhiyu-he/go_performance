package main

import (
	"github.com/ThoseFlowers/thrift/lib/go/thrift"
	"github.com/zhiyu-he/go_performance/simple-thrift/echo"
	"fmt"
)

func main() {
	socket, err := thrift.NewTSocket("127.0.0.1:23333")
	if err != nil {
		panic(err)
	}

	err = socket.Open()
	if err != nil {
		panic(err)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTFastFrameBinaryProtocolFactoryDefault(1024)

	cli := echo.NewEchoServiceClientFactory(transportFactory.GetTransport(socket), protocolFactory)

	err = cli.Hi()
	if err != nil {
		panic(err)
	}
	rsp, err := cli.Do(&echo.EchoReq{SeqID:1, StrDat: "abc", BinDat: []byte("abc")})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", rsp)
}
