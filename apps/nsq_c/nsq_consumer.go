package main

import (
	"github.com/nsqio/go-nsq"
	"fmt"
	"os/signal"
	"syscall"
	"os"
	"time"
)


const (
	topice_name = "fucker"
	channel_name = "mdzz"
)

func main() {

	config := nsq.NewConfig()

	fmt.Printf("%#v\n", config.MaxInFlight)
	config.MaxInFlight = 10
	c, _ := nsq.NewConsumer(topice_name, channel_name, config)

	c.AddConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		time.Sleep(2 * time.Second)
		fmt.Printf("C-%s\n", string(message.Body))
		return nil
	}), 25)

	err := c.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		panic(err)
	}
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)
	<-termChan
}


