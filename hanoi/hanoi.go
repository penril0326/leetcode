package hanoi

import "fmt"

func Hanoi(plateCount int, src, ast, des string) {
	if 1 == plateCount {
		s := fmt.Sprintf("Move plate from %s to %s.", src, des)
		fmt.Println(s)
	} else {
		Hanoi(plateCount-1, src, des, ast)
		Hanoi(1, src, ast, des)
		Hanoi(plateCount-1, ast, src, des)
	}
}
