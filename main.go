package main

import (
	"myapp/routes"
)

func main() {
	// db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
