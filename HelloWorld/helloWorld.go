package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {

	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) { // prefix will be assigned the zero-value (for string it is "")
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	//here it returns prefix if you don't specify what to return, but you can return any string if you want
	return
}

func main() {
	fmt.Println(Hello("Arman", ""))
}
