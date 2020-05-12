package case0682

import (
    "strconv"
)

func calPoints(ops []string) int {
    stack,ans:=make([]int,0),0
    for i:=0;i<len(ops);i++{
        temp:=0
        if ops[i]=="D"{
            temp=stack[len(stack)-1]*2
            stack=append(stack,temp)
            ans+=temp
        }else if ops[i]=="C"{
            temp-=stack[len(stack)-1]
            stack=stack[:len(stack)-1]
            ans+=temp
        } else if ops[i]=="+"{
            temp:=stack[len(stack)-1]+stack[len(stack)-2]
            stack=append(stack,temp)
            ans+=temp
        }else{
            temp,_=strconv.Atoi(ops[i])
            stack=append(stack,temp)
            ans+=temp
        }

    }
    return ans
}