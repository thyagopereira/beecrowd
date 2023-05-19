package leetcode

func majorityElement(nums []int) int {
	hash := make(map[int]int)
	maj := len(nums) / 2

	for i := 0; i <= len(nums); i++ {
		if hash[nums[i]] == maj {
			return nums[i]
		} else {
			hash[nums[i]] += 1
		}
	}

	return 0
}
