package simple_impl

import (
	"fmt"
	"testing"
)

type Q struct {
	id int64
}

func (p *Q) GetKey() int64 {
	return p.id
}

func (p *Q) GetValue() interface{} {
	return p
}

var kv = NewKV(100, 2)

func TestKV(t *testing.T) {
	req := &Req{
		op:     EVENT_SET,
		values: []IOP{&Q{id: 1}, &Q{id: 2}},
	}
	kv.Exec(req)

	var hit = make(map[int64]*Q, 16)
	var missIds = make([]int64, 0, 16)
	query := func(cacheM map[int64]interface{}) {
		for _, k := range []int64{1, 2, 3} {
			if raw, ok := cacheM[k]; ok {
				hit[k] = raw.(*Q)
			} else {
				missIds = append(missIds, k)
			}
		}
	}
	req2 := &Req{
		op:     EVENT_GET,
		keys:   []int64{1},
		query:  query,
		signal: make(chan struct{}),
	}
	kv.Exec(req2)
	<-req2.signal
	fmt.Printf("%v %v\n", hit, missIds)
}
