package case0921

func minAddToMakeValid(S string) int {
	left, right := 0, 0
	for _, item := range S {
		if item == '(' {
			left++
			continue
		}
		if left > 0 {
			left--
		} else {
			right++
		}
	}
	return left + right
}
