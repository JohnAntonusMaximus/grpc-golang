package main

import "time"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		time.Sleep(time.Second * 1)
		ch <- 2
		close(ch)
	}()

	for v := range ch {
		println(v)
	}
}
