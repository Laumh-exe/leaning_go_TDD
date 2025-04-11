package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bounjour, "
	spanish            = "spanish"
	french             = "french"
)

func main() {
	fmt.Println(Hello("Chris", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(language) + name + "!"
}

func getPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
