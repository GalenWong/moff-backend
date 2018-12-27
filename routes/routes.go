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
	AllRoutes.HandleFunc("/user/new", newUser).Methods("POST")
	AllRoutes.HandleFunc("/auth/login", login).Methods("POST")
	AllRoutes.HandleFunc("/auth/validate", validate).Methods("GET")

	return AllRoutes
}
