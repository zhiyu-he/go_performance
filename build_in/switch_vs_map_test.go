package build_in

import "testing"

var s1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
var s2 = []string{"abc", "efg", "hij", "klm"}

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
		10,
		11,
		12,
		13,
		14,
		15,
		16:
		return true
	default:
		return false
	}
}

func IsExceptedValueSwitch2(val string) bool {
	switch val {
	case
		"abc",
		"efg",
		"hij",
		"klm":
		return true
	default:
		return false
	}
}

var m = map[int]bool{
	1:  true,
	2:  true,
	3:  true,
	4:  true,
	5:  true,
	6:  true,
	7:  true,
	8:  true,
	9:  true,
	10: true,
	11: true,
	12: true,
	13: true,
	14: true,
	15: true,
	16: true,
}

var m2 = map[string]bool{
	"abc": true,
	"efg": true,
	"hij": true,
	"klm": true,
}

func IsExceptedValueMap(val int) bool {
	return m[val]
}

func IsExceptedValueMap2(val string) bool {
	return m2[val]
}

func IsExceptedValueSlice(val int) bool {
	for _, item := range s1 {
		if item == val {
			return true
		}
	}
	return false
}

func IsExceptedValueSlice2(val string) bool {
	for _, item := range s2 {
		if item == val {
			return true
		}
	}
	return false

}

func BenchmarkSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueSwitch(16)
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueMap(16)
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueSlice(1)
	}
}

func BenchmarkSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsExceptedValueSlice2("abc")
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
