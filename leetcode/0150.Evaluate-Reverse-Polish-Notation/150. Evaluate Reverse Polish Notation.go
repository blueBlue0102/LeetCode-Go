package leetcode

import "strconv"

func EvaluateReversePolishNotation(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			val1, val2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			result := 0
			switch token {
			case "+":
				result = val1 + val2
			case "-":
				result = val1 - val2
			case "*":
				result = val1 * val2
			case "/":
				result = val1 / val2
			}
			stack = append(stack, result)
		} else {
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[0]
}
