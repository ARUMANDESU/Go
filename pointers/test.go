package main

import "fmt"

type A struct {
	a int
}

func main() {
	a := A{a: 3}
	fmt.Printf("address of a: %p\n", &a)
	// X(a)
	// fmt.Println(a.a)
	Y(&a)
	fmt.Printf("address of a: %p\n", &a)
	fmt.Println(a.a)
}

func X(a A) {
	a = A{a: 1}
}

func Y(a1 *A) {
	fmt.Printf("address of pointer to: %p\n", &*a1)
	fmt.Println("address of a1:", &a1)
	*a1 = A{a: 1}
	fmt.Println("\nafter assignment")
	fmt.Printf("address of pointer to: %p\n", &*a1)
	fmt.Println("address of a1:", &a1)
}
