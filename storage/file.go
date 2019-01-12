package storage

import (
	"io"

	"google.golang.org/api/drive/v3"
)

func (ds *DriveSession) createFile(
	file *drive.File,
	r io.Reader,
) (*drive.File, error) {
	var c = ds.client
	var newfile, err = c.Files.Create(file).Media(r).Do()
	if err != nil {
		return newfile, err
	}
	return newfile, nil
}

func (ds *DriveSession) deleteFile(fileID string) error {
	var c = ds.client
	return c.Files.Delete(fileID).Do()
}

func (ds *DriveSession) getFileMeta(fileID string) (*drive.File, error) {
	var c = ds.client
	return c.Files.Get(fileID).Do()
}

func (ds *DriveSession) getFileStream(fileID string) (io.ReadCloser, error) {
	var c = ds.client
	resp, err := c.Files.Get(fileID).Download()
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
