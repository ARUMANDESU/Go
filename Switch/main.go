package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Switch statements express conditionals across many branches.
	*/

	// Here's a basic switch.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	/*
		You can use commas to separate multiple expressions in the same case statement.
		We use the optional default case in this example as well.
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/*
		Switch without an expression is an alternate way to express if/else logic.
		Here we also show how the case expressions can be non-constants.
	*/
	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	/*
		A type switch compares types instead of values. You can use this to discover the type of an interface value.
		In this example, the variable 't' will have the type corresponding to ins clause.
	*/
	whatAmI := func(i interface{}) { // variable of function type, it is declared and assigned an anonymous function(a function without a name)
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case func(interface{}):
			fmt.Println("lol, did you pass this function to itself as an argument")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(whatAmI)
}
