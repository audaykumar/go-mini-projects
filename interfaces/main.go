package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// printGreetingEb(eb)
	// printGreetingSb(sb)

	// Single printGreeting for both the structs
	printGreeting(eb)
	printGreeting(sb)
}

/*
Both the printGreeting function do the same functionality.
Without interfaces would need to rewrite code
print greeting for english bot
*/
// func printGreetingEb(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

/* print greeting for spanish bot */
// func printGreetingSb(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//receiver variable if not used doen't need to be added
// can use func (eb englishBot) getGreeting string {} or the below way
func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
