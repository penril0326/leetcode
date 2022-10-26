package simplifypath

import (
	"practice/data_structure/stack"
	"strings"
)

// Time complexity: O(n)
// Space complexity: O(n)
func simplifyPath(path string) string {
	directorys := strings.Split(path, "/") // T: O(n) ?
	stk := stack.NewStack()                // S: O(n)
	for _, dir := range directorys {       // T: O(n)
		if dir == ".." {
			stk.Pop()
		} else if (dir == ".") || (dir == "") {
			continue
		} else {
			stk.Push(dir)
		}
	}

	outStack := stack.NewStack() // S: O(n)
	for !stk.IsEmpty() {         // T: O(n)
		outStack.Push(stk.Pop())
	}

	out := []string{}         // S: O(n)
	for !outStack.IsEmpty() { // T: O(n)
		out = append(out, outStack.Pop().(string))
	}

	return "/" + strings.Join(out, "/") // T: O(2n) concat string is O(n), Join is O(n)(?)

	// Total time: O(6n) -> O(n)
	// Total space: O(3n) -> O(n)
}
