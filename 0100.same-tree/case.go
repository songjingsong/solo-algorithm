pacakage case0100

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return recursionTraversal(p,q)
}

func recursionTraversal(p *TreeNode, q *TreeNode)bool{
	if p==nil&&q==nil{
		return true
	}
	if p!=nil &&q !=nil &&p.Val==q.Val{
		return  recursionTraversal(p.Left,q.Left)&&recursionTraversal(p.Right,q.Right)
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
