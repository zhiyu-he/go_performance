package main

import (
	"math/rand"
	"time"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

// 用于模拟在CPU负载高的Case下, chan的调度会受到影响


const (
	loop_size = 1000000 // 100w
	goroutine_size = 1000
)
func  calculateSomething() {
	sum := 0
	for i := 0; i < loop_size; i++ {
		sum += rand.Int()
	}
}

type job struct {
	addTime int
}

var jobChan = make(chan *job)

func emitJob() {
	j := &job{addTime: time.Now().Nanosecond()}
	jobChan<-j
}

func InitCPUCost() {
	for i := 0; i < goroutine_size; i++ {
		go func() {
			interval := int(rand.Int31n(100)) + 1
			t := time.NewTicker(time.Duration(interval) * time.Millisecond)
			for {
				<-t.C
				calculateSomething()
			}
		}()
	}
}


// 模拟单机QPS, 4ms一个, 1s 250qps
func InitQPSEmiter() {
	go func() {
		t := time.NewTicker(time.Millisecond * time.Duration(4))
		for {
			<-t.C
			emitJob()
		}
	}()
}

func ConsumerJob() {
	for {
		select {
		case j := <-jobChan:
			delta := (time.Now().Nanosecond() - j.addTime) / 1000 / 1000
			if delta > 0 {
				fmt.Printf("delta-ms: %d\n", delta)
			}
		}
	}
}

func main() {
	InitCPUCost()

	go ConsumerJob()

	InitQPSEmiter()
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	<-termChan
}