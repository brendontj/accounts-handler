package main

import (
	"cautious-octo-pancake/application"
	"cautious-octo-pancake/internal/bank"
	"cautious-octo-pancake/internal/bank/storage"
	"fmt"
)

func main() {
	fmt.Println("Starting Rest API with Mux Routers")
	app := application.Application{
		Bank:   bank.NewBank(storage.NewMemoryRepository()),
		Router: nil,
	}

	app.Start()
	app.Run(":8080")
}
