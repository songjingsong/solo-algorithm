package case1019

func nextLargerNodes(head *ListNode) []int {
	stack, ans, data := make([]int, 0), make([]int, 0), make([]int, 0)
	for head != nil {
		ans = append(ans, 0)
		data = append(data, head.Val)

		for len(stack) > 0 && head.Val > data[stack[len(stack)-1]] {
			ans[stack[len(stack)-1]] = head.Val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, len(data)-1)
		head = head.Next
	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}
