package main

import (
	"fmt"
)

/*
	Go's structs are typed collections of fields.
	They're useful for grouping data together to form records.
*/

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

// You can safely return a pointer to local variable as a local variable will survive the scope of the function
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main1() {
	//This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	//You can name the fields when initializing a struct.
	fmt.Println(person{
		name: "Alice",
		age:  30,
	})

	//Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	//An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	//It's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	//Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
	fmt.Println(sp.age)

	//Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	// if a struct type only used for a single value, we don't have to give it a name.
	//The value can have an anonymous struct type.
	//This technique is commonly used for table-driven tests
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)

	sp1 := s // sp1 is a copy of s
	fmt.Println(sp1.age)

	// changes made in `sp1` doesn't affect on `s`, because `sp1` is separate copy of `s`
	sp1.age = 54
	fmt.Println(sp1.age)
	fmt.Println(s.age)

	s.age = 67
	fmt.Println("sp1:", sp1.age)
	fmt.Println("sp:", sp.age)
	fmt.Println("s:", s.age)

	fmt.Println("s:", &s)
	fmt.Println("sp:", &sp)
	fmt.Println("sp1:", &sp1)

}

//Go supports embedding of structs and interfaces to express a more seamless composition of types

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container embeds a base. An embedding looks like a field without a name.
type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{num: 1},
		str:  "some name",
	}

	//we can access the base's fields directly on co, e.g. co.num.
	fmt.Printf("co=:num: %vm str: %v}\n", co.num, co.str)

	//Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)

	//Since container embeds base, the methods of base also become methods of a container.
	//Here we invoke a method that was embedded from base directly on co.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	//Embedding structs with methods may be used to bestow interface implementation onto other structs.
	//Here we see that a container now implements the describer interface because it embeds base.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
