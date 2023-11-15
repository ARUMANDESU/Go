package main

import "fmt"

/*
	in Go, a string is in effect a read-only slice of bytes.
*/

func main() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println("")
	// shorter way to generate presentable output
	fmt.Printf("% x\n", sample)

	// The %q (quoted) verb will escape any non-printable byte sequences in a string to the output is unambiguous.
	fmt.Printf("%q\n", sample)

	//If we are unfamiliar or confused by strange values in the string, we can use the “plus” flag to the %q verb.
	//This flag causes the output to escape not only non-printable sequences, but also any non-ASCII bytes, all while interpreting UTF-8.
	//The result is that it exposes the Unicode values of properly formatted UTF-8 that represents non-ASCII data in the string:
	fmt.Printf("%+q\n", sample)

	//This also works for slice of bites
	b := []byte{'\xbd', '\xb2', '\x3d', '\xbc', '\x20', '\xe2', '\x8c', '\x98'}
	fmt.Println(b)

	fmt.Printf("% x\n", sample)

	fmt.Printf("%q\n", sample)

	fmt.Printf("%+q\n", sample)

	for _, v := range sample {
		fmt.Printf("%q ", v)
	}

	fmt.Println("\nUTF-8 and string literals")
	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	// Go source code is always UTF-8
	// rune is the same as code point
	// rune is alias of int32

}
