package main

import (
	"fmt"
	"time"
)

func main() {

	hello := make(chan string)
	goodbye := make(chan string)
	quit := make(chan bool)

	go func(s string) {
		time.Sleep(time.Second * 2)
		hello <- s
	}("Hello!")

	go func(s string, ch chan string) {
		time.Sleep(time.Second * 4)
		ch <- s
	}("Goodbye!", goodbye)

	go func(b bool) {
		time.Sleep(time.Second * 5)
		quit <- b
	}(true)

	for {
		select {
		case msg := <-hello:
			fmt.Printf("Hello Channel: %v\n", msg)
		case msg := <-goodbye:
			fmt.Printf("goodbye Channel: %v\n", msg)
		case <-time.After(time.Second * 7):
			fmt.Println("Deadline reached!")
		case <-quit:
			fmt.Println("Quitting!")
			break
		}
	}

	close(hello)
	close(goodbye)
	close(quit)

}
