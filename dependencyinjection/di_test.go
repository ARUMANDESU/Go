package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Arman")

	got := buffer.String()
	want := "Hello, Arman"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
