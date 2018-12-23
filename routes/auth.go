package routes

import (
	"encoding/json"
	"moff-backend/services"
	"net/http"
)

type loginJSON struct {
	googleID string `json:"id"`
}

type loginJSONRes struct {
	token string `json:"token"`
}

func login(w http.ResponseWriter, r *http.Response) {
	var body loginJSON
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid login format", http.StatusBadRequest)
	}
	var token, err = services.Login(body.googleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var resp = loginJSONRes{token}
	json.NewEncoder(w).Encode(resp)
}
