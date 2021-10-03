package application

import (
	"cautious-octo-pancake/internal/bank"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Application struct {
	bank.Bank
	*mux.Router
}

func (a *Application) Start() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *Application) Run() {
	log.Fatal(http.ListenAndServe(":8080", a.Router))
}
