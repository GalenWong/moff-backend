package storage

import (
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

// using appdata scope, we access an aliased folder that is hidden from user
const moffFolderID = "appDataFolder"
const thumbnailFolderName = "thumbnails"
const audioFolderName = "audios"
const folderMime = "application/vnd.google-apps.folder"

/*
CheckOrCreateFolder checks if the appdata folder is correctly setup.
Call this function before any other drive operation
*/
func (ds *DriveSession) CheckOrSetupDir() error {
	a, err := ds.getOrCreateFolder(audioFolderName)
	if err != nil {
		return err
	}
	t, err := ds.getOrCreateFolder(thumbnailFolderName)
	if err != nil {
		return err
	}
	ds.audios = a
	ds.thumbnails = t
	return nil
}

func (ds *DriveSession) getOrCreateFolder(folderName string) (*drive.File, error) {
	f, err := ds.getFolder(folderName)
	if err != nil {
		return nil, err
	} else if f == nil {
		return ds.createFolder(folderName)
	}
	return f, nil
}

/*
get folder under moffFolder, return nil, nil if not found
*/
func (ds *DriveSession) getFolder(folderName string) (*drive.File, error) {
	var c = ds.client
	var query = fmt.Sprintf(
		"'%v' in parents and name = '%v' and mimeType = '%v'",
		moffFolderID,
		folderName,
		folderMime,
	)
	list, err := c.Files.List().Spaces(moffFolderID).Q(query).Do()
	if err != nil {
		return nil, err
	}

	switch len(list.Files) {
	case 0:
		return nil, nil
	case 1:
		return list.Files[0], nil
	default: // it is impossible to have more than 1 folder
		return nil, fmt.Errorf("More than 1 '%v' folder", folderName)
	}
}

func (ds *DriveSession) createFolder(folderName string) (*drive.File, error) {
	var c = ds.client
	var newFolder = drive.File{}
	newFolder.Name = folderName
	newFolder.MimeType = folderMime
	newFolder.Parents = []string{moffFolderID}

	folder, err := c.Files.Create(&newFolder).Do()
	if err != nil {
		return nil, err
	}
	return folder, nil
}

func (ds *DriveSession) lsAppData() []*drive.File {
	var c = ds.client
	list, err := c.Files.List().Spaces(moffFolderID).PageSize(10).Do()
	if err != nil {
		log.Fatal(err)
	}
	return list.Files
}
