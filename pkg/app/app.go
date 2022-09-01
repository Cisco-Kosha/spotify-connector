package app

import (
	"context"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zmb3/spotify/v2"
	"github.com/zmb3/spotify/v2/auth"

	"github.com/kosha/spotify-connector/pkg/config"
	"github.com/kosha/spotify-connector/pkg/logger"
)

type App struct {
	Router *mux.Router
	Log    logger.Logger
	Cfg    *config.Config
	Client *spotify.Client
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	return router
}

// Initialize creates the necessary scaffolding of the app
func (a *App) Initialize(log logger.Logger) {

	cfg := config.Get()

	a.Cfg = cfg
	a.Log = log

	a.Router = router()

	httpClient := a.getHttpClient()
	a.Client = spotify.New(httpClient)

	a.initializeRoutes()

}

func (a *App) getHttpClient() *http.Client {

	expiresAt, err := time.Parse(time.RFC3339, a.Cfg.GetExpiresAt())
	if err != nil {
		a.Log.Errorf("cannot parse expiresAt as time")
	}

	token := &oauth2.Token{
		AccessToken:  a.Cfg.GetAccessToken(),
		RefreshToken: a.Cfg.GetRefreshToken(),
		Expiry:       expiresAt,
	}

	return spotifyauth.New().Client(context.Background(), token)
}

// Run starts the app and serves on the specified addr
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
