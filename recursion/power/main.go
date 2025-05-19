package main

import "fmt"

func main() {
	m, n := 3, 13
	fmt.Printf("%d Power %d is: %d \n", m, n, power(m, n))
	fmt.Printf("%d Power %d is: %d \n", m, n, pow(m, n))
	m, n = 4, 8
	fmt.Printf("%d Power %d is: %d \n", m, n, power(m, n))
	fmt.Printf("%d Power %d is: %d \n", m, n, pow(m, n))
}

func power(m, n int) int {
	if n == 0 {
		return 1
	}
	return power(m, n-1) * m
}

func pow(m, n int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return pow(m*m, n/2)
	}
	return m * pow(m*m, (n-1)/2)
}
