package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	Head, Tail *ListNode
}

func Init(ls List) List {
	ls.Head = nil
	ls.Tail = nil
	return ls
}

func NewNode(val int) *ListNode {
	new_node := &ListNode{
		Val:  val,
		Next: nil,
	}

	if new_node == nil {
		return nil
	}

	return new_node
}

func Insert(node *ListNode, ls List) List {
	if ls.Head == nil {
		ls.Head = node
		ls.Tail = node
	} else {
		ls.Tail.Next = node
		ls.Tail = node
	}
	return ls
}

func PrintLinkList(head *ListNode) {
	for i := head; i != nil; i = i.Next {
		fmt.Printf("%d ", i.Val)
	}
	fmt.Println()
}

func InputList(n int, ls List) List {
	val := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &val)
		new_node := NewNode(val)
		ls = Insert(new_node, ls)
	}
	return ls
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	var result *ListNode = nil
	if list1 != nil && list2 == nil {
		result = list1
		result.Next = mergeTwoLists(list1.Next, nil)
	} else if list1 == nil && list2 != nil {
		result = list2
		result.Next = mergeTwoLists(nil, list2.Next)
	} else {
		if list1.Val < list2.Val {
			result = list1
			result.Next = mergeTwoLists(list1.Next, list2)
		} else {
			result = list2
			result.Next = mergeTwoLists(list1, list2.Next)
		}
	}
	return result
}

func main() {
	var ls1, ls2 List
	Init(ls1)
	Init(ls2)

	var m, n int
	fmt.Printf("Enter size of list 1: ")
	fmt.Scanf("%d", &m)
	fmt.Printf("Enter val for list 1: ")
	ls1 = InputList(m, ls1)

	fmt.Printf("Enter size of list 2: ")
	fmt.Scanf("%d", &n)
	fmt.Printf("Enter val for list 2: ")
	ls2 = InputList(n, ls2)

	head := mergeTwoLists(ls1.Head, ls2.Head)
	fmt.Printf("Output: ")
	PrintLinkList(head)

}
