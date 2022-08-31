package leetcode_21

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
解题思路: 定义一个空节点，然后遍历两个链表，并同时对比每一次节点的大小。
然后分别添加到单独节点上，最后成为一个单独的链表。
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	p := dummy
	p1 := list1
	p2 := list2
	for {
		if p1 != nil && p2 != nil {
			if p1.Val > p2.Val {
				p.Next = p2
				p2 = p2.Next
			} else {
				p.Next = p1
				p1 = p1.Next
			}
			p = p.Next
		} else {
			break
		}
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy
}
