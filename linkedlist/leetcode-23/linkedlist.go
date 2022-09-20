package leetcode_23

import "sort"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
第一个简单思路，把所有的节点都遍历添加到slice中，然后对slice进行排序，最后拼接出来所有的节点。
就是最后结果，耗时很长，不是好办法。
*/
func mergeKLists(lists []*ListNode) *ListNode {
	result := make([]*ListNode, 0)
	for i := 0; i < len(lists); i++ {
		tmp := lists[i]
		for tmp != nil {
			result = append(result, tmp)
			tmp = tmp.Next
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Val < result[j].Val
	})
	for i := 0; i < len(result)-1; i++ {
		result[i].Next = result[i+1]
	}
	if len(result) == 0 {
		return nil
	}
	result[len(result)-1].Next = nil
	return result[0]
}

/*
采用分治的办法，但是这种目前还看不懂思路。
*/
func mergeKListsDivideAndConquer(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	l1 := mergeKLists(lists[:len(lists)/2])
	l2 := mergeKLists(lists[len(lists)/2:])
	result := &ListNode{}
	tmp := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	} else {
		tmp.Next = l2
	}
	return result.Next
}
