package main

import "fmt"

func main() {

	var num1, num2, sum int

	fmt.Println("Enter number1:")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter number2:")
	fmt.Scanf("%d", &num2)

	sum = num1 + num2

	fmt.Println("sum is ", sum)

}
