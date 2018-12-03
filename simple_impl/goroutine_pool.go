package simple_impl

var (
	worker []Worker
	source []*Source
)

func init() {
	source = make([]*Source, 10)
	worker = make([]Worker, 10)
}

type Pool struct {
	curIdx int
}

type Source struct {
	task Task
	sig  chan struct{}
	cb   *Callback
}

func NewPool(size int) *Pool {
	p := &Pool{
		curIdx: 0,
	}
	for i := 0; i < size; i++ {
		w := func(sig chan struct{}, idx int) {
			for {
				<-sig
				s := source[idx]
				s.cb.result, s.cb.err = s.task()
				source[idx].cb.fin <- struct{}{}
			}
		}
		sig := make(chan struct{}, 1)
		source[i] = &Source{
			sig: sig,
			cb: &Callback{
				fin: make(chan struct{}, 1),
			},
		}
		worker[i] = w
		go worker[i](sig, i)

	}
	return p
}

func (p *Pool) Submit(f Task) *Callback {
	s := source[p.curIdx]
	s.task = f
	s.sig <- struct{}{}
	return s.cb
}

type Callback struct {
	result interface{}
	err    error
	fin    chan struct{}
}

func (p *Callback) Fin() (interface{}, error) {
	<-p.fin
	return p.result, p.err
}

type Task func() (interface{}, error)
type Worker func(sig chan struct{}, idx int)
