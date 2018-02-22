package main

import (
	"fmt"
)

func intoPost(infix string) string {
	//Mapping special characters and hiving them a priority
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	//variable to return the new post fix string
	posfix := []rune{}
	//variable to place each character of the input string into array
	s := []rune{}

	return string(posfix)
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
