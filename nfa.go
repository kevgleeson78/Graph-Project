/*
*App Name: Gmit-Project
*Author:Kevin Gleeson
*Student Number: G00353180
*Version: 1.0
*Resources:
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
		//Concatination character
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
			nfaStack = append(nfaStack, &nfa{initial: frag2.initial, accept: frag2.accept})
			//Union Character
		case '|':
			//get the last character in the nfaStack array
			frag2 := nfaStack[len(nfaStack)-1]
			//get everything up to the last character in the nfaStack array
			nfaStack = nfaStack[:len(nfaStack)-1]
			//Get the last character from the above newly assigned nfaStack
			frag1 := nfaStack[len(nfaStack)-1]
			//get everything up to the last element in the array
			nfaStack = nfaStack[:len(nfaStack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
			//Klanee star character
		case '*':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		//All other characters
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
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

func addState(l []*state, s *state, a *state) []*state {
	l = append(l, s)
	if s != a && s.symbol == 0 {
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}

	return l
}

func pomatch(po string, s string) bool {
	ismatch := false
	//convert infix to postfix from reg.go file
	convert := IntoPost(po)
	fmt.Println(convert)
	ponfa := posRegNfa(convert)

	current := []*state{}
	next := []*state{}

	current = addState(current[:], ponfa.initial, ponfa.accept)

	for _, r := range s {
		for _, c := range current {
			if c.symbol == r {
				next = addState(next[:], c.edge1, ponfa.accept)
			}
		}
		current, next = next, []*state{}
	}

	for _, c := range current {
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
}
func main() {

	fmt.Println(pomatch("ab.c", "c"))

}
