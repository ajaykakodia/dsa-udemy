package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{8, 20, 4, 7, 6, 3, 10, 5, 14, 2}
	index := linearSearch(arr, 7)
	if index == -1 {
		fmt.Println("Element not found ...")
	} else {
		fmt.Printf("Element %d found at index: %d\n", 7, index)
	}

	index = linearSearch(arr, 21)
	if index == -1 {
		fmt.Println("Element not found ...")
	} else {
		fmt.Printf("Element %d found at index: %d\n", 21, index)
	}

	ele := Get(arr, 5)
	fmt.Printf("Element %d found at index 5\n", ele)

	Set(arr, 5, 19)

	ele = Get(arr, 5)
	fmt.Printf("Element %d found at index 5\n", ele)

	ele = Max(arr)
	fmt.Printf("Element %d is max in array %v\n", ele, arr)

	avg := Avg(arr)
	fmt.Printf("Avg of array %v is %v\n", arr, avg)

	sort.Ints(arr)

	fmt.Println(arr)

	index = binarySearch(arr, 10)
	if index == -1 {
		fmt.Println("Element not found ...")
	} else {
		fmt.Printf("Element %d found at index: %d\n", 10, index)
	}

	index = binarySearch(arr, 9)
	if index == -1 {
		fmt.Println("Element not found ...")
	} else {
		fmt.Printf("Element %d found at index: %d\n", 9, index)
	}

	index = binarySearchR(arr, 5, 0, len(arr))
	if index == -1 {
		fmt.Println("Element not found ...")
	} else {
		fmt.Printf("Element %d found at index: %d\n", 5, index)
	}

	index = binarySearchR(arr, 9, 0, len(arr))
	if index == -1 {
		fmt.Println("Element not found ...")
	} else {
		fmt.Printf("Element %d found at index: %d\n", 15, index)
	}
}

func Get(arr []int, index int) int {
	if index < 0 || index >= len(arr) {
		return -1
	}
	return arr[index]
}

func Set(arr []int, index, val int) {
	if index < 0 || index >= len(arr) {
		return
	}
	arr[index] = val
}

func Max(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func Avg(arr []int) float32 {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return float32(sum) / float32(len(arr))
}

func linearSearch(arr []int, ele int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == ele {
			return i
		}
	}
	return -1
}

func binarySearch(arr []int, ele int) int {
	start := 0
	end := len(arr)

	for start <= end {
		mid := (end + start) / 2
		if arr[mid] == ele {
			return mid
		}
		if arr[mid] > ele {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

func binarySearchR(arr []int, ele, start, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if arr[mid] == ele {
		return mid
	}

	if arr[mid] > ele {
		end = mid - 1
	} else {
		start = mid + 1
	}
	return binarySearchR(arr, ele, start, end)
}
