package main

import (
	"fmt"
	"strings"
)

func main() {

	toUppercaseCallback("Hello Callback!", func(v string) {
		fmt.Printf("Callback executed: %s\n", v)
	})
}

func toUppercaseCallback(s string, f func(string)) {
	f(strings.ToUpper(s))
}
