package case0111

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil { //已经到达叶子节点
		return 1
	}
	var ans = int(^uint(0) >> 1)
	if root.Left != nil {
		ans = min(minDepth(root.Left)+1, ans)

	}
	if root.Right！=nil {
		ans = min(minDepth(root.Right)+1, ans)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
