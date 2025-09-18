package services

import (
	"reflect"
	"testing"
)

func mockCheckWebsite(url string) bool {
	return url == "https://apple.com"
}

func TestWebCheck(t *testing.T) {
	websites := []string{
		"https://google.com",
		"https://apple.com",
		"https://stackoverflow.com",
	}

	want := map[string]bool{
		"https://google.com":        false,
		"https://apple.com":         true,
		"https://stackoverflow.com": false,
	}

	got := CheckWebsites(mockCheckWebsite, websites)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v, got %v", want, got)
	}
}
