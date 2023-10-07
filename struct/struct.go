package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := Vertex{2, 3}
	v.X = 4
	fmt.Println(v.X)
}
