package di

import "testing"
import "bytes"

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Samuel")

	got := buffer.String()
	want := "Hello, Samuel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
