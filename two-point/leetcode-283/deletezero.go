package leetcode_283

// 力扣-283
// 解题思路
// 移动所有的零值到末尾，然后也就是删除所有的零值
// 使用27题的零值
func moveZeroes(nums []int) {
	p := removeElement(nums, 0)
	for ; p < len(nums); p++ {
		nums[p] = 0
	}
}

func removeElement(nums []int, value int) int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != value {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
