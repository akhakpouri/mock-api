package managers

import (
	"log"
	"math/rand"
	"mock-api/model"
	"time"

	"go.uber.org/dig"
)

func GetDig() {

	container := dig.New()

	container.Provide(func() string { return "Ali" })

	// it provides a new account to the container using the injected values
	if err := container.Provide(newAccount); err != nil {
		log.Fatal(err)
	}

	if err := container.Invoke(SayHi); err != nil {
		log.Fatal(err)
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
