package application

import (
	"cautious-octo-pancake/internal/account_handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Application struct {
	account_handler.AccountHandler
	*mux.Router
}

func (a *Application) Start() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *Application) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
