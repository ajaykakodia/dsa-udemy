package main

import "fmt"

func main() {
	num := 8
	m := make(map[int]int)
	fmt.Printf("Fibonacci of %d is: %d \n ", num, fab1(num))
	fmt.Printf("Fibonacci of %d is: %d \n ", num, fab2(num))
	fmt.Printf("Fibonacci of %d is: %d \n ", num, fab3(num, m))
}

func fab1(n int) int {
	result := 0
	num1, num2 := 0, 1

	for i := 2; i <= n; i++ {
		result = num1 + num2
		num1 = num2
		num2 = result
	}
	return result
}

func fab2(n int) int {
	if n <= 1 {
		return n
	}

	return fab2(n-1) + fab2(n-2)
}

func fab3(n int, m map[int]int) int {
	if n <= 1 {
		m[n] = n
		return n
	}
	if _, ok := m[n-1]; !ok {
		m[n-1] = fab3(n-1, m)
	}
	if _, ok := m[n-2]; !ok {
		m[n-2] = fab3(n-2, m)
	}
	return m[n-1] + m[n-2]
}
