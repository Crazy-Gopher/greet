package greet

import (
	"fmt"
	"strings"
)

// Hello prints a greeting message for the specified name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func HelloUpper(name string) string {
	return strings.ToUpper(fmt.Sprintf("Hello, %s!", name))
}
