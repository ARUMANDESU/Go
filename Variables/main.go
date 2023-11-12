package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	/*
		var b = 1
		var c = 2
	*/
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // zero-value for integer is 0
	fmt.Println(e)

	var i string // zero-value for string is empty space "  "
	fmt.Println("here ->", i, "<- here")

	var j bool // zero-value for boolean is false
	fmt.Println(j)

	var k float32 // also 0
	fmt.Println(k)

	f := "apple" // var f string = "apple" || var f = "apple"
	fmt.Println(f)

	g, h := 1, 5 // var g, h int = 1, 5 || var g, h = 1, 5
	fmt.Println(g, h)

}
