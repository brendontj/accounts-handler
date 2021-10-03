package main

import (
	"cautious-octo-pancake/application"
	"cautious-octo-pancake/internal/bank"
	"cautious-octo-pancake/internal/bank/storage"
)

func main() {
	app := application.Application{
		Bank:   bank.NewBank(storage.NewMemoryRepository()),
		Router: nil,
	}

	app.Start()
	app.Run(":8080")
}
