package model

import "testing"

var colors = Dictionary{
	"black": "#000",
	"white": "#fff",
}

func TestMap(t *testing.T) {
	got := colors.Search("black")
	want := "#000"

	if got != want {
		t.Errorf("got %q, but really wanted %q", got, want)
	}
}
