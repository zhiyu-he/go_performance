package simple_impl

type Q struct {
	id int64
}

func (p *Q) GetKey() int64 {
	return p.id
}

func (p *Q) GetValue() interface{} {
	return p
}

type OP int

const (
	EVENT_GET OP = iota
	EVENT_SET
	EVENT_CLR
)

type Req struct {
	op     OP
	keys   []int64
	query  F
	values []IOP
	signal chan struct{}
}

type F func(m map[int64]interface{})

type IOP interface {
	GetKey() int64
	GetValue() interface{}
}

type Rsp struct {
	Status int
}

type KV struct {
	cache  map[int64]interface{}
	opChan chan *Req
}

func NewKV(cap int, concurrent int) *KV {
	kv := &KV{
		cache:  make(map[int64]interface{}, cap),
		opChan: make(chan *Req, concurrent),
	}
	go kv.Handle()
	return kv
}

func (p *KV) Exec(req *Req) {
	p.opChan <- req
}

// maybe 自己定义一个handle函数？
func (p *KV) Handle() {
	for {
		req := <-p.opChan
		switch req.op {
		case EVENT_GET:
			req.query(p.cache)
			req.signal <- struct{}{}
		case EVENT_SET:
			for _, val := range req.values {
				p.cache[val.GetKey()] = val.GetValue()
			}
		case EVENT_CLR:
			// todo
		}
	}
}
