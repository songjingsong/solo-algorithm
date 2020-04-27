package case0897

func increasingBST(root *TreeNode) *TreeNode {
	head := &TreeNode{}
	cur := head
	type dfsFunc func(root *TreeNode)
	var dfs dfsFunc
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		root.Left = nil
		cur.Right = root
		cur = root
		dfs(root.Right)
	}
	dfs(root)
	return head.Right
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
