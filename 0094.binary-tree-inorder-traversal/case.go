//https://leetcode-cn.com/problems/binary-tree-inorder-traversal
package case0094

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	recursionTraversal(root, &ans)
	return ans
}

//基于递归的算法
func recursionTraversal(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	recursionTraversal(root.Left, ans)
	*ans = append(*ans, root.Val)
	recursionTraversal(root.Right, ans)
}

//基于栈的遍历
func stackTraversal(root *TreeNode) []int {
	stacks, ans := make([]*TreeNode, 0), make([]int, 0)
	cur := root
	for cur != nil || len(stacks) > 0 {
		for cur != nil {
			stacks = append(stacks, cur)
			cur = cur.Left
		}
		cur = stacks[len(stacks)-1]
		ans = append(ans, cur.Val)
		cur = cur.Right
		stacks = stacks[:len(stacks)-1]
	}
	return ans
}

//莫里斯遍历
func MorrisTraversal(root *TreeNode) []int {
	ans, cur := make([]int, 0), root
	var pre *TreeNode
	for cur != nil {
		if cur.Left == nil {
			ans = append(ans, cur.Val)
			cur = cur.Right
		} else {
			pre = cur.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = cur
			temp := cur
			cur = cur.Left
			temp.Left = nil
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
