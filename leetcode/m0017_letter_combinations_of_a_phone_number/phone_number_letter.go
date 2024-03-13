package lettercombinationsofaphonenumber

import "strconv"

// Time: O(4^n*n), n is length of digits.
//
//	4 is the number of letters which the digit correspond to. In this case '7' and '9' correspond to 4 letters
//
// Space: O(n)
func letterCombinations(digits string) []string {
	graph := [][]byte{
		{}, // 0
		{}, // 1
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
		{'m', 'n', 'o'},
		{'p', 'q', 'r', 's'},
		{'t', 'u', 'v'},
		{'w', 'x', 'y', 'z'},
	}

	ans := []string{}
	backtracking(graph, []byte{}, digits, 0, len(digits), &ans)
	return ans
}

func backtracking(graph [][]byte, curr []byte, digits string, start, end int, ans *[]string) {
	if end == 0 {
		return
	} else if len(curr) == end {
		new := make([]byte, end)
		copy(new, curr)
		*ans = append(*ans, string(new))
		return
	}

	dig, _ := strconv.Atoi(string(digits[start]))
	for _, letter := range graph[dig] {
		curr = append(curr, letter)
		backtracking(graph, curr, digits, start+1, end, ans)
		curr = curr[:len(curr)-1]
	}
}
