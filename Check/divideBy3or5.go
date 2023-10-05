package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the Number:")
	fmt.Scanf("%d", &num)

	if num == 0 {
		fmt.Println("Error")
	} else if num%3 == 0 && num%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("NotFizzBuzz")
	}
}
