package main

type I interface {
	SayHi()
	GoodBye()
}

type Base struct{}

func (p *Base) SayHi() {
	println("Base SayHi")
}

func (p *Base) GoodBye() {
	println("Base GoodBye")
}

type Child struct {
	*Base
}

func (p *Child) SayHi() {
	println("Child SayHi")
}

func main() {
	b := &Base{}
	b.SayHi()

	c := &Child{
		Base: b,
	}
	c.SayHi()
	c.GoodBye()
}
