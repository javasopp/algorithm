package leetcode_19

// 力扣 19 倒数第k个数
// 解题思路:
// 两个指针 p1,p2，先是p1 走n步，然后p2 开始从头走，p1 走完了过后，p2就是你需要找到的地方

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumy := &ListNode{Val: -1}
	dumy.Next = head
	result := findNthEnd(dumy, n+1)
	result.Next = result.Next.Next
	return dumy.Next
}

func findNthEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		if p1 != nil {
			p1 = p1.Next
		}
	}
	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}
