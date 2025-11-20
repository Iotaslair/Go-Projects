package main

import "fmt"

const (
	japanese = "Japanese"
	french   = "French"

	englishHelloPrefix  = "Hello "
	japaneseHelloPrefix = "Konnichiwa "
	frenchHelloPrefix   = "Bonjour "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case japanese:
		prefix = japaneseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("william", ""))
}
