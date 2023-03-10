package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/zmb3/spotify/v2"
	"net/http"
)

// profile godoc
// @Summary Get users profile information
// @Description Get Spotify profile information for the user
// @Tags profile
// @Accept  json
// @Produce  json
// @Param profile path string true "Enter user id"
// @Router /api/v1/profile/{profile} [get]
func (a *App) profile(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	vars := mux.Vars(r)
	profile := vars["profile"]

	user, err := a.Client.GetUsersPublicProfile(context.Background(), spotify.ID(profile))
	if err != nil {
		a.Log.Error(err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, user)

}
