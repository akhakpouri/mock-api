package model

import (
	"bytes"
	"fmt"
	"time"
)

type Account struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
}

func GetAge(birth time.Time) int {
	now := time.Now()
	years := now.Year() - birth.Year()
	if now.Month() < birth.Month() || (now.Month() == birth.Month() && now.Day() < birth.Day()) {
		years--
	}
	return years
}

func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "hello, %s", name)
}

func (a *Account) PrintHi(message string) string {
	return fmt.Sprintf("Hello %s - %s", a.Name, message)
}

func NewAccount(acc Account) Account {
	return Account{Id: acc.Id, Name: acc.Name, BirthDate: acc.BirthDate}
}
