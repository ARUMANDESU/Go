package main

import "fmt"

// A type assertion provides access to an interface value's underlying concrete value.
func main() {

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string) // if ok is false then s will be zero-value of type T . `i.(T)`
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*f1 := i.(float64) //panic
	fmt.Println(f1)*/

}
