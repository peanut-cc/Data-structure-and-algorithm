package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 做过第一个hascycle 这个链表题之后，其实再看这个基本也是一样的思路，哈希表和快慢指针
// 在这里同样是可用的， 哈希表在这里实现，同样复杂度为O(N)
func detectCycle(head *ListNode) *ListNode {
	hMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := hMap[head]; ok {
			return head
		}
		hMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}
