package case0003

func lengthOfLongestSubstring(s string) int {
	ans, slide := 0, make([]int, 128)
	for i, j := 0, 0; j < len(s); j++ {
		i = max(i, slide[s[j]])
		ans = max(ans, j-i+1)
		slide[s[j]] = j + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
