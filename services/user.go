package services

import (
	"errors"
	"moff-backend/databases"
)

var ErrCreateDuplicateUser = errors.New("User already exists")

func CreateNewUser(id string, name string) error {
	var newUser = databases.User{
		ID:       id,
		Name:     name,
		SongList: []databases.Song{},
	}
	if exists, err := databases.UserExists(id); exists && err == nil {
		return ErrCreateDuplicateUser
	} else if err != nil {
		return err
	}
	return databases.CreateUser(&newUser)
}
