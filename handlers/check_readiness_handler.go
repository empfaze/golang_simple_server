package handlers

import (
	"github/empfaze/simple_server/utils"
	"net/http"
)

func CheckReadinessHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}
