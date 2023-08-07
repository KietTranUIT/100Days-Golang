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

func swapNodes(currentNode *ListNode, preNode *ListNode, count int) *ListNode {
	if currentNode == nil {
		return nil
	}
	count++

	if count%2 == 0 {
		fmt.Println(currentNode.val)
		preNode.Next = swapNodes(currentNode.Next, currentNode, count)
		currentNode.Next = preNode
		fmt.Println(preNode.val, currentNode.val)
		return currentNode
	}

	fmt.Println("test: ", currentNode.val, currentNode.Next.val)

	currentNode = swapNodes(currentNode.Next, currentNode, count)
	return currentNode

}

func swapPairs(head *ListNode) *ListNode {
	swapNodes(head, nil, 0)
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
