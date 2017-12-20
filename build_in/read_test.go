package build_in
  
import (
    "encoding/hex"
    "math/rand"
    "testing"
)

// ref: xiaoguoqiao, 特别鸣谢
// 鉴于此, thrift在golang中对list<Obj> a; []*Obj 这种代码简直匪夷所思
// go test -v -bench=. read_test.go -benchmem

type A struct {
    Num1 int
    Num2 int
    Num3 int
    Num4 int
    Num5 int
    str1 string
    str2 string
}

var m1 = []A{}
var m2 = []*A{}

func init() {
    b := make([]byte, 10000)
    rand.Read(b)
    str := hex.EncodeToString(b)
    for i := 0; i < 10; i++ {
        x := A{
            Num1: rand.Int(),
            Num2: rand.Int(),
            Num3: rand.Int(),
            Num4: rand.Int(),
            Num5: rand.Int(),
            str1: str,
            str2: str,
        }
        m1 = append(m1, x)
        y := x
        m2 = append(m2, &y)
    }
}

func BenchmarkRead1(b *testing.B) {
    var sum = 0
    for i := 0; i < b.N; i++ {
        for _, m := range m1 {
            sum += m.Num3
        }
    }
}
func BenchmarkRead2(b *testing.B) {
    var sum = 0
    for i := 0; i < b.N; i++ {
        for _, m := range m2 {
            sum += m.Num3
        }
    }
}
func BenchmarkRead3(b *testing.B) {
	var sum = 0
	for i := 0; i < b.N; i++ {
		for idx := range m1 {
			sum += m1[idx].Num3
		}
	}
}
