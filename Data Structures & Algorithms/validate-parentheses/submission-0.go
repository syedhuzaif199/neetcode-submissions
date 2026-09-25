func isValid(s string) bool {
    runes := []rune(s)
	stack := make([]rune, len(runes))
	top := 0
	for _, r := range runes {
		if r == '(' || r == '{' || r == '[' {
			stack[top] = r
			top += 1
			continue
		}
		
		if top == 0 {
			return false
		}

		top -= 1
		last := stack[top]
		
		if r == ')' && last != '(' {
			return false
		}

		if r == '}' && last != '{' {
			return false
		}

		if r == ']' && last != '[' {
			return false
		}

	}

	if top == 0 {
		return true
	}

	return false
}
