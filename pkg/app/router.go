package app

import (
	"context"
	"encoding/json"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
	"net/http"
	"time"

	"github.com/kosha/spotify-connector/pkg/models"
)

// listConnectorSpecification godoc
// @Summary Get details of the connector specification
// @Description Get all environment variables that need to be supplied for the spotify connector to work
// @Tags specification
// @Accept  json
// @Produce  json
// @Success 200 {object} object
// @Router /api/v1/specification/list [get]
func (a *App) listConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	respondWithJSON(w, http.StatusOK, map[string]string{
		"ACCESS_TOKEN":  "OAuth2 Access Token",
		"REFRESH_TOKEN": "OAuth2 Refresh Token",
		"EXPIRES_AT":    "OAuth2 Token Expiry",
	})
}

// testConnectorSpecification godoc
// @Summary Test if oauth2 access token works against the specification
// @Description Check if the oauth2 credentials are valid
// @Tags specification
// @Accept  json
// @Produce  json
// @Param text body models.Specification false "Enter access token, refresh token and token expiry"
// @Success 200 {object} boolean
// @Router /api/v1/specification/test [post]
func (a *App) testConnectorSpecification(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")

	if (*r).Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	var s models.Specification
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&s); err != nil {
		a.Log.Errorf("Error parsing json payload", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	expiresAt, err := time.Parse(time.RFC3339, s.ExpiresAt)
	if err != nil {
		a.Log.Errorf("cannot parse expiresAt as time")
	}

	token := &oauth2.Token{
		AccessToken:  s.AccessToken,
		RefreshToken: s.RefreshToken,
		Expiry:       expiresAt,
	}

	t, err := spotify.New(spotifyauth.New().Client(context.Background(), token)).Token()

	if err != nil || t == nil {
		respondWithJSON(w, http.StatusOK, "false")
	} else {
		respondWithJSON(w, http.StatusOK, "true")
	}
}
