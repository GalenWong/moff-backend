package routes

import (
	"encoding/json"
	"moff-backend/services"
	"net/http"
)

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
