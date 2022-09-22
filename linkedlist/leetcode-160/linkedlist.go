package leetcode_160

type ListNode struct {
	Val  int
	Next *ListNode
}

// 力扣-160 两个链表相交的节点。
// 解题思路:
// 链表 A & B
// p1 & p2 指针，分别在逻辑上遍历两个链表
// A 遍历完，p1 遍历 B
// B 遍历完，p2 遍历 A
// 循环遍历即可。

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}

// 新的思路
// 直接分别计算 两个链表的长度
// 让两个链表到达末尾节点的位置相同
// 开始让链表走，走到一致的时候，就是交点
func getIntersectionNodeLength(headA, headB *ListNode) *ListNode {
	p1 := headA
	p2 := headB
	lenA, lenB := 0, 0
	for p1 != nil {
		p1 = p1.Next
		lenA++
	}
	for p2 != nil {
		p2 = p2.Next
		lenB++
	}
	if lenA > lenB {
		// A 走 A-B的距离
		for i := 0; i < lenA-lenB; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			p2 = p2.Next
		}
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
