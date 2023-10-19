package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {}

func main() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(10 * time.Second): // Sau 10s in timed out
		fmt.Println("timed out")
	}
}
