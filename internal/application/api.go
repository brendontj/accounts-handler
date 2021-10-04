package application

import (
	"cautious-octo-pancake/internal/account"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Api struct {
	account.Service
	*mux.Router
}

func (a *Api) Start() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *Api) Run(addr string) {
	log.Println(fmt.Sprintf("Api running on port: %v", addr))
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
