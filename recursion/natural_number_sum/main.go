package main

import "fmt"

func main() {
	n := 11
	fmt.Printf("Sum of %d natural number: %d \n", n, sum(n))
	fmt.Printf("Sum of %d natural number by simple formula: %d \n", n, naturalSumByFormula(n))
}

func sum(num int) int {
	if num == 1 {
		return 1
	}
	return sum(num-1) + num
}

func naturalSumByFormula(num int) int {
	return num * (num + 1) / 2
}
