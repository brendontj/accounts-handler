package application

import "net/http"

func (a *Api) initializeRoutes() {
	a.Router.HandleFunc("/balance", a.BalanceHandler).Methods(http.MethodGet)
	a.Router.HandleFunc("/event", a.EventHandler).Methods(http.MethodPost)
	a.Router.HandleFunc("/reset", a.ResetHandler).Methods(http.MethodPost)
}