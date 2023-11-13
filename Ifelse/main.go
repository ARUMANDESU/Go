package main

import "fmt"

func main() {
	// Here is a basic example  | this always outputs "7 is odd"
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Here is if statement without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators like && and || are often useful in conditions.
	if 7%2 == 0 || 8%2 == 0 {
		fmt.Println("eiter 8 or 7 are even")
	}

	/*
		a statement can precede conditionals;
		any variables declared in this statement are available in the current and all subsequent branches.
	*/
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// that means that you can not use num outside
	// it is only available inside if scope
	// fmt.Println(num)  // can't do like this

	// Note that you don't need parentheses around conditions in Go,
	// but that the curly braces are required.
	/*
		if (true != false) { // do not do like this, ok?
			fmt.Println("something like this")
		}
	*/

	/*
		There is no ternary if in Go, so you'll need to use a full if statement
	*/
}
