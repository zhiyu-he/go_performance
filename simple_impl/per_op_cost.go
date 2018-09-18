package simple_impl

type AA struct {
	v int
	p *int
	v64 int64
	v32 int32
}

func m() interface{} {
	a := struct{}{}
	return &a
}


func pass(a AA, newV int) {
	a.v = newV
}

func doCount(l []*AA) int {
	var sum = 0
	for _, a := range l {
		if a.p == nil {
			sum += 1
		}
	}
	return sum
}


func doCount2(l []AA) int {
	var sum = 0
	for _, a := range l {
		if a.p == nil {
			sum += 1
		}
	}
	return sum
}


func mSlice1(size int) {
	_ = make([]*AA, size)
}

func mSlice2(size int) {
	_ = make([]AA, size)
}
