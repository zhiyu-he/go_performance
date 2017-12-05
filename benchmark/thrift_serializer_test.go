package benchmark

// thrift -out . -r --gen go:thrift_import=code.byted.org/gopkg/thrift echo.thrift
// go version go1.9.2 darwin/amd64

import (
	"code.byted.org/gopkg/thrift"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/zhiyu-he/go_performance/benchmark/echo"
	"fmt"
)

var (
	normal *thrift.TSerializer
	opt    *thrift.TSerializer

	normalD *thrift.TDeserializer
	optD    *thrift.TDeserializer

	req = &echo.EchoReq{SeqId: 20171208, StrDat: "echo", MDat: map[string]float64{"ctr":0.123, "cvr": 0.567}}

	byteReq []byte
)

func init() {
	t := thrift.NewTMemoryBufferLen(1024)
	transport := thrift.NewTFramedTransport(t)
	p := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(transport)

	normal = &thrift.TSerializer{
		Transport: t,
		Protocol:  p,
	}

	t2 := thrift.NewTMemoryBufferLen(1024)
	p2 := thrift.NewTFastFrameBinaryProtocolFactoryDefault(1024).GetProtocol(t2)
	opt = &thrift.TSerializer{
		Transport: t2,
		Protocol:  p2,
	}

	normalD = &thrift.TDeserializer{
		Transport: t,
		Protocol:  p,
	}

	optD = &thrift.TDeserializer{
		Transport: t2,
		Protocol:  p2,
	}

	byteReq, _ = normal.Write(req)

}

func TestEqual(t *testing.T) {
	// test serializer equal
	dat, _ := opt.Write(req)
	assert.Equal(t, byteReq, dat)

	// test de-serializer equal
	reqNormal := &echo.EchoReq{}
	reqOPT := &echo.EchoReq{}

	optD.Read(reqOPT, byteReq)
	normalD.Read(reqNormal, byteReq)
	assert.EqualValues(t, reqNormal, reqOPT)

}

func apacheThriftWrite() {
	normal.Write(req)
}

func optThriftWrite() {
	opt.Write(req)
}

func apacheThriftRead(dat []byte) {
	normalD.Read(req, dat)
}

func optThriftRead(dat []byte) {
	optD.Read(req, dat)
}

func BenchmarkApacheThrift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		apacheThriftWrite()
	}
}

func BenchmarkOPTThrift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optThriftWrite()
	}
}

func BenchmarkApacheThriftRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		apacheThriftRead(byteReq)
	}
}

func BenchmarkOPTThriftRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optThriftRead(byteReq)
	}
}
