package main

import "fmt"

func main() {
	var si, r, p, t int

	fmt.Println("Enter Principal=")
	fmt.Scanf("%d", &p)
	fmt.Println("Enter Rate=")
	fmt.Scanf("%d", &r)
	fmt.Println("Enter time in years=")
	fmt.Scanf("%d", &t)
	si = p * r * t / 100

	fmt.Println("Simple Interest=", si)
}
