package storage

import (
	"io"

	"google.golang.org/api/drive/v3"
)

func (ds *DriveSession) GetAudio(id string) (io.ReadCloser, error) {
	return ds.getFileStream(id)
}

func (ds *DriveSession) CreateAudio(r io.Reader) (string, error) {
	f := &drive.File{
		Parents: []string{ds.audios.Id},
	}
	newfile, err := ds.createFile(f, r)
	if err != nil {
		return "", err
	}
	return newfile.Id, nil
}

func (ds *DriveSession) DeleteAudio(id string) error {
	return ds.deleteFile(id)
}
