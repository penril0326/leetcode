package issubsequence

// Time complexity: O(n), which n is length of t
// Space complexity: O(1)
func isSubsequence(s string, t string) bool {
	is, it := 0, 0
	for is < len(s) && it < len(t) {
		if s[is] == t[it] {
			is++
			it++
		} else {
			it++
		}
	}

	return is == len(s)
}
