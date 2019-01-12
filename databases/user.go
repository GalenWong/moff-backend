package databases

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func CreateUser(new *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := GetUserDB().InsertOne(ctx, new)
	return err
}

func UserExists(googleID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := GetUserDB().FindOne(ctx, bson.M{"id": googleID})
	var existing User
	if err := res.Decode(&existing); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		log.Printf(err.Error())
		return true, err
	}
	return true, nil
}

func FindUser(googleID string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := GetUserDB().FindOne(ctx, bson.M{"id": googleID})
	var u User
	if err := res.Decode(&u); err != nil {
		return nil, err
	}
	return &u, nil
}
