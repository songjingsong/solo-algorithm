package case590

 type Node struct {
     Val int
     Children []*Node
 }


func postorder(root *Node) []int {
    ans:=make([]int,0)
    post(root,&ans)
    return ans
}
func post(root *Node,ans*[]int){
    if root==nil{
        return
    }
    if root.Children!=nil &&len(root.Children)>0{
        for _, item := range root.Children {
            post(item,ans)
        }
    }
    *ans=append(*ans,root.Val)
}