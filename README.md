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
6: This will run the application to the terminal window.

7: Enter a Infix notation regular expression on the first prompt.

8: Enter the string you want to match.

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

