package main

import "fmt"

type diagonalMatrix struct {
	arr  []int
	size int
}

func NewDiagonalMatrix(size int) *diagonalMatrix {
	return &diagonalMatrix{
		arr:  make([]int, size),
		size: size,
	}
}

func (d *diagonalMatrix) Get(row, col int) int {
	if row > d.size && col > d.size {
		return 0
	}
	if row == col {
		return d.arr[row]
	}
	return 0
}

func (d *diagonalMatrix) Set(row, col, val int) {
	if row > d.size && col > d.size {
		return
	}
	if row == col {
		d.arr[row] = val
	}
}

func (d *diagonalMatrix) Display() {
	for i := 0; i < d.size; i++ {
		for j := 0; j < d.size; j++ {
			if j == i {
				fmt.Printf("%d ", d.arr[i])
			} else {
				fmt.Printf("%d ", 0)
			}
		}
		fmt.Println("")
	}
}

func main() {
	dm := NewDiagonalMatrix(5)
	dm.Display()
	println("--------------------------")
	dm.Set(0, 0, 8)
	dm.Set(1, 1, 3)
	dm.Set(2, 2, 5)
	dm.Set(3, 3, 4)
	dm.Set(4, 4, 9)
	dm.Display()
	fmt.Println("Element :", dm.Get(2, 2))
}
