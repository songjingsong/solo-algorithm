package case0054

func kthLargest(root *TreeNode, k int) int {
	return dfs(root,&k)
}
func dfs(root *TreeNode,k *int)int  {
	if root==nil{
		return 0
	}
	r:=dfs(root.Right,k)
	if *k==0{
		return r
	}
	*k-=1
	if *k==0{
		return root.Val
	}
	return dfs(root.Left,k)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
