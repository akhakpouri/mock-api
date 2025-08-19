package model

import (
	"testing"
	"time"
)

func TestGetAge(t *testing.T) {
	layout := "01/02/2006"
	birthDate, err := time.Parse(layout, "11/01/1984")
	if err != nil {
		t.Errorf("error occured. %v", err)
	}

	age := GetAge(birthDate)
	t.Logf("account is %v years old", age)
}
