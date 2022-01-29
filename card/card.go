package card

import "fmt"

/*
 * init function added
 */
func init() {
	fmt.Println("Simple interest package initialized")
}

// Hello returns a greeting for the named person.
func Marriage(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
