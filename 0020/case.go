package case0020

func isValid(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) > 0 && ((s[i] == '}' && stack[len(stack)-1] == '{') ||
			(s[i] == ']' && stack[len(stack)-1] == '[') ||
			(s[i] == ')' && stack[len(stack)-1] == '(')) {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return len(stack) == 0
}
