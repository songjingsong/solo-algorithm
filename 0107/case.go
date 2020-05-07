package case 


func levelOrderBottom(root *TreeNode) [][]int {
	ans:=make([][]int,0)
	dfs(root,&ans,0)
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
        ans[i], ans[j] = ans[j], ans[i]
    }
	return ans
}

func dfs(root *TreeNode,ans *[][]int,level int){
	if root==nil{
		return
	}
	if len(*ans)<=level{
		*ans=append(*ans,make([]int,0))
	}
	(*ans)[level]=append((*ans)[level],root.Val)
	dfs(root.Left,ans,level+1)
	dfs(root.Right,ans,level+1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}