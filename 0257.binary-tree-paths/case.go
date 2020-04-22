package case0257

func binaryTreePaths(root *TreeNode) []string {
	ans:=make([]string,0)
	recursion(root,&ans,"")
	return ans 
}

func sumNumbers(root *TreeNode) int {
	ans := 0
	recursion(root, &ans, 0)
	return ans
}
func recursion(root *TreeNode, ans *[]string, temp string) {
	if root == nil {
		return
	}
	if temp==""{
		temp+=fmt.Sprintf("%v",root.Val)
	}else{
		temp+=fmt.Sprintf("->%v",root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*ans=append(*ans,temp)
		return
	}
	recursion(root.Left, ans, temp)
	recursion(root.Right, ans, temp)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
