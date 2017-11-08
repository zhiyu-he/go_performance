package build_in

import "testing"

type Ad struct {
	Operators []Operator
}

func (p *Ad) Filter() bool {
	for _, op := range p.Operators {
		if !op.Match() {
			return false
		}
	}
	return true
}

func (p *Ad) Filter2() bool {
	if len(p.Operators) == 0 {
		return true
	}
	for _, op := range p.Operators {
		if !op.Match() {
			return false
		}
	}
	return true
}


type Operator struct {

}

func (p *Operator) Match() bool {
	return true
}


const const_size =  25000

var adsPtr []*Ad

func init () {
	adsPtr = make([]*Ad, const_size)
	for i := 0; i < const_size; i++ {
		adsPtr[i] = &Ad{
			Operators: nil,
		}
	}
}

func DoNormal() {
	for _, ad := range adsPtr {
		if ad.Filter() {

		}
	}
}

func DoOPT() {
	for _, ad := range adsPtr {
		if ad.Filter2() {

		}
	}
}

func DoOPT2() {
	for _, ad := range adsPtr {
		if len(ad.Operators) == 0 || ad.Filter2() {

		}
	}
}

func BenchmarkDoNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoNormal()
	}
}

func BenchmarkDoOPT1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoOPT()
	}
}

func BenchmarkDoOPT2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoOPT2()
	}
}
