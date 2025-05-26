package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11, 12}
	missingNumber := findMissingElements(arr)
	fmt.Println("Missing Number:", missingNumber)

	missingNumber = findMissingElements1(arr)
	fmt.Println("Missing Number:", missingNumber)

	arr = []int{6, 7, 8, 9, 10, 11, 13, 14, 17, 18, 19}
	res := findMissingElements2(arr)
	fmt.Println("Missing Numbers:", res)

	arr = []int{3, 6, 8, 8, 10, 12, 15, 15, 15, 20}
	res = findDuplicateElementInSortedArray(arr)
	fmt.Println("Duplicate Numbers:", res)
	arr = []int{8, 3, 6, 4, 6, 5, 6, 8, 2, 7}
	res = findDuplicateElementInUnSortedArray(arr)
	fmt.Println("Duplicate Numbers:", res)
	arr = []int{8, 3, 6, 4, 6, 5, 6, 8, 2, 7}
	res = findDuplicateElementInUnSortedArray1(arr)
	fmt.Println("Duplicate Numbers:", res)

	arr = []int{1, 3, 4, 5, 6, 8, 9, 10, 12, 14}
	sortedSumArray(arr, 10)

	arr = []int{6, 3, 8, 10, 16, 7, 5, 2, 9, 14}
	sumPair(arr, 10)
	sumPair1(arr, 10)

	max, min := maxNMinInSingleRun(arr)
	fmt.Printf("max: %v, min: %v in array: %v", max, min, arr)
}

func maxNMinInSingleRun(arr []int) (int, int) {
	max, min := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		} else if arr[i] < min {
			min = arr[i]
		}
	}

	return max, min
}

func sumPair(arr []int, sum int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if val, ok := m[arr[i]]; ok {
			fmt.Printf("Sum of %v at index %v and %v at index %v is %v\n", arr[val], val, arr[i], i, sum)
		} else {
			m[sum-arr[i]] = i
		}
	}
}

func sumPair1(arr []int, sum int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == sum {
				fmt.Printf("Sum of %v at index %v and %v at index %v is %v\n", arr[i], i, arr[j], j, sum)
			}
		}
	}
}

func sortedSumArray(arr []int, sum int) {
	i, j := 0, len(arr)-1
	for i < j {
		if arr[i]+arr[j] > sum {
			j--
		} else if arr[i]+arr[j] < sum {
			i++
		} else {
			fmt.Printf("Sum of %v at index %v and %v at index %v is %v\n", arr[i], i, arr[j], j, sum)
			i++
			j--
		}
	}
}

func findMissingElements(arr []int) int {
	lastEle := arr[len(arr)-1]

	sumOfNNaturalNumber := lastEle * (lastEle + 1) / 2
	sumOfArray := 0
	for i := 0; i < len(arr); i++ {
		sumOfArray = sumOfArray + arr[i]
	}

	return sumOfNNaturalNumber - sumOfArray
}

func findMissingElements1(arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i]-i > 1 {
			return arr[i] - 1
		}
	}

	return -1
}

func findMissingElements2(arr []int) []int {
	indexGap := arr[0]
	res := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i]-i > indexGap {
			missingNumbers := arr[i] - i - indexGap
			for j := 0; j < missingNumbers; j++ {
				res = append(res, arr[i]-j-1)
			}
			indexGap = arr[i] - i
		}
	}

	return res
}

func findDuplicateElementInSortedArray(arr []int) []int {
	res := []int{}
	for i := 1; i < len(arr); i++ {
		if arr[i-1] == arr[i] && (len(res) == 0 || res[len(res)-1] != arr[i]) {
			res = append(res, arr[i])
		}
	}
	return res
}

func findDuplicateElementInUnSortedArray(arr []int) []int {
	res := []int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] != -1 {
			for j := i + 1; j < len(arr); j++ {
				if arr[i] == arr[j] && (len(res) == 0 || res[len(res)-1] != arr[i]) {
					res = append(res, arr[i])
					arr[j] = -1
				}
			}
		}
	}
	return res
}

func findDuplicateElementInUnSortedArray1(arr []int) []int {
	res := []int{}
	m := make(map[int]int)
	for i := 0; i < len(arr)-1; i++ {
		m[arr[i]]++
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
