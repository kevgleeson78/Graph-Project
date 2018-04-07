/*
*App Name: Gmit-Project
*Author:Kevin Gleeson
*Student Number: G00353180
*Version: 1.0
*Resources:
*
 */

package main

// IntoPost Capitol function name for exporting to nfa.go
func IntoPost(infix string) string {
	//Mapping special characters and giving them a priority
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
	//variable to return the new post fix string
	//Array of runes needs to be cast to string.
	posfix := []rune{}
	//variable to place each character of the input string into stack array
	s := []rune{}
	//for loop to iterate over the input string with index of each rune within the string.
	for _, r := range infix {
		switch {
		//if the array has an open bracket
		case r == '(':
			//put that rune into the array s
			s = append(s, r)
			//if the rune is a close bracket
		case r == ')':
			//for loop to run until the character ( is found
			for s[len(s)-1] != '(' {
				//append each rune to the end of the array posfix
				posfix = append(posfix, s[len(s)-1])
				//select and store in s everything apart from the last rune in the array s.
				s = s[:len(s)-1]
			}
			//remove open bracket by gettint everything in s array apart from the last character.
			s = s[:len(s)-1]
			//case for mapped runes greater than 0. 0 in runes are equal to null.
		case specials[r] > 0:
			//loop while s greater than 0 and the value of the special rune is less than the value of the character of the top of the stack.
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				//append each rune to the end of the array posfix
				posfix = append(posfix, s[len(s)-1])
				//select and store in s everything apart from the last rune in the array s.
				s = s[:len(s)-1]
			}
			//Append s from stack to s
			s = append(s, r)
		default:
			posfix = append(posfix, r)
		}
	}
	for len(s) > 0 {
		//Append anything that is left in the stack.
		posfix = append(posfix, s[len(s)-1])
		s = s[:len(s)-1]
	}

	return string(posfix)
}

/*
	//Answer ab.c*. => ('a' followed by a 'b' followed by 0 or more 'c')
	fmt.Println("infix: ", "a.b.c*")
	fmt.Println("PostFix: ", intoPost("a.b.c*"))

	//Answer abd|.* => ()
	fmt.Println("infix: ", "(a.(b|d))*")
	fmt.Println("PostFix: ", intoPost("(a.(b|d))*"))
	//Answer abd|.c*.
	fmt.Println("infix: ", "a.(b|d).c*")
	fmt.Println("PostFix: ", intoPost("a.(b|d).c*"))
	//Answer abb.+.c.
	fmt.Println("infix: ", "a.(b.b)+.c")
	fmt.Println("PostFix: ", intoPost("a.(b.b)+.c"))
*/
