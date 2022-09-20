package main

import (
	leetcode_19 "algorithm/linkedlist/leetcode-19"
	"fmt"
)

func main() {
	//	3
	p1 := &leetcode_19.ListNode{Val: 1}
	p2 := &leetcode_19.ListNode{Val: 2}
	p3 := &leetcode_19.ListNode{Val: 3}
	p4 := &leetcode_19.ListNode{Val: 4}
	p5 := &leetcode_19.ListNode{Val: 5}
	p1.Next = p2
	p2.Next = p3
	p3.Next = p4
	p4.Next = p5
	p5.Next = nil
	fmt.Println("132132")
}
