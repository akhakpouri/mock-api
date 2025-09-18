package managers

import (
	"fmt"
	"math/rand"
	"mock-api/model"
	"time"
)

func SayHi(acct model.Account) {
	msg := acct.PrintHi("Welcome to the mock-api universe!")
	fmt.Println(msg)
}

func AlwaysSayHi() {
	names := []string{
		"Ali",
		"Bob",
		"Jon",
		"Ben",
		"Nell",
		"Alec",
	}
	channel := make(chan model.Account)

	for name := range names {
		go func() {
			acct := model.Account{Name: names[name]}
			channel <- acct
		}()
	}

	for range names {
		r := <-channel
		SayHi(r)
	}
}

func newAccount(name string) model.Account {
	id := rand.Intn(100)
	year := 1984
	day := 1
	month := time.November
	birthDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return model.Account{Id: id, Name: name, BirthDate: birthDate}

}
