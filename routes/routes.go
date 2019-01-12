package routes

import (
	"github.com/gorilla/mux"
)

/*
AllRoutes returns all the routes in a mux.Router
which can be used for http server
*/
func AllRoutes() *mux.Router {
	var AllRoutes = mux.NewRouter()
	AllRoutes.HandleFunc("/songs", authMidware(songList)).Methods("GET")
	AllRoutes.HandleFunc("/song/{id}", authMidware(downloadSong)).Methods("GET")
	AllRoutes.HandleFunc("/song", authMidware(uploadSong)).Methods("POST")
	AllRoutes.HandleFunc("/song/{id}", authMidware(deleteSong)).Methods("DELETE")
	AllRoutes.HandleFunc("/user/new", newUser).Methods("POST")
	AllRoutes.HandleFunc("/auth/login", login).Methods("POST")
	AllRoutes.HandleFunc("/auth/validate", validate).Methods("GET")
	AllRoutes.HandleFunc("/auth/google/token", authMidware(updateAuthToken)).Methods("POST")

	return AllRoutes
}
