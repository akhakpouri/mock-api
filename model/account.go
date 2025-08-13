package model

import "time"

type Account struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birthDate"`
	age       int
}

func GetAge(birth time.Time) int {
	now := time.Now()
	years := now.Year() - birth.Year()
	if now.Month() < birth.Month() || (now.Month() == birth.Month() && now.Day() < birth.Day()) {
		years--
	}
	return years
}
