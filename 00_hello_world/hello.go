package main

import "fmt"

const (
	french = "fr"
	spanish = "es"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, country string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(country) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
		default:
			prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
