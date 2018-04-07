/*
*App Name: Gmit-Project
*Author:Kevin Gleeson
*Student Number: G00353180
*Version: 1.0
*Resources:
*https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
*https://web.microsoftstream.com/video/96e6f4cc-b390-4531-ba7f-84ad6ab01f47
*https://web.microsoftstream.com/video/d08f6a02-23ec-4fa1-a781-585f1fd8c69e
*https://web.microsoftstream.com/video/946a7826-e536-4295-b050-857975162e6c
*https://web.microsoftstream.com/video/68a288f5-4688-4b3a-980e-1fcd5dd2a53b
*https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d
*https://regex101.com/r/xvnqoo/1/
*https://en.wikipedia.org/wiki/Shunting-yard_algorithm
*https://en.wikipedia.org/wiki/Thompson%27s_construction
*http://codeidol.com/community/perl/know-the-precedence-of-regular-expression-operator/14215/
*
*
 */

package main

// IntoPost Capitol function name for exporting to nfa.go
func IntoPost(infix string) string {
	//Mapping special characters and giving them a priority
	specials := map[rune]int{'*': 10, '+': 9, '.': 8, '|': 7}
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

	//Answer abd|.* => ('a' followed by  'b' or 'a' followed by  'b' 0 or more times)
	fmt.Println("infix: ", "(a.(b|d))*")
	fmt.Println("PostFix: ", intoPost("(a.(b|d))*"))

	//Answer abd|.c*. => ('a' followed by 'b' or 'a' followed by 'd' followed by 'c' 0 or more times)
	fmt.Println("infix: ", "a.(b|d).c*")
	fmt.Println("PostFix: ", intoPost("a.(b|d).c*"))

	//Answer abb.+.c. => ('a' followed by one or more 'bb' followed by 'c' )
	fmt.Println("infix: ", "a.(b.b)+.c")
	fmt.Println("PostFix: ", intoPost("a.(b.b)+.c"))
*/
