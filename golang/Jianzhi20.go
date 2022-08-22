package main

/* 自動機 */

func isNumber(s string) bool {
	status := map[int]map[rune]int{
		0: {' ': 0, 's': 1, 'd': 2, '.': 4},
		1: {'d': 2, '.': 4},
		2: {'d': 2, '.': 3, 'e': 5, ' ': 8},
		3: {'d': 3, 'e': 5, ' ': 8},
		4: {'d': 3},
		5: {'s': 6, 'd': 7},
		6: {'d': 7},
		7: {'d': 7, ' ': 8},
		8: {' ': 8},
	}

	cur := 0
	for _, each := range s {
		var next rune
		switch each {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			next = 'd'
		case '+', '-':
			next = 's'
		case 'e', 'E':
			next = 'e'
		case '.', ' ':
			next = each
		default:
			next = '?'
		}
		if nextstat, ok := status[cur][next]; ok {
			cur = nextstat
		} else {
			return false
		}
	}
	if cur == 2 || cur == 3 || cur == 7 || cur == 8 {
		return true
	}
	return false
}

// func main() {
// 	s := "1 "
// 	isNumber(s)
// }
