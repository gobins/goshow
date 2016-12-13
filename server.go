package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var count int

func sayhelloWorld(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 10)
	count++
	response := strconv.Itoa(count) + " Hello World"
	fmt.Println(response)
	fmt.Fprintf(w, response) // send data to client side
}

func main() {
	http.HandleFunc("/", sayhelloWorld) // set router
	fmt.Println("HTTP Server Listening on port 9090")
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
