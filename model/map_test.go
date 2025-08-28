package model

import (
	"testing"
)

var colors = Dictionary{
	"black": "#000",
	"white": "#fff",
}

func printMapError(got string, want string, t testing.TB) {
	t.Helper()

	if got != want {
		t.Errorf("failed. wanted %q, but got %q", want, got)
	} else {
		t.Logf("success. wanted %q, and got %q", want, got)
	}
}

func TestMap(t *testing.T) {
	t.Run("Known color", func(t *testing.T) {
		got, _ := colors.Search("white")
		want := "#fff"
		printMapError(got, want, t)
	})
	t.Run("Unknown color", func(t *testing.T) {
		got, err := colors.Search("purple")
		want := "#000"
		if err == nil {
			t.Fatal("Excepted an error, but didn't get one")
		}
		printMapError(got, want, t)
	})
	t.Run("Add color", func(t *testing.T) {
		key, value := "gray", "#999"
		colors.Add(key, value)

		want := value
		got, err := colors.Search(key)

		if err != nil {
			t.Fatalf("color %q wasn't added correctly.", key)
		}
		printMapError(got, want, t)
	})
}
