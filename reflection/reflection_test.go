package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got '%s', want '%s'", got[0], expected)
	}

	if len(got) != 1 {
		t.Errorf("wrong nubmer of function calls, got %d, wanted %d", len(got), 1)
	}
}

func TestWalkExtended(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		}, {
			"Pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		}, {
			"Nested fields",
			struct {
				Name    string
				Profile struct {
					Age  int
					City string
				}
			}{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		}, {
			"Slices",
			[]Profile{
				{33, "London"},
				{25, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		}, {
			"Arrays",
			[2]Profile{
				{33, "London"},
				{42, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		}, {
			"Map",
			map[string]string{
				"Foo":  "Bar",
				"Bazz": "Buzz",
			},
			[]string{"Bar", "Buzz"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
