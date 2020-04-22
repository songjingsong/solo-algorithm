package case0129

func sumNumbers(root *TreeNode) int {
	ans := 0
	recursion(root, &ans, 0)
	return ans
}
func recursion(root *TreeNode, ans *int, temp int) {
	if root == nil {
		return
	}
	temp = temp*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*ans += temp
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
