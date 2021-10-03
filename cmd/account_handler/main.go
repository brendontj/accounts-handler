package main

import (
	"cautious-octo-pancake/application"
	"cautious-octo-pancake/internal/account_handler"
	"cautious-octo-pancake/internal/account_handler/storage"
)

func main() {
	app := application.Application{
		AccountHandler: account_handler.NewAccountHandler(storage.NewMemoryRepository()),
		Router:         nil,
	}

	app.Start()
	app.Run(":8000")
}
