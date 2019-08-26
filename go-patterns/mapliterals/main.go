package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		34.34, -75.34,
	},
	"Google": Vertex{
		54.67, -56.76,
	},
}

func main() {
	fmt.Println(m)
}
