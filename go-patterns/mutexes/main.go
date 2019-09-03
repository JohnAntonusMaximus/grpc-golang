package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	value int
	sync.Mutex
}

// you can run this program with the -race flag (go run -race main.go)
// to detect a race condition. Comment out the locks and rerun the program
// to see the race detector at runtime

func main() {
	counter := Counter{}

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer counter.Unlock()
			counter.Lock()
			counter.value++
			fmt.Printf("Value: %v\n", counter.value)
		}(i)
	}
	time.Sleep(time.Second)
	defer counter.Unlock()
	counter.Lock()
	println(counter.value)
}
