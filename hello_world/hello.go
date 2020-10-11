package main

import "fmt"

const frenchHelloPrefix = "Bonjour, "
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const french = "French"
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return chooseLanguage(language) + name
}

func chooseLanguage(language string) (prefix string) {
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

func main() {
	fmt.Printf(Hello("John", ""))
}
