package storage

import (
	"context"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

const tokFile = "credentials.json"

var scopes = []string{drive.DriveAppdataScope}

var config *oauth2.Config

func init() {
	// reads secret file
	config = getAPIConfig(tokFile)
}

/*
getAPIConfig should not be called elsewhere other then init
*/
func getAPIConfig(secretPath string) *oauth2.Config {
	b, err := ioutil.ReadFile(secretPath)
	if err != nil {
		log.Fatalf("Cannot read client secret file: %v", secretPath)
	}
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		log.Fatalf("Cannot parse client secret file to config: %v", secretPath)
	}
	return config
}

func getClient(userToken *oauth2.Token, config *oauth2.Config) (*drive.Service, error) {
	httpclient := config.Client(context.Background(), userToken)
	return drive.New(httpclient)
}

/*
DriveSession is a wrapper that wraps around the google drive client.
*/
type DriveSession struct {
	client     *drive.Service
	thumbnails *drive.File
	audios     *drive.File
}

func GetDriveSession(userToken *oauth2.Token) (*DriveSession, error) {
	client, err := getClient(userToken, config)
	if err != nil {
		return nil, err
	}
	var drive = &DriveSession{client, nil, nil}
	return drive, nil
}
