package main

import (
	"fmt"
	"math"
)

//Go supports constants of character, string, boolean, and numeric values

// const declares a constants value
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	//Constant expressions perform arithmetic with arbitrary precision(high precision without losing accuracy due to the limitations of fixed-size data types like float64.).
	const d = 3e20 / n // 3 * 10^20 / n
	fmt.Println(d)
	/*
		Numeric constants in Go don't have a type until they are given one.
		However, when used in a context that requires a type, the compiler assigns a default type based on the value.
	*/
	fmt.Printf("%T\n", d) // float64

	//numeric constant has no type until it's given one
	fmt.Println(int64(d))
	fmt.Printf("%T\n", int64(d)) //int64

	/*
		A number can be given a type by using it in a context that requires one,
		such as a variable assignment or function call.
		For example, here math.Sin expects a float64.
	*/
	fmt.Println(math.Sin(n))
	requiresUint(n) //requires uint but we can use constant number
	requiresUint(d)
}

func requiresUint(a uint) {
	// do something
}
