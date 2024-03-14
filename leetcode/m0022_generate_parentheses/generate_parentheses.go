package generateparentheses

// Time: O(2^2n/n*sqrt(n)), it's Catalan number.
// Space: O(n)
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	backtracking([]byte{}, 0, 0, n, &ans)
	return ans
}

func backtracking(curr []byte, left, right, n int, ans *[]string) {
	if len(curr) == n*2 {
		*ans = append(*ans, string(curr))
		return
	}

	var neighbors []byte
	if left == right {
		neighbors = []byte{'('}
	} else if left == n {
		neighbors = []byte{')'}
	} else {
		neighbors = []byte{'(', ')'}
	}

	for _, parenthesis := range neighbors {
		curr = append(curr, parenthesis)
		if parenthesis == '(' {
			backtracking(curr, left+1, right, n, ans)
		} else {
			backtracking(curr, left, right+1, n, ans)
		}
		curr = curr[:len(curr)-1]
	}
}
