package case1315

func sumEvenGrandparent(root *TreeNode) int {
    var ans int

    dfs(nil,nil,root,&ans)

    return ans
}

func dfs(grand,father,root *TreeNode,ans *int){
    if root==nil{
        return
    }
    if grand!=nil&& grand.Val%2==0{
        *ans+=root.Val
    }
    dfs(father,root,root.Left,ans)
    dfs(father,root,root.Right,ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
