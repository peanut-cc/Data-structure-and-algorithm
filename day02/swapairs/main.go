package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 这种方法其实不难，主要是对于链表这种数据结构操作需要熟悉，这几天要多看几遍加强这种思维训练方式
func swapPairs(head *ListNode) *ListNode {
	// 这里声明一个临时ListNode{} 对象，并且它 的Next为Head
	pre := &ListNode{}
	pre.Next = head
	temp := pre
	// 这里的判断其实就是为了保证每次两两交换的节点是存在的
	for temp.Next != nil && temp.Next.Next != nil {
		start := temp.Next
		end := temp.Next.Next
		temp.Next = end
		start.Next = end.Next
		end.Next = start
		temp = start
	}
	return pre.Next
}

// 重复练习一遍
func swapPairs2(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return pre.Next
}

// 再次重复练习一遍
func swapPairs3(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return pre.Next
}

// 自己默写一遍，加强理解
func swapPairs4(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return pre.Next
}

func swapPairs5(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		end.Next = start
		start.Next = end.Next
		tmp = start
	}
	return pre.Next
}

// 周一重复练习
func swapPairs6(head *ListNode) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	tmp := pre
	for tmp.Next != nil && tmp.Next.Next != nil {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return pre.Next
}

// 用递归的思路实现
// 这种递归的实现其实是从后往前两两反转，需要多写写，初看代码不是非常容易理解
func swapPairsNew(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsNew(next.Next)
	next.Next = head
	return next
}

func swapPairsNew2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsNew2(next.Next)
	next.Next = head
	return next
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
	fmt.Printf("%#v\n", origin)
	// for origin != nil {
	// 	fmt.Printf("%#v\n", origin.Val)
	// 	if origin.Next == nil {
	// 		break
	// 	}
	// 	origin = origin.Next
	// }
	res := swapPairsNew(origin)
	for res != nil {
		fmt.Printf("%#v\n", res.Val)
		if res.Next == nil {
			break
		}
		res = res.Next
	}
}
