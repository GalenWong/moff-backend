package services

import (
	"errors"
	"fmt"
	"log"
	"moff-backend/databases"

	jwt "github.com/dgrijalva/jwt-go"
)

var ErrUserNotExist = errors.New("Cannot Find User")

// TODO: change to read config
var secret = []byte("deadbeef314moff")

func Login(googleID string) (string, error) {
	if exists, err := databases.UserExists(googleID); !exists && err == nil {
		return "", ErrUserNotExist
	} else if err != nil {
		return "", err
	}
	var user, err = databases.FindUser(googleID)
	if err != nil {
		return "", err
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func decodeToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
}

func ValidateToken(tokenStr string) (bool, error) {
	token, err := decodeToken(tokenStr)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Println(claims["id"])
		return true, nil
	}
	return false, err
}

func ValidateAndGetID(tokenStr string) (string, error) {
	token, err := decodeToken(tokenStr)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string), nil
	}
	return "", err
}
