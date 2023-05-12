package leetcode

func tribonacci(n int) int {
	dp := []int{0, 1, 1}

	for i := 3; i <= n; i++ {
		trib := dp[i-1] + dp[i-2] + dp[i-3]
		dp = append(dp, trib)
	}

	return dp[n]
}
