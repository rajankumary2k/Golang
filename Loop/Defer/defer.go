package main

import "fmt"

func main() {
	i := 0

	for i = 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
}
