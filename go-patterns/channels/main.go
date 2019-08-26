package main

import (
	"time"
)

func main() {

	length := 10

	ch := make(chan string, length)

	for i := 0; i < 10; i++ {
		go func() {
			ch <- "Hello World"
			time.Sleep(time.Second)
		}()
	}

}
