package main

import "fmt"

func main() {
	msgs := make(chan string)
	go func() {msgs <- "ping"} ()
	msg := <-msgs
	fmt.Println(msg)
}