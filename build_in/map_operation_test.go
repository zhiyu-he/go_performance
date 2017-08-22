package build_in


var m = make(map[int64]int64, 120000)

func init() {
	for i := 0; i < 120000; i++ {
		m[int64(i)]= int64(i)
	}
}


func emptyMap() {
	for k := range  m {
		delete(m, k)
	}
}

