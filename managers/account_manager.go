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

func newAccount(name string) model.Account {
	id := rand.Intn(100)
	year := 1984
	day := 1
	month := time.November
	birthDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return model.Account{Id: id, Name: name, BirthDate: birthDate}

}
