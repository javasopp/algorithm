package leetcode_26

// 力扣-26 删除数组中的重复项
// 解题思路
// 快慢指针
// 让快指针走，每次比较一次，如果快指针的位置不等于慢指针的位置，慢指针往前走一步，赋值快指针的值到慢指针
// 最后返回慢指针的位置+1 就是最终的位置。

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow
}
