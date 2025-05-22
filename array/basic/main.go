package main

import "fmt"

func main() {
	arr := [5]int{}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	fmt.Println(arr)

	delete(&arr, 2)

	fmt.Println(arr)
}

func delete(arr *[5]int, index int) {
	if index < 0 && index >= len(arr) {
		return
	}

	for i := index; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = 0
}
