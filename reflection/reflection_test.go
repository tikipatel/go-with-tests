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
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
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
