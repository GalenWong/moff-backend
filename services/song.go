package services

import (
	"io"
	"moff-backend/databases"
	"moff-backend/storage"
)

func UploadSong(userID string, r io.Reader) error {
	ds, err := getValidDriveSession(userID)
	if err != nil {
		return err
	}
	fileID, err := ds.CreateAudio(r)
	if err != nil {
		return err
	}
	newSong := &databases.Song{ID: fileID}

	return addToSongList(userID, newSong)
}

func DownloadSong(userID string, songID string) (io.ReadCloser, error) {
	ds, err := getValidDriveSession(userID)
	if err != nil {
		return nil, err
	}
	return ds.GetAudio(songID)
}

func RemoveSong(userID string, songID string) error {
	ds, err := getValidDriveSession(userID)
	if err != nil {
		return err
	}
	if err := ds.DeleteAudio(songID); err != nil {
		return err
	}
	if err := databases.DeleteSong(userID, songID); err != nil {
		return err
	}
	return nil
}

func GetSongList(userID string) ([]databases.Song, error) {
	return databases.GetUserSongList(userID)
}

func getValidDriveSession(userID string) (*storage.DriveSession, error) {
	ds, err := getUserDriveSession(userID)
	if err != nil {
		return nil, err
	}
	if err := ds.CheckOrSetupDir(); err != nil {
		return nil, err
	}
	return ds, nil
}

func getUserDriveSession(userID string) (*storage.DriveSession, error) {
	tok, err := databases.GetUserAuth(userID)
	if err != nil {
		return nil, err
	}
	return storage.GetDriveSession(tok)
}

func addToSongList(userID string, song *databases.Song) error {
	return databases.CreateSong(userID, song)
}
