package simple_thrift

import (
	"testing"
	"github.com/zhiyu-he/go_performance/simple-thrift/echo"
	apacheThrift "github.com/apache/thrift/lib/go/thrift"
)

func BenchmarkApacheThriftRead(b *testing.B) {
	req := &echo.EchoReq{
		SeqID: 1,
		StrDat: "LogID",
		BinDat: []byte{0x01,0x02, 0x03},
	}
	buffer := apacheThrift.NewTMemoryBufferLen(1024)
	factory := apacheThrift.NewTBinaryProtocolFactoryDefault()
	p := factory.GetProtocol(buffer)
	req.Write(p)
	buf := *buffer.Buffer
	for i := 0; i < b.N; i++ {
		req.Read(p)
		*buffer.Buffer = buf
	}
}
