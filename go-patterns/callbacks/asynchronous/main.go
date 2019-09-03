package main

import (
	"fmt"
	"strings"
	"sync"
)

var wait sync.WaitGroup

func main() {
	wait.Add(1)
	toUppercaseAsync("Hello Callback!", func(v string) {
		toUppercaseAsync("Hello Callback!", func(v string) {
			fmt.Printf("Inner callback executed: %s\n", v)
			wait.Done()
		})
		fmt.Printf("Outer callback executed: %s\n", v)

	})
	fmt.Println("Waiting for async response...")
	wait.Wait()
}

func toUppercaseAsync(s string, f func(string)) {
	go func() {
		f(strings.ToUpper(s))
	}()
}
