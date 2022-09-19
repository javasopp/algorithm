package main

import (
	leetcode_86 "algorithm/leetcode-86"
	"fmt"
)

func main() {
	//	3
	p1 := &leetcode_86.ListNode{Val: 1}
	p2 := &leetcode_86.ListNode{Val: 4}
	p3 := &leetcode_86.ListNode{Val: 3}
	p4 := &leetcode_86.ListNode{Val: 2}
	p5 := &leetcode_86.ListNode{Val: 5}
	p6 := &leetcode_86.ListNode{Val: 2}
	p1.Next = p2
	p2.Next = p3
	p3.Next = p4
	p4.Next = p5
	p5.Next = p6
	p6.Next = nil
	fmt.Println("132132")
	leetcode_86.Partition(p1, 3)
}
