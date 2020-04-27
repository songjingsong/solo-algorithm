package case0404

func sumOfLeftLeaves(root *TreeNode) int {
	ans := 0
	dfs(root, false, &ans)
	return ans
}

func dfs(root *TreeNode, isleft bool, ans *int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && isleft {
		*ans += root.Val
		return
	}
	dfs(root.Left, true, ans)
	dfs(root.Right, false, ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
