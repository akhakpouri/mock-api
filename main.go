package main

import "mock-api/app"

func main() {
	var a app.App
	a.CreateRoutes()
	a.Run()
}
