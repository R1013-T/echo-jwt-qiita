package main

import (
	"echo-jwt/database"
	"echo-jwt/routes"
)

func main() {
	database.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
