package leetcode_27

// 力扣-27 移除元素
// 解题思路
// 类似于原地删除数组的操作。
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
