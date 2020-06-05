package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	music "apimusic/pkg/music"

	"github.com/gorilla/mux"
)

type api struct {
	router     http.Handler
	repository music.MusicRepository
}

// Server ...
type Server interface {
	Router() http.Handler
}

// New ...
func New(repo music.MusicRepository) Server {
	a := &api{repository: repo}

	r := mux.NewRouter()
	r.HandleFunc("/", a.basepage).Methods(http.MethodGet)
	r.HandleFunc("/musics", a.fetchMusics).Methods(http.MethodGet)
	r.HandleFunc("/musics/{ID:[a-zA-Z0-9_]+}", a.fetchMusic).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) basepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Main page!")

}

func (a *api) fetchMusics(w http.ResponseWriter, r *http.Request) {
	musics, _ := a.repository.FetchMusics(r.Context())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musics)
}

func (a *api) fetchMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	music, err := a.repository.FetchMusicByID(r.Context(), vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
		json.NewEncoder(w).Encode("Music Not found")
		return
	}

	json.NewEncoder(w).Encode(music)
}
