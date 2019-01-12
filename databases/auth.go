package databases

import (
	"context"
	"fmt"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"golang.org/x/oauth2"
)

func GetUserAuth(googleID string) (*oauth2.Token, error) {
	user, err := FindUser(googleID)
	if err != nil {
		return nil, err
	}
	return user.Token, nil
}

func UpdateUserAuth(googleID string, tok *oauth2.Token) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := GetUserDB().UpdateOne(ctx,
		bson.M{"id": googleID},
		bson.D{
			{"$set", bson.D{{"token", tok}}},
		})
	if err != nil {
		return err
	}

	if res.ModifiedCount != 1 {
		return fmt.Errorf("Updated %v User documents auth token", res.ModifiedCount)
	}
	if res.MatchedCount != 1 {
		return fmt.Errorf("Updated %v User documents auth token", res.MatchedCount)
	}
	return nil
}
