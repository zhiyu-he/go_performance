package simple_impl

import "container/heap"

type Dirty struct {
	ts int64
	item interface{}
	idx int
}

type MinHeap []*Dirty


func newMinHeap(cap int) MinHeap {
	return make(MinHeap, 0, cap)
}

func (p MinHeap) Len() int {
	return len(p)
}

func (p MinHeap) Less(i, j int) bool {
	return p[i].ts < p[j].ts
}

func (p MinHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].idx = i
	p[j].idx = j
}

func (p *MinHeap) Push(x interface{}) {
	n := len(*p)
	c := cap(*p)
	if n + 1 < c {
		npq := make(MinHeap, n, c*2)
		copy(npq, *p)
		*p = npq
	}
	*p = (*p)[0:n+1]
	dirty := x.(*Dirty)
	dirty.idx = n
	(*p)[n] = dirty
}

func (p *MinHeap) Pop() interface{} {
	n := len(*p)
	c := cap(*p)
	if n < (c/2) && c > 25 {
		npq := make(MinHeap, n, c/2)
		copy(npq, *p)
		*p = npq
	}
	dirty := (*p)[n-1]
	dirty.idx = -1
	*p = (*p)[0 : n-1]
	return dirty
}

func (p *MinHeap) PeekAndShift(max int64) (*Dirty, int64) {
	if p.Len() == 0 {
		return nil, 0
	}

	item := (*p)[0]
	if item.ts > max {
		return nil, item.ts - max
	}
	heap.Remove(p, 0)
	return item, 0
}

type DelayPool struct {
	pure chan interface{}
	dirtyHeap MinHeap
	cleanFunc func(i interface{})
}


