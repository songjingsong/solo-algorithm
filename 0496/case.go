package case0496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int, 0)
	stack := make([]int, 0)
	for _, v := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < v {
			dict[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}
		dict[v] = -1
		stack = append(stack, v)
	}
	ans := make([]int, 0)
	for _, v := range nums1 {
		ans = append(ans, dict[v])
	}
	return ans
}
