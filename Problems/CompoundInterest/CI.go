package main

import (
	"fmt"
	"math"
)

func main() {

	var a, r, t, p, ci float64

	fmt.Println("Enter Principal=")
	fmt.Scanf("%g", &p)
	fmt.Println("Enter Rate=")
	fmt.Scanf("%g", &r)
	fmt.Println("Enter time in years=")
	fmt.Scanf("%g", &t)

	a = p * math.Pow((1+r/100), t)
	ci = a - p
	fmt.Println("Amount= ", a)

	fmt.Println("Compound Interest=", ci)
}
