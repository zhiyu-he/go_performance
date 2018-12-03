package build_in

import (
	"testing"
	"sync"
)

func Test1(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Wait()
	t.Log("Here!")
}
