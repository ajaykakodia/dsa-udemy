package main

import "fmt"

// nCr - n!/(r!*(n-r)!)
func main() {
	n, r := 5, 3

	fmt.Printf("Combination of %d char in total %d char is: %d \n", r, n, comb1(n, r))
	fmt.Printf("Combination of %d char in total %d char is: %d \n", r, n, comb(n, r))
}

// by pascal triangle solution
func comb(n, r int) int {
	if r == 0 || n == r {
		return 1
	}

	return comb(n-1, r-1) + comb(n-1, r)
}

// simple formula
func comb1(n, r int) int {
	t1 := fact(n)
	t2 := fact(r)
	t3 := fact(n - r)

	return t1 / (t2 * t3)
}

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return fact(num-1) * num
}
