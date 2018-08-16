package simple_impl

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Queue struct {
	tail *node
	head *node
	len  int64
	lock *sync.Mutex
}

type node struct {
	data interface{}
	next *node
}

func NewQueue() *Queue {
	q := &Queue{
		lock: new(sync.Mutex),
		len:  0,
	}
	return q
}

func (p *Queue) Push(item interface{}) {
	p.lock.Lock()
	node := &node{
		data: item,
	}
	if p.tail == nil {
		p.head = node
		p.tail = node
	} else {
		p.tail.next = node
		p.tail = node
	}
	p.len++
	p.lock.Unlock()
}

func (p *Queue) Pop() interface{} {
	p.lock.Lock()
	if p.head == nil {
		p.lock.Unlock()
		return nil
	}
	i := p.head
	p.head = i.next

	if p.head == nil {
		p.tail = nil
	}
	p.len--
	p.lock.Unlock()
	return i.data
}

func (p *Queue) Len() int64 {
	return atomic.LoadInt64(&p.len)
}

type ObjectPool struct {
	queue Queue
	size  int
	name  string
	f     func() interface{}
}

func NewObjectPool(name string, size int, f func() interface{}) *ObjectPool {
	pool := &ObjectPool{
		name:  name,
		size:  size,
		f:     f,
		queue: Queue{len: 0, lock: new(sync.Mutex)},
	}
	for i := 0; i < size; i++ {
		pool.queue.Push(f())
	}
	pool.Stats()
	return pool
}

func (p *ObjectPool) Stats() {
	go func() {
		for {
			fmt.Printf("[ObjectPool] name: %s pool-size: %d queue-size: %d\n", p.name, p.size, p.queue.Len())
			time.Sleep(10 * time.Second)
		}
	}()
}

func (p *ObjectPool) Borrow() interface{} {
	return p.queue.Pop()
}

func (p *ObjectPool) Return(i interface{}) {
	p.queue.Push(i)
}
