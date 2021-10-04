package application

import (
	"net/http"
)

func (a *Api) ResetHandler(w http.ResponseWriter, _ *http.Request){
	a.Service.Reset()
	respondWithTextValue(w, http.StatusOK, "OK")
}