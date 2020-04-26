package case0938

func rangeSumBST(root *TreeNode, L int, R int) int {
	var ans int
	dfs(root, L, R, &ans)
	return ans
}

func dfs(root *TreeNode, L int, R int, ans *int) {
	if root == nil {
		return
	}
	if root.Val >= L && root.Val <= R {
		*ans += root.Val
	}
	if root.Val > L {
		dfs(root.Left, L, R, ans)
	}
	if root.Val < R {
		dfs(root.Right, L, R, ans)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
