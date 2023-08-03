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

func lenLinkList(head *ListNode) (int, *ListNode) {
	count := 0
	var tail *ListNode
	for i := head; i != nil; i = i.Next {
		if i.Next == nil {
			tail = i
		}
		count++
	}
	return count, tail
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sz, tail := lenLinkList(head)
	index := sz - n
	k := 0
	node := head
	var pr *ListNode = nil

	for node != nil {
		if k == index {
			if node == head && sz == 1 {
				head = nil
				return head
			} else if node == head {
				head = node.Next
				return head
			} else if node == tail {
				pr.Next = nil
				tail = pr
				return head
			} else {
				pr.Next = node.Next
				node.Next = nil
				return head
			}
		}
		pr = node
		node = node.Next
		k++
	}
	return head
}

func PrintLinkList(head *ListNode) {
	for i := head; i != nil; i = i.Next {
		fmt.Printf("%d ", i.Val)
	}
	fmt.Println()
}

func main() {
	var ls List
	Init(ls)

	var s int
	fmt.Printf("Enter size of linked list: ")
	fmt.Scanf("%d", &s)

	var val int

	fmt.Printf("Enter vals: ")
	for i := 0; i < s; i++ {
		fmt.Scanf("%d", &val)
		new_node := NewNode(val)
		ls = Insert(new_node, ls)
	}
	fmt.Println("Before: ")
	PrintLinkList(ls.Head)

	var n int
	fmt.Printf("Enter n: ")
	fmt.Scanf("%d", &n)

	ls.Head = removeNthFromEnd(ls.Head, n)
	fmt.Println("After: ")
	PrintLinkList(ls.Head)
}
