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

func (ll *linkedList) Delete(index int) {
	if index >= ll.Len || ll.First == nil {
		return
	}
	ll.Len--

	if index == 0 {
		ll.First = ll.First.Next
	} else {
		var pn *node
		cn := ll.First
		for i := 0; i < index; i++ {
			pn = cn
			cn = cn.Next
		}
		pn.Next = cn.Next
	}
}

func RDelete(cn *node, index int) *node {
	if index == 0 {
		return cn.Next
	}
	if cn.Next != nil {
		cn.Next = RDelete(cn.Next, index-1)
	}
	return cn
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

func (ll *linkedList) Search(item int) int {
	cn := ll.First
	pos := 0
	for cn != nil {
		if cn.Data == item {
			return pos
		}
		cn = cn.Next
		pos++
	}
	return -1
}

func (ll *linkedList) Sum() int {
	cn := ll.First
	sum := 0
	for cn != nil {
		sum += cn.Data
		cn = cn.Next
	}
	return sum
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
	RDisplay(ll3.First)
	fmt.Println("Number of Node in Linked List 1:", ll.Count())
	fmt.Println("Number of Node in Linked List 3:", RCount(ll3.First))
	ll.Delete(4)
	ll.Display()
	ll.Delete(5)
	ll.Display()
	ll.Delete(5)
	ll.Display()
	ll3.First = RDelete(ll3.First, 2)
	ll3.Display()
	ll3.First = RDelete(ll3.First, 4)
	ll3.Display()
	ll3.First = RDelete(ll3.First, 5)
	ll3.Display()
	fmt.Println("Number 5 found in ll1: ", ll.Search(8) != -1)
	fmt.Println("Number 17 found in ll3: ", ll3.Search(17) != -1)
	fmt.Println("Sum of all element of ll2: ", ll2.Sum())
}
