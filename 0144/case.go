package cae0144

func preorderTraversal(root *TreeNode) []int {
	var ans []int
	pre(root, &ans)
	return ans
}
func pre(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	*ans = append(*ans, root.Val)
	pre(root.Left, ans)
	pre(root.Right, ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
