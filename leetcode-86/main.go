package leetcode_86

import "fmt"

func main() {
	//	3
	slice := []int{1, 4, 3, 2, 5, 2}
	p := &ListNode{Val: -1}
	for i := 0; i < 6; i++ {
		node := &ListNode{Val: slice[i]}
		p.Next = node
	}
	fmt.Println("132132")
}
