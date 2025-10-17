package managers

import (
	"testing"
)

func TestWalker(t *testing.T) {
	expect := "Ali"
	got := []string{}

	x := struct {
		Name string
	}{expect}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of array elements. got %d, but wanted %d", len(got), 1)
	}
}
