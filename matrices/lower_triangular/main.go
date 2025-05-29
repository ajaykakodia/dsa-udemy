package main

import "fmt"

type lowerTriangularMatrix struct {
	arr  []int
	size int
}

func getIndex(row, col int) int {
	return row*(row+1)/2 + col
}

func NewLowerTriangularMatrix(size int) *lowerTriangularMatrix {
	return &lowerTriangularMatrix{
		arr:  make([]int, size*(size+1)/2),
		size: size,
	}
}

func (d *lowerTriangularMatrix) Get(row, col int) int {
	if row <= col {
		return d.arr[getIndex(row, col)]
	}
	return 0
}

func (d *lowerTriangularMatrix) Set(row, col, val int) {
	if col > row {
		return
	}
	d.arr[getIndex(row, col)] = val
}

func (d *lowerTriangularMatrix) Display() {
	for row := 0; row < d.size; row++ {
		for col := 0; col < d.size; col++ {
			if col > row {
				fmt.Printf("%d ", 0)
			} else {
				fmt.Printf("%d ", d.arr[getIndex(row, col)])
			}
		}
		fmt.Println("")
	}
}

func main() {
	dm := NewLowerTriangularMatrix(5)
	dm.Display()
	println("--------------------------")
	dm.Set(0, 0, 8)
	dm.Set(1, 0, 3)
	dm.Set(1, 1, 4)
	dm.Set(2, 0, 8)
	dm.Set(2, 2, 5)
	dm.Set(3, 3, 4)
	dm.Set(4, 0, 10)
	dm.Set(4, 2, 11)
	dm.Set(4, 4, 9)
	dm.Display()
	fmt.Println("Element :", dm.Get(2, 2))
}
