package case0739

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		ans[i] = 0
	}
	return ans
}
