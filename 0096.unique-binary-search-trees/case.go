package case0096

func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0] = 1
	nums[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= n; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}
