package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println("Hello")
			time.Sleep(time.Second)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Done!")

}
