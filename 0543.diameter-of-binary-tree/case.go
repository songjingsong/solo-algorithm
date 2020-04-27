package case0543

func diameterOfBinaryTree(root *TreeNode) int {
	ans:=1
	depth(root,&ans)
	return ans-1
}
func depth(root *TreeNode,ans *int)int{
	if root==nil{
		return 0
	}
	l:=depth(root.Left,ans)
	r:=depth(root.Right,ans)
	*ans=max(l+r+1,*ans)
	return max(l,r)+1
}
func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}