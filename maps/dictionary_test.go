package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assetStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assetStrings(t, err.Error(), want)
	})
}

func TestDictionaryAdd(t *testing.T) {
	definition := "this is just a test"
	dictionary := Dictionary{}
	dictionary.Add("test", definition)

	got, err := dictionary.Search("test")
	want := definition

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assetStrings(t, got, want)
}

func assetStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
