package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Arman", "")
		want := "Hello, Arman"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

// It is important that you test are clear specification of what the cade need to do.
// But there is repeated code when we check if the message is what we expect

// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy,
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper() is needed to tell the test suite that this  method is a helper.
	// by doing this when it fails the line number reported will be in our function call rather than inside our test helper.
	// You can check this out by commenting t.Helper() and make some errors in testing
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
