package main

import (
	"fmt"
	"testing"
	"time"
)

func TestStartInstance(t *testing.T) {
	singleton := GetInstance()
	singletonTwo := GetInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singletonTwo.AddOne()
	}

	fmt.Printf("Before loop current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}

	singleton.Stop()
}
