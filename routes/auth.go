package routes

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"moff-backend/services"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

const userIDctx = "userID"

type loginJSON struct {
	GoogleID string `json:"id"`
}

type loginJSONRes struct {
	Token string `json:"token"`
}

func login(w http.ResponseWriter, r *http.Request) {
	var body loginJSON
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid login format", http.StatusBadRequest)
	}
	var token, err = services.Login(body.GoogleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var resp = loginJSONRes{token}
	json.NewEncoder(w).Encode(resp)
}

var invalidAuthHeader = errors.New("Invalid Authorization Header")

func extractAuthToken(h *http.Header) (string, error) {
	var authHeader = h.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("No Authorization Header Provided")
	}
	var headerSplit = strings.Split(authHeader, " ")
	if len(headerSplit) < 2 || headerSplit[0] != "Bearer" || headerSplit[1] == "" {
		return "", invalidAuthHeader
	}
	return headerSplit[1], nil

}
func validate(w http.ResponseWriter, r *http.Request) {
	var authHeader, err = extractAuthToken(&r.Header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	if valid, err := services.ValidateToken(authHeader); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if valid == false {
		http.Error(w, invalidAuthHeader.Error(), http.StatusForbidden)
		return
	}

	w.Write([]byte("Valid"))
}

/**
AuthMidware checks and validates the Authorization header.
If the token is valid, the user ID is set in context
*/
func authMidware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authHeader, err = extractAuthToken(&r.Header)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		if id, err := services.ValidateAndGetID(authHeader); err != nil || id == "" {
			http.Error(w, invalidAuthHeader.Error(), http.StatusForbidden)
			return
		} else {
			ctx := context.WithValue(r.Context(), userIDctx, id)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

func retrieveUserIDFromCtx(r *http.Request) (string, error) {
	userID, ok := r.Context().Value(userIDctx).(string)
	if !ok || userID == "" {
		return "", errors.New("Cannot retrieve userID from ctx")
	}
	return userID, nil
}

func updateAuthToken(w http.ResponseWriter, r *http.Request) {
	userID, err := retrieveUserIDFromCtx(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	var tok oauth2.Token
	if err := json.NewDecoder(r.Body).Decode(&tok); err != nil {
		log.Println(err)
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	if err := services.UpdateUserAuth(userID, &tok); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
