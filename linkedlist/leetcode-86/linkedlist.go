package leetcode_86

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
拆分当前传递的链表，对当前链表进行拆分
生成两个链表，第一个存放小于提供值的节点，
第二个存放大于提供值的节点。
根据传递的值，遍历链表，比较每个节点大小，并同时存放到两个链表中，然后拼接一下，返回即可。
*/
func Partition(head *ListNode, x int) *ListNode {
	dumy1 := &ListNode{Val: -1}
	dumy2 := &ListNode{Val: -1}
	p := head
	p1 := dumy1
	p2 := dumy2
	for {
		log.Info(p1 == dumy1)
		if p != nil {
			if p.Val >= x {
				p2.Next = p
				p2 = p2.Next
				log.Infof("我是当前判断p2的值:%d,接下来的节点:%v\n", p2.Val, p2.Next)
				fmt.Println(p1 == dumy1)
			} else {
				p1.Next = p
				p1 = p1.Next
				log.Infof("我是当前判断p1的值:%d,接下来的节点:%v\n", p1.Val, p1.Next)
				log.Infof("我是当前判断dumy1的值:%d,接下来的节点:%v\n", dumy1.Val, dumy1.Next)
				fmt.Println(p1 == dumy1)

			}
			temp := &ListNode{Val: -1}
			temp = p.Next
			log.Infof("我是当前替换temp的值:%v\n", p.Next)
			p.Next = nil
			p = temp
			log.Infof("我是替换结束的p的值:%v\n", p)
			fmt.Println()
		} else {
			break
		}
	}
	log.Infof("到这里，我是p1:%v", p1)
	log.Infof("到这里，我是dumy1:%v\n", dumy1)
	fmt.Println()
	log.Infof("到这里，我是p2:%v", p2)
	log.Infof("到这里，我是dumy2:%v", dumy2)
	fmt.Println()
	p1.Next = dumy2.Next
	return dumy1.Next
}
