package leetcode

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	var ans []int
	for i <= j {
		if numbers[i]+numbers[j] == target {
			ans = append(ans, i+1)
			ans = append(ans, j+1)
			break
		} else if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}

	return ans
}
