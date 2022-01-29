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

// Engagement commited..
func Engagement(boy string, girl string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome to the Engagement boy , Are you happy to marry me %v", boy, girl)
	return message
}
