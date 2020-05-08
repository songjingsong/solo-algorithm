package case0230

func kthSmallest(root *TreeNode, k int) int {
    var ans int

    dfs(root,&k,&ans)
    return ans
}

func dfs(root *TreeNode,k *int,ans *int){
    if root==nil ||*k==0{
        return
    }
    dfs(root.Left,k,ans)
    *k--
    if *k==0{
        *ans=root.Val
    }
    dfs(root.Right,k,ans)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}