package routes

import (
	"encoding/json"
	"io"
	"moff-backend/databases"
	"moff-backend/services"
	"net/http"

	"github.com/gorilla/mux"
)

/*
SongList is the endpoint that returns the song list for a specific user
*/
func songList(w http.ResponseWriter, r *http.Request) {
	userID, err := retrieveUserIDFromCtx(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	ls, err := services.GetSongList(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var resp = struct {
		List []databases.Song
	}{List: ls}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func uploadSong(w http.ResponseWriter, r *http.Request) {
	userID, err := retrieveUserIDFromCtx(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if r.Body == nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.UploadSong(userID, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

/**
 * /song/{id}
 */
func downloadSong(w http.ResponseWriter, r *http.Request) {
	userID, err := retrieveUserIDFromCtx(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	var songID = mux.Vars(r)["id"]
	if songID == "" {
		http.Error(w, "empty id", http.StatusBadRequest)
		return
	}

	song, err := services.DownloadSong(userID, songID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if song == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer song.Close()

	if _, err := io.Copy(w, song); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/**
 * /song/{id}
 */
func deleteSong(w http.ResponseWriter, r *http.Request) {
	userID, err := retrieveUserIDFromCtx(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	var songID = mux.Vars(r)["id"]
	if songID == "" {
		http.Error(w, "empty id", http.StatusBadRequest)
		return
	}

	if err := services.RemoveSong(userID, songID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
