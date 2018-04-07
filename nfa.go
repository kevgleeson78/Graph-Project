/*
*App Name: Gmit-Project
*Author:Kevin Gleeson
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
	nfaStack := []*nfa{}
	for _, r := range posFix {
		switch r {
		case '.':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]

			frag1.accept.edge1 = frag2.initial
			nfaStack = append(nfaStack, &nfa{initial: frag2.initial, accept: frag2.accept})

		case '|':
			frag2 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			frag1 := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		case '*':
			frag := nfaStack[len(nfaStack)-1]
			nfaStack = nfaStack[:len(nfaStack)-1]
			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}
			nfaStack = append(nfaStack, &nfa{initial: &initial, accept: &accept})
		}
	}
	if len(nfaStack) != 1 {
		fmt.Println("Uh Oh:", len(nfaStack), nfaStack)
	}
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
