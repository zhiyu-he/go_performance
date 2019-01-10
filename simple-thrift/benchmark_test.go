package simple_thrift

import (
	"testing"
	"github.com/zhiyu-he/go_performance/simple-thrift/echo"
	apacheThrift "github.com/apache/thrift/lib/go/thrift"
)

func BenchmarkApacheThriftRead(b *testing.B) {
	req := &echo.EchoReq{
		SeqId: 1,
		StrDat: "LogID",
		BinDat: []byte{0x01,0x02, 0x03},
	}
	req.AdIds = make([]int64, 0, 3000)
	for i := 0; i < 3000; i++ {
		req.AdIds = append(req.AdIds, int64(i))
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
