package case0559

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Children == nil {
		return 1
	}
	temp:=0
	for _, item := range root.Children {
		if item == nil {
			continue
		}
		temp=max(temp,maxDepth(item))
	}
	return temp+1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Node struct {
	Val      int
	Children []*Node
}
