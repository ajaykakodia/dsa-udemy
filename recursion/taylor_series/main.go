package main

import "fmt"

// What is Taylor series - exponent - e^x = 1+ x/1+ x^2/2! + x^3/3! + x^4/4! + ........ + x^n/n!
func main() {
	m, n := 3, 15
	fmt.Printf("%d exponent %d is: %v \n", m, n, taylorSeries1(m, n))
	fmt.Printf("%d exponent %d is: %v \n", m, n, taylorSeries2(m, n))
	fmt.Printf("%d exponent %d is: %v \n", m, n, taylorSeries3(m, n))
	fmt.Printf("%d exponent %d is: %v \n", m, n, taylorSeries4(m, n))
}

var ts float32 = 1

// Horner's Rule
func taylorSeries3(m, n int) float32 {
	if n == 0 {
		return ts
	}
	ts = 1 + (float32(m)/float32(n))*ts
	return taylorSeries3(m, n-1)
}

// looping
func taylorSeries4(m, n int) float32 {
	result := float32(1)
	for ; n > 0; n-- {
		result = 1 + (float32(m)/float32(n))*result
	}
	return result
}

var (
	pw  int = 1
	fac int = 1
)

func taylorSeries2(m, n int) float32 {
	if n == 0 {
		return 1
	}
	result := taylorSeries2(m, n-1)
	pw = m * pw
	fac = fac * n
	return result + float32(pw)/float32(fac)
}

func taylorSeries1(m, n int) float32 {
	if n == 0 {
		return 1
	}
	p := pow(m, n)
	f := fact(n)
	return taylorSeries1(m, n-1) + float32(p)/float32(f)
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

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return fact(num-1) * num
}
