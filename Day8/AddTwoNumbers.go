package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

type List struct {
	Head, Tail *ListNode
}

func InitialList(ls List) List {
	ls.Head = nil
	ls.Tail = nil
	return ls
}

func NewListNode(val int) *ListNode {
	newNode := new(ListNode)

	if newNode == nil {
		return nil
	}

	newNode.Val = val
	newNode.Next = nil
	return newNode
}

func InsertListNode(ls List, val int) List {
	newNode := NewListNode(val)

	if newNode == nil {
		panic("newNode is nil")
	}

	if ls.Head == nil {
		ls.Head = newNode
		ls.Tail = newNode
	} else {
		ls.Tail.Next = newNode
		ls.Tail = newNode
	}

	return ls
}

func InputListNode() *ListNode {
	var n int
	fmt.Println("Enter number of elements: ")
	fmt.Scanf("%d", &n)

	var number int
	var ls List
	ls = InitialList(ls)

	fmt.Println("Enter value: ")
	for i:=0; i<n; i++ {
		fmt.Scanf("%d", &number)
		ls = InsertListNode(ls, number)
	}
	return ls.Head
}

func UpdateCarry(number int) {
	if number < 10 {
		carry = 0
		return
	}

	carry = number / 10
}

var carry int = 0
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	newNode := new(ListNode)
	if l1 == nil && l2 == nil {
		if carry != 0 {
			newNode.Val = 1
			newNode.Next = nil
			carry = 0
			return newNode
		}
		return nil
	}

	var l1Next, l2Next *ListNode

	if l1 != nil && l2 == nil {
		newNode.Val = l1.Val + 0 + carry
		l1Next = l1.Next
		l2Next = nil
	} else if l2 != nil && l1 == nil {
		newNode.Val = l2.Val + 0 + carry
		l2Next = l2.Next
		l1Next = nil
	} else {
		newNode.Val = l1.Val + l2.Val + carry
		l1Next = l1.Next
		l2Next = l2.Next
	}
	UpdateCarry(newNode.Val)
	newNode.Val %= 10
	newNode.Next = addTwoNumbers(l1Next, l2Next)
	return newNode
}

func main() {
	l1 := InputListNode()
	l2 := InputListNode()

	l3 := addTwoNumbers(l1, l2)
	fmt.Printf("Add two numbers is: ")
	for i:=l3; i!=nil; i=i.Next {
		fmt.Printf("%d ", i.Val)
	}
	fmt.Println()
}