package main

import "fmt"

/**
 * when testing code, it is important to separate your "domain" code from the outside world
 * i.e. side effects
 * in this case, fmt.Println() is the side-effect since it prints to stdout
 * and the string we give it is our domain

 * so in order for us to easily test this, we separate these concerns
 */

// func main() {
// 	fmt.Println("Hello, world")
// }

// new code for easier testing

// create constants to capture the meaning of values
// "Hello, " by itself doesn't seem to be something worth creating a constant for
// but what happens if you want to say Hello in other languages?
// if you just have the string as is, then you've hardcoded it and it is difficult to change
// but if you change it to a constant, you can introduce other constants

// can group constants in a block like this
const (
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// greeting is a NAMED RETURN VALUE
// this basically creates a variable called greeting and assigned a "zero" value
// for int, a zero value is 0 while for strings, it is an empty string ""

// THIS FUNCTION STARTS WITH A lower case letter
// public = CAPITAL LETTER
// private = lowercase letter
func greetingPrefix(language string) (greeting string) {
	switch language {
	case spanish:
		greeting = spanishHelloPrefix
	case french:
		greeting = frenchHelloPrefix
	default:
		greeting = englishHelloPrefix
	}

	// since we have a NAMED RETURN VALUE, we can just return without the value
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
