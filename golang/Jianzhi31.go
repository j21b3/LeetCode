package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, 0)
	i, j := 0, 0
	for i < len(pushed) {
		if len(stack) == 0 && pushed[i] == popped[j] {
			i++
			j++
		} else if len(stack) != 0 && popped[j] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			j++
		} else {
			stack = append(stack, pushed[i])
			i++
		}
	}
	for len(stack) > 0 && j < len(popped) {
		if stack[len(stack)-1] == popped[j] {
			j++
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return true
}
