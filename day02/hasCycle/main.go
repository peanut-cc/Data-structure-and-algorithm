package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 哈希表方式
// 这种方式是比较容易想到的，也是比较容易理解的
// 这种的时间复杂度是O(N)
func hasCycle(head *ListNode) bool {
	aMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := aMap[head]; ok {
			return true
		}
		aMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 重复练习增加理解
func hasCycleHash(head *ListNode) bool {
	aMap := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := aMap[head]; ok {
			return true
		}
		aMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 快慢指针的实现，这是一种非常好的思路
// 这种实现方式也非常容易理解
// 这种实现的复杂度也是O(N)
func hasCycleSlowFast(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || slow == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

// 重复练习增加理解
func hasCycleSlowFast2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for slow != fast {
		if slow == nil || fast == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}




