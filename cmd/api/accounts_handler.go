package main

import (
	"cautious-octo-pancake/internal/account"
	"cautious-octo-pancake/internal/application"
	"cautious-octo-pancake/internal/database"
)

func main() {
	app := application.Api{
		Service: account.NewAccountHandler(database.NewMemoryRepository()),
		Router:  nil,
	}

	app.Start()
	app.Run(":8000")
}
