package databases

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func GetUserSongList(userID string) ([]Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := GetUserDB().FindOne(ctx, bson.M{"id": userID})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var user User
	if err := res.Decode(&user); err != nil {
		return nil, err
	}
	return user.SongList, nil
}

func CreateUser(new *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := GetUserDB().InsertOne(ctx, new)
	return err
}

func UserExists(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := GetUserDB().FindOne(ctx, bson.M{"id": id})
	var existing User
	if err := res.Decode(&existing); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		log.Printf(err.Error())
		return true, err
	}
	return true, nil
}

func FindUser(id string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := GetUserDB().FindOne(ctx, bson.M{"id": id})
	var u User
	if err := res.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
