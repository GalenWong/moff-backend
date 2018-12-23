package routes

import (
	"encoding/json"
	"net/http"

	"moff-backend/services"

	"github.com/gorilla/mux"
)

/*
SongList is the endpoint that returns the song list for a specific user
*/
func songList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var resp = struct {
		Message string   `json:"message"`
		List    []string `json:"list"`
	}{
		"hello",
		[]string{"thank you next"},
	}
	json.NewEncoder(w).Encode(resp)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	var newUser *UserInfo
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := services.CreateNewUser(newUser.ID, newUser.Name); err == services.ErrCreateDuplicateUser {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
AllRoutes returns all the routes in a mux.Router
which can be used for http server
*/
func AllRoutes() *mux.Router {
	var AllRoutes = mux.NewRouter()
	AllRoutes.HandleFunc("/songlist", songList).Methods("GET")
	AllRoutes.HandleFunc("/user/new", newUser).Methods("POST")

	return AllRoutes
}
