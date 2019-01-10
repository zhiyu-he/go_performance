package simple_thrift

import (
	"testing"
	"github.com/zhiyu-he/go_performance/simple-thrift/echo"
	apacheThrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/stretchr/testify/assert"
)

var req *echo.EchoReq
var dat []byte
var btf *apacheThrift.TBufferedTransportFactory = apacheThrift.NewTBufferedTransportFactory(1024)

func init() {
	req = &echo.EchoReq{
		SeqID: 1,
		StrDat: "LogID",
		BinDat: []byte{0x01,0x02, 0x03},
	}
	req.AdIds = make([]int64, 0, 3000)
	for i := 0; i < 3000; i++ {
		req.AdIds = append(req.AdIds, int64(i))
	}

	trans := apacheThrift.NewTMemoryBuffer()
	factory := apacheThrift.NewTBinaryProtocolFactoryDefault()
	p := factory.GetProtocol(trans)
	req.Write(p)
	p.Flush()
	dat = trans.Bytes()
}

func TestCorrect(t *testing.T) {
	var r *echo.EchoReq = &echo.EchoReq{}
	trans := apacheThrift.NewTMemoryBuffer()
	trans.Write(dat)
	factory := apacheThrift.NewTBinaryProtocolFactoryDefault()
	p := factory.GetProtocol(btf.GetTransport(trans))
	e := r.Read(p)
	assert.Nil(t, e)
	assert.EqualValues(t, req, r)
}


func BenchmarkApacheThriftWrite(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		trans := apacheThrift.NewTMemoryBuffer()
		factory := apacheThrift.NewTBinaryProtocolFactoryDefault()
		p := factory.GetProtocol(btf.GetTransport(trans))
		req.Write(p)
		p.Flush()
	}
}

func BenchmarkApacheThriftRead(b *testing.B) {
	var r *echo.EchoReq = &echo.EchoReq{}
	for i := 0; i <= b.N; i++ {
		trans := apacheThrift.NewTMemoryBuffer()
		trans.Write(dat)
		factory := apacheThrift.NewTBinaryProtocolFactoryDefault()
		p := factory.GetProtocol(btf.GetTransport(trans))
		r.Read(p)
	}
}
