package main

import (
	"fmt"
)

func intoPost(infix string) string {
	posfix := ""

	return posfix
}
func main() {
	//Answer ab.c*
	fmt.Println("infix: ", "a.b.c*")
	fmt.Println("PostFix: ", intoPost("a.b.c*"))
	//Answer abd|.*
	fmt.Println("infix: ", "(a.(b|d))*")
	fmt.Println("PostFix: ", intoPost("(a.(b|d))*"))
	//Answer abd|.c
	fmt.Println("infix: ", "a.(b|d).c")
	fmt.Println("PostFix: ", intoPost("a.(b|d).c"))
	//Answer abb.+.c.
	fmt.Println("infix: ", "a.(b.b)+.c")
	fmt.Println("PostFix: ", intoPost("a.(b.b)+.c"))

}
