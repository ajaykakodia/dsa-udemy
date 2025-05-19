package main

import "fmt"

func main() {
	fmt.Println("starting new begging...")
	n := 10
	fmt.Printf("Sum of %d natural number: %d \n", n, sum(n))
}

func sum(num int) int {
	if num == 1 {
		return 1
	}
	return sum(num-1) + num
}
