package simple_impl

import (
	"sync"
	"unsafe"
)

type SliceObject []interface{}

type SliceHeader struct {
	ptr unsafe.Pointer
	len int
	cap int
}

type SimpleSlicePool struct {
	poolSize int
	sliceCap int

	startIndex int
	endIndex   int
	slices     []*SliceObject

	mu *sync.Mutex
}

func NewSimpleSlicePool(poolSize, sliceCap int) *SimpleSlicePool {
	slices := make([]*SliceObject, poolSize)
	for i := 0; i < poolSize; i++ {
		s := make(SliceObject, 0, sliceCap)
		slices[i] = &s
	}
	return &SimpleSlicePool{
		poolSize:   poolSize,
		sliceCap:   sliceCap,
		startIndex: 0,
		endIndex:   poolSize - 1,
		slices:     slices,
		mu:         new(sync.Mutex),
	}
}

func (p *SimpleSlicePool) _getIndex() int {
	var idx int = -1
	p.mu.Lock()

	if p.startIndex != p.endIndex {
		idx = p.startIndex
		p.startIndex = (p.startIndex + 1) % p.poolSize
		p.mu.Unlock()
	} else {
		p.mu.Unlock()
	}
	return idx
}

func (p *SimpleSlicePool) _putIndex() {
	p.mu.Lock()
	p.endIndex = (p.endIndex + 1) % p.poolSize
	p.mu.Unlock()
}

func (p *SimpleSlicePool) GetSliceObject() (*[]interface{}, int) {
	idx := p._getIndex()
	if idx == -1 {
		return nil, idx
	} else {
		return (*[]interface{})(unsafe.Pointer(p.slices[idx])), idx
	}
}

func (p *SimpleSlicePool) PutSliceObject(slicePtr *[]interface{}, idx int) {
	//p.slices[idx] = (*SliceObject)(unsafe.Pointer(&s))
	tmpPtr := (*SliceHeader)(unsafe.Pointer(slicePtr))
	(*tmpPtr).len = 0
	p._putIndex()
}
