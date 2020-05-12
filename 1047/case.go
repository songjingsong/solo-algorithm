package case

func removeDuplicates(S string) string {
	ans := make([]byte, 0)
	for i:=0;i<=len(S);i++{
		if len(ans)>0 && ans[len(ans)-1]==S[i]{
			ans=ans[:len(ans)-1]
			continue
		}
		ans=append(ans,S[i])
	}
	return string(ans)
}