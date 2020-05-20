package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 这是一种非常常规的思路，也是应该立马想起来的答案
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTmp
	}
	return curr
}

// 这个思路非常有意思，需要好好理解
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// p 就是原来链表的最后一个
	p := reverseList2(head.Next)
	fmt.Println(p.Val)
	fmt.Println(head.Val)
	head.Next.Next = head
	head.Next = nil
	fmt.Println(head.Val)
	return p
}
}

func main() {
	var curr *ListNode
	var origin *ListNode
	for i := 1; i <= 4; i++ {
		one := &ListNode{
			Val: i,
		}
		if i == 1 {
			origin = one
		}
		if curr != nil {
			curr.Next = one
		}
		curr = one
	}
	//fmt.Println(origin)
	res := reverseList2(origin)
	fmt.Println(res)
}