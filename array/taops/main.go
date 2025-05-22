package main

import "fmt"

func main() {
	arr1 := []int{3, 5, 7, 9, 10, 15, 17, 19}
	arr2 := []int{2, 4, 6, 8, 10, 11, 14, 15, 16, 18, 20, 22}
	res := merge(arr1, arr2)
	fmt.Println(res)

	res = union(arr1, arr2)
	fmt.Println(res)

	arr3 := []int{1, 5, 6, 8, 10, 14, 22, 23}
	res = unionSort(arr1, arr3)
	fmt.Println(res)
	res = intersect(arr1, arr2)
	fmt.Println(res)
	res = intersectSort(arr1, arr3)
	fmt.Println(res)

	res = difference(arr1, arr2)
	fmt.Println(res)
	res = differenceSort(arr1, arr3)
	fmt.Println(res)
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))

	i, j, k, l, m := 0, 0, len(arr1), len(arr2), 0

	for i < k && j < l {
		if arr1[i] < arr2[j] {
			result[m] = arr1[i]
			i++
			m++
		} else {
			result[m] = arr2[j]
			j++
			m++
		}
	}

	for i < k {
		result[m] = arr1[i]
		m++
		i++
	}

	for j < l {
		result[m] = arr2[j]
		m++
		j++
	}

	return result
}

func difference(arr1, arr2 []int) []int {
	res := []int{}

	m := make(map[int]bool)
	for i := 0; i < len(arr2); i++ {
		m[arr2[i]] = true
	}

	for i := 0; i < len(arr1); i++ {
		if !m[arr1[i]] {
			res = append(res, arr1[i])
		}
	}

	return res
}

func differenceSort(arr1, arr2 []int) []int {
	res := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			j++
		} else {
			i++
			j++
		}
	}

	for ; i < len(arr1); i++ {
		res = append(res, arr1[i])
	}

	return res
}

func intersect(arr1, arr2 []int) []int {
	res := []int{}

	m := make(map[int]bool)
	for i := 0; i < len(arr1); i++ {
		m[arr1[i]] = true
	}

	for i := 0; i < len(arr2); i++ {
		if m[arr2[i]] {
			res = append(res, arr2[i])
		}
	}

	return res
}

func intersectSort(arr1, arr2 []int) []int {
	res := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			i++
		} else if arr2[j] < arr1[i] {
			j++
		} else {
			res = append(res, arr1[i])
			i++
			j++
		}
	}

	return res
}

func union(arr1, arr2 []int) []int {
	res := []int{}

	m := make(map[int]bool)
	for i := 0; i < len(arr1); i++ {
		res = append(res, arr1[i])
		m[arr1[i]] = true
	}

	for i := 0; i < len(arr2); i++ {
		if !m[arr2[i]] {
			res = append(res, arr2[i])
		}
	}

	return res
}

func unionSort(arr1, arr2 []int) []int {
	res := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else if arr2[j] < arr1[i] {
			res = append(res, arr2[j])
			j++
		} else {
			res = append(res, arr1[i])
			i++
			j++
		}
	}

	for ; i < len(arr1); i++ {
		res = append(res, arr1[i])
	}

	for ; j < len(arr2); j++ {
		res = append(res, arr2[j])
	}

	return res
}
