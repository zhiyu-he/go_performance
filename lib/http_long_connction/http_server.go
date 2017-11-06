package main

import (
	"fmt"
	"io"
	"net/http"
)

func HandlePut(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Here: %v\n", req.RequestURI)
	io.WriteString(w, "something-else")
}

func main() {
	http.HandleFunc("/put", HandlePut)

	err := http.ListenAndServe("127.0.0.1:4102", nil)
	if err != nil {
		panic(err)
	}
}
