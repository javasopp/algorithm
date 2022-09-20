package leetcode_876

// 力扣 876 查找链表重点
// 思路:
// 使用快慢指针，慢指针走一步，快指针走两步。
// 当快指针走完过后，那么慢指针的位置，就是最终的位置。

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
