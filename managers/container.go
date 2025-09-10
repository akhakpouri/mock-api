package managers

import (
	"log"

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
