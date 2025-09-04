package managers

import (
	"fmt"
	"mock-api/model"
)

func SayHi(acct model.Account) {
	msg := acct.PrintHi("Welcome to the mock-api universe!")
	fmt.Println(msg)
}
