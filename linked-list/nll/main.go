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

func RDisplay(first *node) {
	if first == nil {
		fmt.Println("Nil")
		return
	}
	fmt.Printf("%d -> ", first.Data)
	RDisplay(first.Next)
}

func RInsert(cn *node, data, index int) *node {
	if index == 0 {
		cd := node{
			Data: data,
			Next: cn,
		}
		return &cd
	}
	cn.Next = RInsert(cn.Next, data, index-1)
	return cn
}

func (ll *linkedList) Insert(data, index int) {
	if index > ll.Len {
		return
	}
	ll.Len++
	cd := node{
		Data: data,
	}
	if index == 0 {
		cd.Next = ll.First
		ll.First = &cd
	} else {
		cn := ll.First
		for i := 0; i < index-1; i++ {
			cn = cn.Next
		}
		cd.Next = cn.Next
		cn.Next = &cd
	}
}

func (ll *linkedList) Count() int {
	cn := ll.First
	count := 0
	for cn != nil {
		count++
		cn = cn.Next
	}
	return count
}

func RCount(cn *node) int {
	if cn == nil {
		return 0
	}
	return 1 + RCount(cn.Next)
}

func main() {
	ll := createLinkedList([]int{3, 6, 4, 8, 10, 14, 17})
	ll.Display()
	// RDisplay(ll.First)
	ll2 := linkedList{}
	ll2.Insert(2, 0)
	ll2.Insert(3, 1)
	ll2.Insert(4, 2)
	ll2.Insert(5, 0)
	ll2.Insert(7, 2)
	ll2.Display()
	ll3 := linkedList{}
	ll3.First = RInsert(ll3.First, 2, 0)
	ll3.First = RInsert(ll3.First, 3, 1)
	ll3.First = RInsert(ll3.First, 5, 2)
	ll3.First = RInsert(ll3.First, 7, 3)
	ll3.First = RInsert(ll3.First, 1, 2)
	ll3.First = RInsert(ll3.First, 10, 3)
	ll3.Display()
	fmt.Println("Number of Node in Linked List 1:", ll.Count())
	fmt.Println("Number of Node in Linked List 3:", RCount(ll3.First))
}
