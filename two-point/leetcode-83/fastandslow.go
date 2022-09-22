package leetcode_83

// 力扣-83 删除数组中重复元素
// 解题思路
// 和快慢指针删除数组类似

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
