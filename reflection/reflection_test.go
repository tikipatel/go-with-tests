package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "CHris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong nubmer of function calls, got %d, wanted %d", len(got), 1)
	}
}
