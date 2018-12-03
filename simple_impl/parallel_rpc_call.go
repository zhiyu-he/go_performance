package simple_impl

import "sync"

/*
	Author: dongtiancong.
*/
type Function func() (interface{}, error)

type CallBack struct {
	task   Function
	result interface{}
	err    error
	fin    sync.WaitGroup
}

func (p *CallBack) Fin() (interface{}, error) {
	p.fin.Wait()
	return p.result, p.err
}

func Future4(f Function) Function {
	var result interface{}
	var err error
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		result, err = f()
	}()
	return func() (interface{}, error) {
		wg.Wait()
		return result, err
	}
}

func Future(f Function) Function {
	var result interface{}
	var err error
	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		result, err = f()
	}()
	return func() (interface{}, error) {
		<-c
		return result, err
	}
}

func Future2(f Function) Function {
	var result interface{}
	var err error
	var sig c
	select {
	case sig = <-signalList:
	default:
		sig = make(c, 1)
	}
	go func() {
		sig<-struct{}{}
		result, err = f()
	}()
	return func() (interface{}, error) {
		<-sig
		signalList<-sig
		return result, err
	}
}


func Future3(f Function) (cb *CallBack) {
	cb = &CallBack{
		task: f,
		fin: sync.WaitGroup{},
	}
	cb.fin.Add(1)
	go func() {
		cb.result, cb.err = cb.task()
		cb.fin.Done()
		//cb.fin<- struct{}{}
	}()
	return cb
}

type c chan struct{}

var signalList chan c

func init() {
	signalList = make(chan c, 2000)
	for i := 0; i < 2000; i++ {
		signalList<-make(c, 1)
	}
}

