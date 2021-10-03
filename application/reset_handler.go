package application

import (
	"net/http"
)

func (a *Application) ResetHandler(w http.ResponseWriter, r *http.Request){
	a.AccountHandler.Reset()
	respondWithTextValue(w, http.StatusOK, "OK")
}