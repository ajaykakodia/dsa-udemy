package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Println(&arr[i], "---", v)
	}
}
