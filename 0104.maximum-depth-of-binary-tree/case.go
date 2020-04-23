package case0104

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftd := maxDepth(root.Left) + 1
	rightd := maxDepth(root.Right) + 1

	if leftd > rightd {
		return leftd
	}
	return rightd
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
