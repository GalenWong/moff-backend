package databases

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type userDB struct {
	Client *mongo.Client
}

var userDBinst *userDB

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	userDBinst = &userDB{Client: c}
}

func GetUserDB() *mongo.Collection {
	// may replace with diff db for diff purpose
	return userDBinst.Client.Database("testing").Collection("User")
}
