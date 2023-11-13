package main

import "fmt"

/*
for is go's only looping construct.
Here are some basic types of for loops.
*/

func main() {
	// with a single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println(" ")

	// a classic initial/condition/after 'for' loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	fmt.Println(" ")

	/*
		'for' without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	*/
	for {
		fmt.Println("loop")
		break
	}

	// We can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println(" ")

	returnOnlyEvenNumbers(3, 23)

}

func returnOnlyEvenNumbers(from int, until int) {
	var n int
	if from <= 0 {
		n = 2
	} else {
		n = from
	}

	if n%2 == 1 {
		n++
	}

	for ; n <= until; n += 2 {
		fmt.Println(n)
	}
}
