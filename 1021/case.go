package case

func removeOuterParentheses(S string) string {
    left,right,index:=0,0,0
    var ans string
    for i:=0;i<len(S);i++{
        if S[i]=='('{
            left++
        }
        if S[i]==')'{
            right++
        }
        if left==right{
            ans+=S[index+1:i]
            index=i+1
        }
    }
    return ans   
}