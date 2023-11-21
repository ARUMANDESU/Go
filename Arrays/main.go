package main

import "fmt"

/*
	I Go, an array is a numbered sequence of elements of a specific length.
	In typical Go code, slices are much more common;
	arrays are useful in sme special scenarios.
*/

func main() {
	/*
		Here we create an array 'a' that will hold exactly 5 integers.
		The type of elements and length are both part of the array's type.
		By default, an array is zero-valued, which for integers means 0s
	*/

	var a [5]int
	fmt.Println("emp:", a) //emp: [0 0 0 0 0]

	/*
		We can set a value at an index using the ```array[index] = value``` syntax, and get a value with array[index].
	*/
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a)) // len: 5

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// Array type are one-dimensional, but you can compose type to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// Also for 2d array you can declare and initialize in one line.
	twoDTwo := [2][3]int{{1, 2, 3}, {1, 3, 5}}
	fmt.Println("2d2:", twoDTwo)
	fmt.Printf("%T\n", twoDTwo) //[2][3]int

	c := [...]string{"Penn", "Teller"} // compiler count the array elements for you
	fmt.Printf("arr: %v, type: %T\n", c, c)
}

func Sum(numbers []int) (sum int) {
	for _, num := range numbers {
		sum += num
	}
	return
}

func SumAll(numsToSum ...[]int) []int {
	sum := make([]int, len(numsToSum))

	for i, nums := range numsToSum {
		sum[i] = Sum(nums)
	}

	return sum
}
