package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"slices"
)

/*
	Slices are an important data type in Go, giving a more powerful interface to sequences than arrays.

*/

func main() {
	/*
		Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
		An uninitialized slice equals to nil and has length 0.
	*/
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // [] true true
	/*
		The zero value of a slice is nil.
		The `len` and `cap` functions will both return 0 for a nil slice.
	*/
	/*
		To create an empty slice with non-zero length, use the builtin `make`.
		Here we make a slice of strings of length 3 (initially zero-valued).
		By default a new slice’s capacity is equal to its length;
		if we know the slice is going to grow ahead of time,
		it’s possible to pass a capacity explicitly as an additional parameter to `make`.

		func make([]T, len, cap) []T
	*/
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) //emp: [  ] len: 3 cap: 3

	s = make([]string, 3, 5)                               //length: 3, capacity: 5
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) //emp: [  ] len: 3 cap: 5

	/*
		The `make` function takes a type, a length, and an optional capacity.
		When called, make allocates an array and returns a slice that refers to that array.
	*/

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s, "len:", len(s), "cap:", cap(s))
	fmt.Println("get:", s[2])

	// `len` returns the length of the slice as expected.
	fmt.Println("len:", len(s))

	/*
		In addition to these basic operations, slices support several more that make them richer than arrays.
		One is the builtin `append`, which returns a slice containing one or more new values.
		Note that we need to accept a return value from `append` as we may get a new slice value.
	*/

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s, "len:", len(s), "cap:", cap(s)) //apd: [a b c d e f] len: 6 cap: 10

	/*
		Slices can also be copy’d.
		Here we create an empty slice `c` of the same length as `s` and copy into `c` from `s`.
	*/
	c := make([]string, len(s)) // by default capacity is equal to length, that's why len and cap are 6
	copy(c, s)
	fmt.Println("cpy:", c, "len:", len(c), "cap:", cap(c)) //cpy: [a b c d e f] len: 6 cap: 6

	/*
		Slices support a “slice” operator with the syntax slice[low:high].
		For example, this gets a slice of the elements s[2], s[3], and s[4].
	*/
	l := s[2:5]
	fmt.Println("sl1:", l, "len:", len(l), "cap:", cap(l)) //sl1: [c d e] len: 3 cap:8

	//This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl2:", l, "len:", len(l), "cap:", cap(l)) //sl2: [a b c d e] len: 5 cap: 10

	//And this slices up from (and including) s[2].
	l = s[2:]
	fmt.Println("sl3:", l, "len:", len(l), "cap:", cap(l)) //sl3: [c d e f] len: 4 cap: 8

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t, "len:", len(t), "cap:", cap(t)) //dcl: [g h i] len: 3 cap: 3

	// The slices package contains a number of useful utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	/*
		Slices can be composed into multi-dimensional data structures.
		The length of the inner slices can vary, unlike with multi-dimensional arrays.
	*/
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoDTwo := [][]int{{1, 2, 3, 4}, {1, 5, 3}}
	fmt.Println("2d2: ", twoDTwo)
	fmt.Printf("%T\n", twoDTwo) //[][]int

	/*
		A slice can also be formed by "slicing" an existing slice or array.
		Slicing is done by specifying a half-open range with two indices separated by a colon.
		For example, the expression b[1:4] creates a slice including elements 1 through 3 of b (the indices of the resulting slice will be 0 through 2)
	*/

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b1 := b[1:4]
	fmt.Printf("b1: %q, len: %d, cap: %d\n", b1, len(b1), cap(b1))

	//The start and end indices of a slice expression are optional; they default to zero and the slice’s length respectively:
	b1 = b[:2]
	fmt.Printf("b1: %q, len: %d, cap: %d\n", b1, len(b1), cap(b1))
	b1 = b[2:]
	fmt.Printf("b1: %q, len: %d, cap: %d\n", b1, len(b1), cap(b1))
	b1 = b[:]
	fmt.Printf("b1: %q, len: %d, cap: %d\n", b1, len(b1), cap(b1))

	//This is also the syntax to create a slice given an array
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	sl := x[:] // a slice referencing the storage of x
	fmt.Printf("sl: %v, len: %d, cap: %d\n", sl, len(sl), cap(sl))

	/*
		Slicing does not copy the slice’s data. It creates a new slice value that points to the original array.
		This makes slice operations as efficient as manipulating array indices.
		Therefore, modifying the elements (not the slice itself) of a re-slice modifies the elements of the original slice:
	*/
	d := []byte{'r', 'o', 'a', 'd'}
	fmt.Printf("d: %q, len: %d, cap: %d\n", d, len(d), cap(d))
	e := d[2:]
	fmt.Printf("e: %q, len: %d, cap: %d\n", e, len(e), cap(e))
	e[1] = 'm'
	fmt.Printf("e: %q, len: %d, cap: %d\n", e, len(e), cap(e))
	fmt.Printf("d: %q, len: %d, cap: %d\n", d, len(d), cap(d))

	/*
		A slice cannot be grown beyond its capacity.
		Attempting to do so will cause a runtime panic, just as when indexing outside the bounds of a slice or array.
		Similarly, slices cannot be re-sliced below zero to access earlier elements in the array.
	*/

	// To append one slice to another, use `...` to expand the second argument to a list of arguments.
	a := []string{"arman", "someone1"}
	f := []string{"someone2", "someone2", "someone3"}
	a = append(a, f...)
	fmt.Printf("a: %v, len: %d, cap: %d\n", a, len(a), cap(a))

}

/*
As mentioned earlier, re-slicing a slice doesn’t make a copy of the underlying array.
The full array will be kept in memory until it is no longer referenced.
Occasionally this can cause the program to hold all the data in memory when only a small piece of it is needed.
For example, this FindDigits function loads a file into memory and searches it for the first group of consecutive numeric digits, returning them as a new slice.
*/
var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return digitRegexp.Find(b)
}

/*
	This code behaves as advertised, but the returned []byte points into an array containing the entire file.
	Since the slice references the original array, as long as the slice is kept around the garbage collector can’t release the array;
	the few useful bytes of the file keep the entire contents in memory.
*/

// To fix this problem one can copy the interesting data to a new slice before returning it:
func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))
	copy(c, b)
	return c
}

// A more concise version of this function could be constructed by using append. This is left as an exercise for the reader
func CopyDigitsMy(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	return append(b)
}
