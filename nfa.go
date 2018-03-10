package main

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}
type nfa struct {
	initial *state
	accept  *state
}

func posRegNfa(posfix string) {

}

func main() {
	nfa := posRegNfa("ab.c*|")
}
