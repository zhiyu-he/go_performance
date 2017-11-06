package simple_impl

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

var idList []int

func init() {
	const size = 30000
	idList = make([]int, size)
	for i := 0; i < size; i++ {
		idList[i] = rand.Int()
	}
}

func shuffle(ids []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(ids) - 1; i >= 0; i-- {
		j := r.Intn(i + 1)
		ids[j], ids[i] = ids[i], ids[j]
	}
}

func TestShuffle(t *testing.T) {
	const size = 30000
	idList1 := make([]int, size)
	idList2 := make([]int, size)
	for i := 0; i < size; i++ {
		num := rand.Int()
		idList1[i] = num
		idList2[i] = num
	}

	fmt.Printf("l1: %#v\n", idList1)
	fmt.Printf("l2: %#v\n", idList2)
	shuffle(idList1)
	sort.Ints(idList1)
	sort.Ints(idList2)

	fmt.Printf("l1: %#v\n", idList1)
	fmt.Printf("l2: %#v\n", idList2)

	for i := 0; i < size; i++ {
		if idList1[i] != idList2[i] {
			panic("abc")
		}
	}
}

func BenchmarkShuffle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		shuffle(idList)
	}
}
