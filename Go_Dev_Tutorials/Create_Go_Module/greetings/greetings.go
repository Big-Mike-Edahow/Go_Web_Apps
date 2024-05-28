// greetings.go

package greetings

import "fmt"

// Return a greeting for the named person.
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v! Welcome!", name)

	return message
}
