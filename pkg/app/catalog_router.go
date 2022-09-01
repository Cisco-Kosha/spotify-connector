package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zmb3/spotify/v2"
)

// searchCatalog godoc
// @Summary Scan artists, albums, playlists and tracks catalog for the provided search term
// @Description Search spotify catalog
// @Tags catalog
// @Accept  json
// @Produce  json
// @Param catalog path string true "Enter catalog e.g., albums, tracks, playlists, artists"
// @Param q query string true "Search term"
// @Success 200 {object} object
// @Router /api/v1/search/{catalog} [get]
func (a *App) searchCatalog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	catalog := vars["catalog"]

	catalogItems := []string{
		"albums", "tracks", "playlists", "artists",
	}
	validCatalog := false

	for _, item := range catalogItems {
		if catalog == item {
			validCatalog = true
		}
	}

	if !validCatalog {
		respondWithError(w, http.StatusBadRequest, "invalid catalog. must be one of {artists, tracks, playlists, albums}")
		return
	}

	searchTerm := r.FormValue("q")
	if searchTerm == "" {
		respondWithError(w, http.StatusBadRequest, "search parameter is not set")
		return
	}

	// search containing the term
	results, err := a.Client.Search(context.Background(), searchTerm,
		spotify.SearchTypePlaylist|spotify.SearchTypeAlbum|spotify.SearchTypeArtist|spotify.SearchTypeTrack)
	if err != nil {
		a.Log.Errorf("Error searching the catalog for the term. Error: %s", err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if catalog == "albums" {
		// handle album results
		if results.Albums != nil {
			fmt.Println("Albums:")
			for _, item := range results.Albums.Albums {
				fmt.Println("   ", item.Name)
			}
			respondWithJSON(w, http.StatusOK, results.Albums.Albums)
		}
	} else if catalog == "tracks" {
		// handle tracks results
		if results.Tracks != nil {
			fmt.Println("Tracks:")
			for _, item := range results.Tracks.Tracks {
				fmt.Println("   ", item.Name)
			}
			respondWithJSON(w, http.StatusOK, results.Tracks.Tracks)
		}
	} else if catalog == "playlists" {
		// handle playlists results
		if results.Playlists != nil {
			fmt.Println("Playlists:")
			for _, item := range results.Playlists.Playlists {
				fmt.Println("   ", item.Name)
			}
			respondWithJSON(w, http.StatusOK, results.Playlists.Playlists)
		}
	} else if catalog == "artists" {
		// handle artists results
		if results.Artists != nil {
			fmt.Println("Artists:")
			for _, item := range results.Artists.Artists {
				fmt.Println("   ", item.Name)
			}
			respondWithJSON(w, http.StatusOK, results.Artists.Artists)
		}
	}

}
