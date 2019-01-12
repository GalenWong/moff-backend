package storage

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
)

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func TestRun(t *testing.T) {
	token, err := tokenFromFile("token.json")
	if err != nil {
		log.Fatal(err)
	}

	ds, err := GetDriveSession(token)
	if err != nil {
		log.Fatal("Fail to init session " + err.Error())
	}
	r := strings.NewReader("file content")
	newfile := &drive.File{
		Name: "testfile",
		// Spaces:  []string{moffFolderID},
		Parents: []string{moffFolderID},
	}
	f, err := ds.createFile(newfile, r)
	if err != nil {
		panic(err)
	}
	log.Println(f)
	if f == nil {
		panic("f is nil")
	}
	fmeta, err := ds.getFileMeta(f.Id)
	if err != nil {
		panic(err)
	}
	if fmeta == nil {
		panic("fmeta is nil")
	}
	log.Printf("creation id: %v, getMeta id: %v", f.Id, fmeta.Id)
	if f.Id != fmeta.Id {
		panic("unequal ids")
	}

	fstream, err := ds.getFileStream(f.Id)
	if err != nil {
		panic(err)
	}
	defer fstream.Close()
	content := make([]byte, 15)
	_, err = fstream.Read(content)
	if err != nil {
		panic(err)
	}
	log.Println(string(content))
}
