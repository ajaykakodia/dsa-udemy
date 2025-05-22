package main

import "fmt"

func main() {
	arr1 := [5]int{}
	fmt.Println(arr1)
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i + 1
	}
	fmt.Println(arr1)

	delete(&arr1, 2)

	fmt.Println(arr1)

	arr := []int{5, 3, 2, 6, 8, 10, 13, 23, 17, 21}
	fmt.Println(arr)
	reverse(arr)
	fmt.Println(arr)
	LShift(arr, 3)
	fmt.Println(arr)
	arr = []int{5, 3, 2, 6, 8, 10, 13, 23, 17, 21}
	fmt.Println(arr)
	LShift(arr, 8)
	fmt.Println(arr)
	arr = []int{5, 3, 2, 6, 8, 10, 13, 23, 17, 21}
	fmt.Println(arr)
	RShift(arr, 3)
	fmt.Println(arr)
	arr = []int{5, 3, 2, 6, 8, 10, 13, 23, 17, 21}
	fmt.Println(arr)
	RShift(arr, 8)
	fmt.Println(arr)
}

func reverse(arr []int) {
	i, j, temp := 0, len(arr)-1, 0
	for i < j {
		temp = arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
}

func LShift(arr []int, shiftBy int) {
	if shiftBy > len(arr) {
		return
	}
	j := 0
	for i := shiftBy; i < len(arr); i++ {
		arr[j] = arr[i]
		j++
	}

	for i := 0; i < shiftBy; i++ {
		arr[len(arr)-1-i] = 0
	}
}

func RShift(arr []int, shiftBy int) {
	if shiftBy > len(arr) {
		return
	}
	j := len(arr) - 1
	for i := len(arr) - shiftBy - 1; i >= 0; i-- {
		arr[j] = arr[i]
		j--
	}

	for i := 0; i < shiftBy; i++ {
		arr[i] = 0
	}
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
