package case0965

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root, root.Val)
}
func dfs(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	return dfs(root.Left, val) && dfs(root.Right, val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
