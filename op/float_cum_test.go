package op

import (
	"testing"
	"math/rand"
	"github.com/h2so5/half"
)

var f16 []half.Float16
var f32 []float32
var f64 []float64

const num = 1e6
func init() {
	for i := 0; i < num; i++ {
		f16 = append(f16, half.NewFloat16(rand.Float32()))
		f32 = append(f32, rand.Float32())
		f64 = append(f64, rand.Float64())
	}
}

func doFloat16() half.Float16 {
	var f16sum half.Float16 = 1.0
	for _, i := range f16 {
		f16sum *= i
	}
	return f16sum
}

func doFloat32() float32 {
	var f32sum float32 = 1.0
	for _, i := range f32 {
		f32sum *= i
	}
	return f32sum
}

func doFloat64() float64 {
	var f64sum float64 = 1.0
	for _, i := range f64 {
		f64sum *= i
	}
	return f64sum
}

func BenchmarkDoFloat16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doFloat16()
	}
}

func BenchmarkDoFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doFloat32()
	}
}

func BenchmarkDoFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doFloat64()
	}
}
