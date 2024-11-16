package greet

import "fmt"

// Hello prints a greeting message for the specified name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
