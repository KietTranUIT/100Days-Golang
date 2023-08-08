package main

import "fmt"

type List struct {
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	val  int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	new_node := &ListNode{
		val:  val,
		Next: nil,
	}

	return new_node
}

func Push(new_node *ListNode, ls List) List {
	if ls.Head == nil && ls.Tail == nil {
		ls.Head = new_node
		ls.Tail = new_node
	} else {
		ls.Tail.Next = new_node
		ls.Tail = new_node
	}
	return ls
}

func swapData(a, b int) (int, int) {
	return b, a
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	head.val, head.Next.val = swapData(head.val, head.Next.val)
	fmt.Println("Test: ", head.val, head.Next.val)
	head.Next.Next = swapPairs(head.Next.Next)
	return head
}

func PrintListNode(head *ListNode) {
	for i := head; i != nil; i = i.Next {
		fmt.Printf("%d ", i.val)
	}
	fmt.Println()
}

func main() {
	var n, val int
	ls := List{
		Head: nil,
		Tail: nil,
	}

loop:
	fmt.Printf("Enter size of list node: ")
	fmt.Scanf("%d", &n)

	if n < 0 || n > 100 {
		fmt.Println("Please enter again size of list node!")
		goto loop
	}

	fmt.Printf("Enter list value:  ")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &val)
		new_node := NewNode(val)
		ls = Push(new_node, ls)
	}

	fmt.Printf("Before: ")
	PrintListNode(ls.Head)

	ls.Head = swapPairs(ls.Head)

	fmt.Printf("After: ")
	PrintListNode(ls.Head)
}
