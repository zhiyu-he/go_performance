package simple_impl

import (
	"fmt"
	"testing"
)

var p *Pool

func init() {
	p = NewPool(10)
}

func TestSubmitTask(t *testing.T) {
	cb := p.Submit(func() (interface{}, error) {
		return "fuck", nil
	})
	res, err := cb.Fin()
	fmt.Printf("res: %v err: %v\n", res, err)

}
