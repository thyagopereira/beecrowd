package leetcode

func rotate(nums []int, k int) {

	k = k % len(nums)
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

	l, r = 0, k-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}

	l, r = k, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l, r = l+1, r-1
	}
}
