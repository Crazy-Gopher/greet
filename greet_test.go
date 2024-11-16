package greet

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Go!"
	result := Hello("Go")
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
