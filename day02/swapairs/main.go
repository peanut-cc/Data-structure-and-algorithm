package main


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