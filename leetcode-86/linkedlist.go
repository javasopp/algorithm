package leetcode_86

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
分解链表，根据特定的值，进行分解成2个链表
定义两个头结点，分别添加对应的值即可。
*/
func partition(head *ListNode, x int) *ListNode {
	dumy1 := &ListNode{Val: -1}
	dumy2 := &ListNode{Val: -1}
	p := head
	p1 := dumy1
	p2 := dumy2
	for {
		if p != nil {
			if p.Val > x {
				p2.Next = p
				p2 = p2.Next
			} else {
				p1.Next = p
				p1 = p1.Next
			}
			temp := &ListNode{Val: -1}
			temp = p.Next
			p.Next = nil
			p = temp
		} else {
			break
		}
	}
	p1.Next = dumy2.Next
	return dumy1.Next
}
