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
 */

package main

import (
	"fmt"
)

//Struct that holds rune characters, and two pointers
type state struct {
	//rune used to represent arrows pointing to either edge1 or edge2 or both states
	symbol rune
	edge1  *state
	edge2  *state
}

//Struct to hold the fragment of each state and keep track of which state is in the initial or accepted state.
type nfa struct {
	initial *state
	accept  *state
}

//Function to take in a postfix string as a param and return a pointer to nfa struct
func posRegNfa(posFix string) *nfa {
	//Array Stack to hold empty nfa pointers
	nfaStack := []*nfa{}
	//loop over each rune from posFix String input
	for _, r := range posFix {
		//switch statement for specila characters in the string
		switch r {
		//Case for Concatination character
		case '.':
			//get the last character in the nfaStack array
			frag2 := nfaStack[len(nfaStack)-1]
			//get everything up to the last character in the nfaStack array
			nfaStack = nfaStack[:len(nfaStack)-1]
			//Get the last character from the above newly assigned nfaStack
			frag1 := nfaStack[len(nfaStack)-1]
			//get everything up to the last element in the array
			nfaStack = nfaStack[:len(nfaStack)-1]
			//Joining the two fragments together
			//frag1 first edge points to frag2 initial state.
			frag1.accept.edge1 = frag2.initial
			//append a new pointer to nfaStack of frag1 initial state and frag2 accept state.
			nfaStack = append(nfaStack, &nfa{initial: frag1.initial, accept: frag2.accept})

		//Case for Union Character
		case '|':
			//get the last character in the nfaStack array
			frag2 := nfaStack[len(nfaStack)-1]
			//get everything up to the last character in the nfaStack array
			nfaStack = nfaStack[:len(nfaStack)-1]
			//Get the last character from the above newly assigned nfaStack
			frag1 := nfaStack[len(nfaStack)-1]
			//get everything up to the last element in the array
			nfaStack = nfaStack[:len(nfaStack)-1]
			//New state edge1 points to Frag1.initial and edge2 points to frag2.initial
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			//New accept state
			accept := state{}
			//frag1 accept and edge point to the new accept state.
			frag1.accept.edge1 = &accept
			//frag2 accept and edge point to the new accept state.
			frag2.accept.edge1 = &accept
			//Append pointer to nfaStack with initial and accept states above as pointers.
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

		//Case for Klanee star character
		case '*':
			//get the last character in the nfaStack array
			frag := nfaStack[len(nfaStack)-1]

			//get everything up to the last character in the nfaStack array
			nfaStack = nfaStack[:len(nfaStack)-1]
			//New accept state
			accept := state{}
			//New state edge1 points to Frag.initial and edge2 points to the new accept state.
			initial := state{edge1: frag.initial, edge2: &accept}
			//Join frag accept edge1 to frag initial state.
			frag.accept.edge1 = frag.initial
			//Join frag accept edge2 to new accept state.
			frag.accept.edge2 = &accept
			//Append pointer to nfaStack with initial and accept states above as pointers.
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})

			//Case One or more
		case '+':
			//get the last character in the nfaStack array
			frag := nfaStack[len(nfaStack)-1]
			//get everything up to the last character in the nfaStack array
			nfaStack = nfaStack[:len(nfaStack)-1]
			//New accept state
			accept := state{}
			//New state edge1 points to Frag.initial and edge2 points to the new accept state.
			initial := state{edge1: frag.initial, edge2: &accept}
			//Join frag accept edge1 to frag initial state.
			frag.accept.edge1 = &initial

			//Append to nfaStack with initial and accept state above as a pointer.
			nfaStack = append(nfaStack, &nfa{initial: frag.initial, accept: &accept})
		//All other characters
		default:
			//New accept state
			accept := state{}
			//New initial state with symbol from state struct and edge1 pointing to new accept state
			initial := state{symbol: r, edge1: &accept}
			//Append pointer to nfaStack with initial and accept states above as pointers.
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}
	//Error check that there is only one element left on the stack
	if len(nfaStack) != 1 {
		fmt.Println("Uh Oh:", len(nfaStack), nfaStack)
	}
	//Return the last nfa on the stack
	return nfaStack[0]
}

//Recursive function to add states to the nfa.
func addState(l []*state, s *state, a *state) []*state {
	//Append the state that has been passed in.
	l = append(l, s)
	//If the rune is equal to 0 there is an e arrow comming from that state.
	//And s not equal to a.
	if s != a && s.symbol == 0 {
		//Keep passing l, s.edge1 and a until s.symbol !=0.
		l = addState(l, s.edge1, a)
		//Condition to check for non null value
		if s.edge2 != nil {
			//Keep Keep passing l, s.edge2 and a until s.edge2 becomes null.
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

//Function To take in a postfix regular expression and a string.
//This returns true if the string matches the postfix regular expression.
func pomatch(po string, s string) bool {
	//Variable for a matching regular expression.
	ismatch := false
	//convert infix to postfix from reg.go file
	//convert := IntoPost(po)
	//fmt.Println(convert)
	//Pass the param po from pomatch to the  posRegNfa function and Store the result into the variable ponfa.
	ponfa := posRegNfa(po)
	//New state to keep track of the current state.
	current := []*state{}
	//New state takes a pointer from current if an arrow is pointing from current
	next := []*state{}
	//Assign current to a function that has current, ponfa accept and initial states.
	current = addState(current[:], ponfa.initial, ponfa.accept)
	//loop over each rune in String s
	for _, r := range s {
		//For each rune in Strings loop over the current array
		for _, c := range current {
			//Condition to check if the rune in current is equal to the rune in the input String s.
			if c.symbol == r {
				//Assign next to addState function that has next, edge1 for current and ponfa accept.
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		//Clear the next array and copy the current array to the new empty next array
		current, next = next, []*state{}
	}
	// Loop to go over the current array
	for _, c := range current {
		//Condition if c in loop == ponfa accept state
		if c == ponfa.accept {
			//Assign ismatch to true.
			ismatch = true
			break
		}
	}
	//Return true or false depending on whether the regular expression matches the string.
	return ismatch
}
func main() {

	//fmt.Println(pomatch("a.b.c*", "abccccc"))
	//fmt.Println(pomatch("(a.(b|d))*", "adadad"))
	//fmt.Println(pomatch("a.(b|d).c*", "ad"))
	fmt.Println(pomatch("abb.+.c.", "abbbbbbbbc"))
}
