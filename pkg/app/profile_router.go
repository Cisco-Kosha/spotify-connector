package app

import (
	"context"
	"fmt"
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

	vars := mux.Vars(r)
	profile := vars["profile"]

	user, err := a.Client.GetUsersPublicProfile(context.Background(), spotify.ID(profile))
	if err != nil {
		a.Log.Error(err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("User ID:", user.ID)
	fmt.Println("Display name:", user.DisplayName)
	fmt.Println("Spotify URI:", string(user.URI))
	fmt.Println("Endpoint:", user.Endpoint)
	fmt.Println("Followers:", user.Followers.Count)

	respondWithJSON(w, http.StatusOK, user)

}
