package main

import "fmt"

func main() {
	n := 7
	fmt.Printf("Factorial of %d is : %d \n", n, factorial(n))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return factorial(num-1) * num
}
