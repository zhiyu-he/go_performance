package main

import (
	"github.com/nsqio/go-nsq"
	"time"
	"strconv"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const (
	topice_name = "fucker"
)

func main() {

	config := nsq.NewConfig()

	writer1, _ :=nsq.NewProducer("127.0.0.1:23333", config)
	writer2, _ := nsq.NewProducer("127.0.0.1:4150", config)
	t1 := time.NewTicker(time.Millisecond * 500)
	t2 := time.NewTicker(time.Second * 2)
	go SendMsg(writer1, t1)
	go SendMsg(writer2, t2)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	<-termChan
}

func SendMsg(p *nsq.Producer, ticket *time.Ticker) {
	times := 0
	for {
		<-ticket.C
		msg := strconv.Itoa(times) + "-p"
		fmt.Printf("MSG: %s\n", msg)
		err := p.Publish(topice_name, []byte(msg))
		if err != nil {
			fmt.Printf("%#v\n", err)
		}
		times += 1
	}
}
