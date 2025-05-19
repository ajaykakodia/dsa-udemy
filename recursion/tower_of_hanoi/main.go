package main

import "fmt"

func main() {
	n := 4
	toh(n, "A", "B", "C")
}

func toh(n int, t1, t2, t3 string) {
	if n == 1 {
		fmt.Printf("%s --> %s\n", t1, t3)
		return
	}
	toh(n-1, t1, t3, t2)
	fmt.Printf("%s --> %s\n", t1, t3)
	toh(n-1, t2, t1, t3)
}
