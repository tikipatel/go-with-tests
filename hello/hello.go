package hello

import "fmt"

const english = "English"
const englishHelloPrefix = "Hello, "

const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "

const french = "French"
const frenchHelloPrefix = "Bonjour, "

func main() {
	fmt.Println(Hello("world", "English"))
}

// Hello is a function that takes in name and language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return
}

// UntestedFunction is an untested function
func UntestedFunction() {
	fmt.Println("This is code from an untested function.")

}
