package simplifypath

import (
	"practice/leetcode"
	"strings"
)

// Time complexity: O(n)
// Space complexity: O(n)
func simplifyPath(path string) string {
	directorys := strings.Split(path, "/") // T: O(n) ?
	stack := leetcode.NewStack()           // S: O(n)
	for _, dir := range directorys {       // T: O(n)
		if dir == ".." {
			stack.Pop()
		} else if (dir == ".") || (dir == "") {
			continue
		} else {
			stack.Push(dir)
		}
	}

	outStack := leetcode.NewStack() // S: O(n)
	for !stack.IsEmpty() {          // T: O(n)
		outStack.Push(stack.Pop())
	}

	out := []string{}         // S: O(n)
	for !outStack.IsEmpty() { // T: O(n)
		out = append(out, outStack.Pop().(string))
	}

	return "/" + strings.Join(out, "/") // T: O(2n) concat string is O(n), Join is O(n)(?)

	// Total time: O(6n) -> O(n)
	// Total space: O(3n) -> O(n)
}
