package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func displayList(node *ListNode) {
	schema := ""
	for node != nil {
		schema += fmt.Sprintf("(%d)-> ", node.Val)
		node = node.Next
	}
	fmt.Println(schema)
}

func search(node *ListNode, val int) int {
	pos := 0
	for node != nil {
		if node.Val == val {
			return pos
		} else {
			pos++
		}

		node = node.Next
	}

	return 0
}

func insert(node *ListNode, val int, after int) *ListNode {
	newNode := ListNode{
		Val: val,
	}

	for node != nil {
		if node.Val == after {
			newNode.Next = node.Next
			node.Next = &newNode
			break
		}
		node = node.Next
	}
	return node
}

func reverse(node *ListNode) *ListNode {

	// init pointers
	var prev, next *ListNode
	current := node

	for current != nil {
		// 0. Before changing the current.Next, store the next node in temp
		next = current.Next

		// 1. Reverse the direction
		current.Next = prev

		// 2. take step forward for prev & current
		// prev as curr and curr as next
		prev = current
		current = next
	}
	return prev

}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	reversedHalf := reverse(findMiddle(head).Next)

	for reversedHalf != nil {
		if reversedHalf.Val != head.Val {
			return false
		}
		reversedHalf = reversedHalf.Next
		head = head.Next
	}

	return true
}

func main() {
	node := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	// displayList(&node)
	// fmt.Println(search(&node, 3))
	// displayList(insert(&node, 4, 3))
	// displayList(reverse(&node))
	fmt.Println(isPalindrome(&node))
}
