package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Encoding int

const (
	XML Encoding = iota
	JSON
)

type User struct {
	UserID string `xml:"userId" json:"userId"`
	Email  string `xml:"email" json:"email"`
	Score  int    `xml:"score" json:"score"`
}

var Users []User

func (u *User) Equal(v User) bool {
	if u.UserID != v.UserID || u.Email != v.Email || u.Score != v.Score {
		return false
	}
	return true
}

func (u *User) encode(format Encoding) ([]byte, error) {
	var encoded []byte = nil
	var err error
	switch format {
	case XML:
		encoded, err = xml.Marshal(u)
	case JSON:
		encoded, err = json.Marshal(u)
	default:
		errors.New("unknow encoding format")
	}
	return encoded, err
}

func (u *User) fromEncoded(format Encoding, encoded []byte) error {
	recovered := User{}
	var err error
	switch format {
	case XML:
		err = xml.Unmarshal(encoded, &recovered)
	case JSON:
		err = json.Unmarshal(encoded, &recovered)
	default:
		err = errors.New("unknow encoding format")
	}

	if err == nil {
		*u = recovered
	}

	return err
}

func (u *User) write(encoded []byte, path string) error {
	err := ioutil.WriteFile(path, encoded, os.ModePerm)
	return err
}

func (u *User) read(path string) ([]byte, error) {
	encoded, err := ioutil.ReadFile(path)
	return encoded, err
}

func (u *User) ToEncodedFile(format Encoding, filePath string) error {
	encoded, err := u.encode(format)
	if err != nil {
		return err
	}
	err = u.write(encoded, filePath)
	return err
}

func (u *User) FromEncodedFile(format Encoding, filePath string) error {
	encoded, err := u.read(filePath)
	if err != nil {
		return err
	}
	err = u.fromEncoded(format, encoded)
	return err
}

func testWriteXML(t *testing.T) {
	tmpDir := os.TempDir()
	for _, u := range Users {
		f := tmpDir + u.UserID + ".xml"
		err := u.ToEncodedFile(XML, f)
		if err != nil {
			t.Error(err)
		}
	}
}

func testWriteJSON(t *testing.T) {
	tmpDir := os.TempDir()
	for _, u := range Users {
		f := tmpDir + u.UserID + ".json"
		err := u.ToEncodedFile(JSON, f)
		if err != nil {
			t.Error(err)
		}
	}
}

func testReadXML(t *testing.T) {
	tmpDir := os.TempDir()
	for _, u := range Users {
		f := tmpDir + "/" + u.UserID + ".xml"
		newUSer := User{}
		err := newUSer.FromEncodedFile(XML, f)
		if err != nil {
			t.Error(err)
		}
		if !newUSer.Equal(u) {
			t.Errorf(fmt.Sprintf("found %v, expected %v", newUSer, u))
		}
	}
}

func testReadJSON(t *testing.T) {
	tmpDir := os.TempDir()
	for _, u := range Users {
		f := tmpDir + "/" + u.UserID + ".json"
		newUSer := User{}
		err := newUSer.FromEncodedFile(JSON, f)
		if err != nil {
			t.Error(err)
		}
		if !newUSer.Equal(u) {
			t.Errorf(fmt.Sprintf("found %v, expected %v", newUSer, u))
		}
	}
}

func testXML(t *testing.T) {
	t.Run("Action=Write", testWriteXML)
	t.Run("Action=Read", testReadXML)

	tmpDir := os.TempDir()
	for _, u := range Users {
		f := tmpDir + "/" + u.UserID + ".xml"
		_ = os.Remove(f)
	}
}

func testJSON(t *testing.T) {
	t.Run("Action=Write", testWriteJSON)
	t.Run("Action=Read", testReadJSON)

	tmpDir := os.TempDir()
	for _, u := range Users {
		f := tmpDir + "/" + u.UserID + ".json"
		_ = os.Remove(f)
	}
}

func TestEncoding(t *testing.T) {
	t.Run("Action=XML", testXML)
	t.Run("Action=JSON", testJSON)
}

func TestMain(m *testing.M) {
	UserA := User{"UserA", "usera@gmail.com", 42}
	UserB := User{"UserB", "userb@gmail.com", 43}
	Users = []User{UserA, UserB}

	os.Exit(m.Run())
}
