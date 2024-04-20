package leetcode

func ValidParentheses(s string) bool {
	stack := []rune{}
	for _, runeValue := range s {
		if runeValue == '(' || runeValue == '[' || runeValue == '{' {
			stack = append(stack, runeValue)
		} else if runeValue == ')' {
			if !checkParenthes(&stack, '(') {
				return false
			}
		} else if runeValue == ']' {
			if !checkParenthes(&stack, '[') {
				return false
			}
		} else if runeValue == '}' {
			if !checkParenthes(&stack, '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}

// 驗證 stack head 的值是否等於 parenthes
// 若相同，則回傳 true，並且將其 pop
// 否則回傳 false
func checkParenthes(stack *[]rune, parenthes rune) bool {
	if len(*stack) == 0 {
		return false
	}

	val := (*stack)[len(*stack)-1]
	if val != parenthes {
		return false
	} else {
		*stack = (*stack)[:len(*stack)-1]
		return true
	}
}
