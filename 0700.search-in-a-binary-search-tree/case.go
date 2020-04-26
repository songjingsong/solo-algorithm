package case700

func searchBST(root *TreeNode, val int) *TreeNode {
	if root==nil{
		return nil
	}
	if root.Val==val{
		return root
	}
	if root.Val>val{
		return searchBST(root.Left,val)
	}
	return searchBST(root.Right,val)
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
