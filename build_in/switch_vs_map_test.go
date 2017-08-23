package build_in

import "testing"

func IsExceptedValueSwitch(val int) bool {
	switch val {
	case
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10:
		return true
	default:
		return false
	}
}

func IsExceptedValueSwitch2(val string) bool {
	switch val {
	case
		"abc",
		"efg":
		return true
	default:
		return false
	}
}


var m = map[int]bool{
	1: true,
}

var m2 = map[string]bool{
	"abc": true,
	"efg": true,
}


func IsExceptedValueMap(val int) bool {
	return m[val]
}

func IsExceptedValueMap2(val string) bool {
	return m2[val]
}


func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueSwitch(1)
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueMap(1)
	}
}

func BenchmarkSwitch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueSwitch2("abc")
	}
}

func BenchmarkMap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueMap2("abc")
	}
}
