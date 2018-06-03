package main

// go build -gcflags -m xx.go
import "fmt"

type MiniModel struct {
	ValType int32

	ValType2     *int32
	IsXXTypeBool bool
}

func (p *MiniModel) IsXXType() bool {
	return p.ValType == 1 || p.ValType == 8 || p.ValType == 9
}

func (p *MiniModel) IsXXType3() bool {
	if p.ValType2 != nil && *p.ValType2 == 1 {
		return true
	}
	return false
}

func (p *MiniModel) IsXXType2() bool {
	return p.IsXXTypeBool
}

func (p *MiniModel) Set() {
	p.IsXXTypeBool = p.IsXXType()
}

func main() {

	mm := &MiniModel{}

	flag := mm.IsXXType()
	mm.IsXXType2()

	fmt.Printf("%v\n", flag)
}
