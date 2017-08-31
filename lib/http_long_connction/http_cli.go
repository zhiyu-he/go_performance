package main

import (
	"net/http"
//	"io/ioutil"
	"fmt"
	"net"
	"time"
	"io/ioutil"
	"io"
)

const URL = "http://127.0.0.1:4102"



func doPost(client *http.Client, url, topic string) (string, error) {
	uri := url + "/put?topic=" + topic

	fmt.Printf("URL: %v\n", uri)
	resp, err := client.Post(uri, "plain/text", nil)
	if err != nil {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	//io.Copy(ioutil.Discard,resp.Body)

	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil

}

func doPost2(client *http.Client, url, topic string) (string, error) {
	uri := url + "/put?topic=" + topic

	fmt.Printf("URL: %v\n", uri)
	resp, err := client.Post(uri, "plain/text", nil)
	if err != nil {
		return "", err
	}

	err = resp.Body.Close()
	return "", err
}


func main() {

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				conn, err  := net.DialTimeout(network, addr, 50 * time.Millisecond)
				fmt.Printf("LocalAddr: %v RemoteAddr: %v\n", conn.LocalAddr(), conn.RemoteAddr())
				return conn, err
			},
			MaxIdleConnsPerHost: 2, // 单机可以保持10个长链接, 默认为2个
		},
	}
	// 循环中没有Connection的创建, 说明进行了复用, test-case1

	for {
		go doPost(client, URL, "mq_ad1")
		go doPost(client, URL, "mq_ad2")
		time.Sleep(2 * time.Second)
	}


	/*
	// 由于server返回了msg, 所以这个消息cli不读的情况下, 就重新创建链接 test-case2
	for {
		go doPost2(client, URL, "mq_ad1")
		go doPost2(client, URL, "mq_ad2")
		time.Sleep(2 * time.Second)
	}
	*/

}