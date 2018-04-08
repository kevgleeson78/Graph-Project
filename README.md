# Graph-Project Regular Expression Engine. 

This is a repository with an example of a Regular expression engine built using a shunting yard algorithm along allong with an nfa data structure.
This program was built using the [Go](https://golang.org/) programming language.

Author: [Kevin Gleeson](https://github.com/kevgleeson78)

Third year student at:[GMIT](http://gmit.ie) Galway

## Cloning, compiling and running the application.

1: Download [git](https://git-scm.com/downloads) to your machine if not already installed.

1.1: Download [go](https://golang.org/dl/) if not already installed.

2: Open git bash and cd to the folder you wish to hold the repository.
Alternatively you can right click on the folder and select git bash here.
This will open the git command prompt in the folder selected.
 
 3: To clone the repository type the following command in the terminal making sure you are in the folder needed for the repository.
```bash
>git clone https://github.com/kevgleeson78/Graph-Project
```
4: To compile the application cd to the folder and type the following 
```bash
> go build 
```
This will compile and create an executable file from the .go files and give it the name of the folder.

5: To run the application ensure you cd to folder the application is held.
Type the following command
```bash
>./Graph-Project
```
6: This will run the server on port 8080. 

7: navigate to http://localhost:8080 in your web browser.

7: Enter a Infix notation regular expression on the first text field.

8: Enter the string you want to match in the second text field and press enter.

## The purpose of this applictaion
The purpose of this program is to take an infix regular expression and transform it to postfix notation.
From there using Thompsons construction the regular expression is split up to accept and initial states one characte at a time.
The characters '.' to concat, '|' to or , '*' for Klanee(0 or more) and '+' for one or more are used.

## Shunting yard algorithm.
The shunting yard algorithm is used to transform the infix regular expression to postfix.
```GO
//Mapping special characters and giving them a priority
	specials := map[rune]int{'*': 10, '+': 9, '.': 8, '|': 7}
 ```
A map of runes with priorities attached are used for shifting over the characters by precedence.

```GO
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
```
## Thompsons NFA

Thompsons construction is used to create a nfa of initial and accept states for the postfix string.

```GO
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
```

Structs are used to keep track of initail and accept states.

```GO
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
```

The pomatch function is used to take in a postfix regular expression and a string.
this returns true if the string matches the postfix regular expression.

```GO
func pomatch(po string, s string) bool {
	//Variable for a matching regular expression.
	ismatch := false
	//convert infix to postfix from reg.go file
	convert := IntoPost(po)
	fmt.Println(convert)
	//Pass the param po from pomatch to the  posRegNfa function and Store the result into the variable ponfa.
	ponfa := posRegNfa(convert)
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
```

Finally from the main method user input is used to enter the infix regular expression and the string to compare against it.

```GO
func main() {
	//fmt.Println(pomatch("a.(b.b)+.c", "abbc"))
	//fmt.Println(pomatch("a.b.c*", "abccccc"))
	//fmt.Println(pomatch("(a.(b|d))*", "adadad"))
	//fmt.Println(pomatch("a.(b|d).c*", "ad"))
	fmt.Print("Enter Infix String: ")
	var infix string
	fmt.Scanln(&infix)
	fmt.Print("Enter String to match: ")
	var match string
	fmt.Scanln(&match)

	fmt.Println(pomatch(infix, match))

}
```
