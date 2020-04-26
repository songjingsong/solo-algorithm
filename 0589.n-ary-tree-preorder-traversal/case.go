package case0589

func preorder(root *Node) []int {
	ans:=make([]int,0)
	order(root,&ans)
	return ans
}
func order(root *Node,ans*[]int){
	if root==nil{
		return
	}
	*ans=append(*ans,root.Val)
	if root.Children!=nil &&len(root.Children)>0{
		for _, item := range root.Children {
			order(item,ans)
		}
	}
}
type Node struct {
	Val int
	Children []*Node
}
