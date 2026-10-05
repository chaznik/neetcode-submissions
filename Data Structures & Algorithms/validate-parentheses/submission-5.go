func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
    stack := []rune{}

	for i := 0; i < len(s); i++ {
		bracket := rune(s[i])
		if bracket == '(' || bracket == '{' || bracket == '[' {
			stack = append(stack, bracket)
		} else {
			if len(stack) != 0 {
				top := stack[len(stack) - 1]

				if top == '(' && bracket == ')' {
					stack = stack[:len(stack) - 1]
				} else if top == '{' && bracket == '}' {
					stack = stack[:len(stack) - 1]
				} else if top == '[' && bracket == ']' {
					stack = stack[:len(stack) - 1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
