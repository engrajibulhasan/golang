package main

import "fmt"

const englishHelloPrefix = "Hello "
const banglaHelloPrefix = "আসসালামুয়ালাইকুম "
const bangla = "Bangla"
const spanishHelloPrefix = "Hola "
const spanish = "Spanish"

func message(name string, language string) string {
	if name == "" {
		name = "world"
	}
	prefix := englishHelloPrefix

	switch language {
	case bangla:
		prefix = banglaHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(message("Rajibul", "Bangla"))
}
