package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrCouldNotFindWord = errors.New("could not find the word you were looking for")

var m map[string]string

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

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
