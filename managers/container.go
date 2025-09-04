package managers

import (
	"log"
	"mock-api/model"
	"time"

	"go.uber.org/dig"
)

func GetDig() {
	year := 1984
	day := 1
	month := time.November
	birthDate := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	acct := model.Account{Id: 1, Name: "Ali", BirthDate: birthDate}
	container := dig.New()

	container.Provide(func() model.Account { return acct })

	// it provides a new account to the container using the injected values
	if err := container.Provide(acct); err != nil {
		log.Fatal(err)
	}

	if err := container.Invoke(SayHi); err != nil {
		log.Fatal(err)
	}
}
