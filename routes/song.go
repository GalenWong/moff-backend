package routes

import (
	"encoding/json"
	"net/http"
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
