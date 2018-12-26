package simple_impl

import (
	"testing"
	"sync"
	"strconv"
	"time"
)

var runTimes int = 20
var loopTimes int = 1e6

func noLoop() {
	time.Sleep(time.Millisecond*1)
	/*
	for i := 0; i < loopTimes; i++ {
	}
	*/
}

func BenchmarkNoLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noLoop()
	}
}

func funcGroupWithFuture3() []*CallBack {
	var callbackList = make([]*CallBack, 0, runTimes)
	for i := 0; i < runTimes; i++ {
		cb := Future3(func() (interface{}, error) {
			noLoop()
			return strconv.Itoa(i), nil
		})
		callbackList = append(callbackList, cb)
	}
	return callbackList
}

func funcGroupWithFuture2() []Function {
	var funcList = make([]Function, 0, runTimes)
	for i := 0; i < runTimes; i++ {
		f := Future2(func() (interface{}, error) {
			noLoop()
			return strconv.Itoa(i), nil
		})
		funcList = append(funcList, f)
	}
	return funcList
}

func funcGroupWithFuture() []Function {
	var funcList = make([]Function, 0, runTimes)
	for i := 0; i < runTimes; i++ {
		f := Future(func() (interface{}, error) {
			noLoop()
			return strconv.Itoa(i), nil
		})
		funcList = append(funcList, f)
	}
	return funcList
}

func funcGroupWithFuture4() []Function {
	var funcList = make([]Function, 0, runTimes)
	for i := 0; i < runTimes; i++ {
		f := Future4(func() (interface{}, error) {
			noLoop()
			return strconv.Itoa(i), nil
		})
		funcList = append(funcList, f)
	}
	return funcList
}

func funcGroupWithFuture5() []Function {
	var funcList = make([]Function, 0, runTimes)
	for i := 0; i < runTimes; i++ {
		f := Future5(func() (interface{}, error) {
			noLoop()
			return strconv.Itoa(i), nil
		})
		funcList = append(funcList, f)
	}
	return funcList
}

func funcGroupWithRawGO() {
	var funcList = make([]func() (interface{}, error), runTimes)
	var wg = sync.WaitGroup{}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		f := func() (interface{}, error) {
			defer wg.Done()
			noLoop()
			return "resp", nil
		}
		go f()
		funcList = append(funcList, f)
	}
	wg.Wait()
}

func doRPCCall() {
	result := funcGroupWithFuture()
	for _, fn := range result {
		fn()
	}
}

func doRPCCall2() {
	result := funcGroupWithFuture2()
	for _, fn := range result {
		fn()
	}
}

func doRPCCall3() {
	result := funcGroupWithFuture3()
	for _, cb := range result {
		_, _ = cb.Fin()
	}
}

func doRPCCall4() {
	result := funcGroupWithFuture4()
	for _, fn := range result {
		fn()
	}
}

func doRPCCall5() {
	result := funcGroupWithFuture5()
	for _, fn := range result {
		fn()
	}
}


func BenchmarkDoRPCCall(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		doRPCCall()
	}
	b.StopTimer()
}
func BenchmarkDoRPCCall2(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		doRPCCall2()
	}
	b.StopTimer()
}

func BenchmarkDoRPCCall3(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		doRPCCall3()
	}
	b.StopTimer()
}

func BenchmarkDoRPCCall4(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		doRPCCall4()
	}
	b.StopTimer()
}

func BenchmarkDoRPCCall5(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		doRPCCall5()
	}
	b.StopTimer()
}

func BenchmarkFuncGroup2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcGroupWithRawGO()
	}
}
