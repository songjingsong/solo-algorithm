package case1305

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {

	r1 := InOrder(root1)
	r2 := InOrder(root2)
	result := make([]int, 0)
	for len(r1) > 0 && len(r2) > 0 {
		if r1[0] > r2[0] {
			result = append(result, r2[0])
			r2 = r2[1:]
		} else {
			result = append(result, r1[0])
			r1 = r1[1:]
		}
	}
	if len(r1) != 0 {
		result = append(result, r1...)
	}

	if len(r2) != 0 {
		result = append(result, r2...)
	}
	return result
}

func InOrder(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	node := root
	result := make([]int, 0)

	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)

			node = node.Left
		}
		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			node = node.Right
		}
	}
	return result
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
