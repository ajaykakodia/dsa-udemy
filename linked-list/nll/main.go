package main

import "fmt"

type node struct {
	Data int
	Next *node
}

type linkedList struct {
	First *node
	Len   int
}

func createLinkedList(arr []int) *linkedList {
	ll := linkedList{}
	var cNode *node
	for i := 0; i < len(arr); i++ {
		ll.Len++
		cn := node{
			Data: arr[i],
		}
		if ll.First != nil {
			cNode.Next = &cn
			cNode = &cn
			continue
		}
		ll.First = &cn
		cNode = &cn
	}
	return &ll
}

func (ll *linkedList) Display() {
	cn := ll.First

	for cn != nil {
		fmt.Printf("%d -> ", cn.Data)
		cn = cn.Next
	}

	fmt.Println("Nil")
}

func main() {
	ll := createLinkedList([]int{3, 6, 4, 8, 10, 14, 17})
	ll.Display()
}
