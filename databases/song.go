package databases

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

func GetUserSongList(googleID string) ([]Song, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := GetUserDB().FindOne(ctx, bson.M{"id": googleID})
	if res.Err() != nil {
		return nil, res.Err()
	}
	var user User
	if err := res.Decode(&user); err != nil {
		return nil, err
	}
	return user.SongList, nil
}

func CreateSong(googleID string, new *Song) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := GetUserDB().UpdateOne(ctx,
		bson.M{"id": googleID},
		bson.D{
			{
				"$push", bson.D{{"songlist", new}},
			},
		})
	return err
}

func DeleteSong(googleID string, songID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := GetUserDB().UpdateOne(ctx,
		bson.M{"id": googleID},
		bson.D{
			{
				"$pull", bson.D{{"songlist.id", songID}},
			},
		})

	return err
}
