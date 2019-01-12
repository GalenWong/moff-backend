package databases

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"golang.org/x/oauth2"
)

type Song struct {
	ID string `json:"id" bson:"id"`
}

type User struct {
	OID      primitive.ObjectID `json:"oid" bson:"_id"`
	ID       string             `json:"id" bson:"id"`
	Name     string             `json:"name" bson:"name"`
	SongList []Song             `json:"songlist" bson:"songlist"`
	Token    *oauth2.Token      `json:"token" bson:"token"`
}
