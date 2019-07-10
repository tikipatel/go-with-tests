package maps

import (
	"errors"
)

// Dictionary is a map with string values and string keys
type Dictionary map[string]string

// ErrCouldNotFindWord is an error that describes that a word is not found
var ErrCouldNotFindWord = errors.New("could not find the word you were looking for")

var m map[string]string

// Search is a function that will search a `Dictionary` for a given key
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrCouldNotFindWord
	}
	/* If we try to add to a nil map, the program will crash.
	m[word] = "word"
	fmt.Printf("added %s to m", m[word])
	*/
	return definition, nil
}

// Add is a function that will add the value for a sepicified key
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
