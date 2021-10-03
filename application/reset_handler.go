package application

import (
	"net/http"
)

func (a *Application) ResetHandler(w http.ResponseWriter, r *http.Request){
	a.Bank.Reset()
	respondWithTextValue(w, http.StatusOK, "OK")
}