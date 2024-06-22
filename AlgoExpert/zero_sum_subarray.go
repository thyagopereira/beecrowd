package main

func ZeroSumSubarray(nums []int) bool {
	aux := make(map[int]int)

	aux[0] = 1
	currentSum := 0
	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]
		if aux[currentSum] != 0 {
			return true
		} else {
			aux[currentSum] += 1
		}
	}

	return false
}
func main() {
	input := []int{-5, -5, 2, 3, -2}

	if !ZeroSumSubarray(input) {
		panic("Wrong answer.")
	}
}
