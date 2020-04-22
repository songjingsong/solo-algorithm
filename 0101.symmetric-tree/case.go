package case 0101

func isSymmetric(root *TreeNode) bool {
	if root==nil{
		return true
	}
	return  symmetric(root.Left,root.Right)
}

func symmetric(a,b *TreeNode)bool{
    if a==nil && b==nil{
        return true
    }
    if a==nil ||b==nil ||a.Val!=b.Val{
        return false
    }
    return symmetric(a.Left,b.Right)&& symmetric(a.Right,b.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
