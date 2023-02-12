package main

import (
	"fmt"
)

type AuthorID int

type Author struct {
	ID   AuthorID
	Name string
}

func (ID AuthorID) Valid() bool {
	return ID > 0
}

func GetAuthor(id AuthorID) (*Author, error) {
	if !id.Valid() {
		// errors.New: エラーを作成する verbを使う場合はfmt.Errorfを使う
		//return nil, errors.New("GetAuthor: invalid ID")

		return nil, fmt.Errorf("GetAuthor: invalid ID: %d", id)
	}
	return nil, nil
}

type Book struct {
	AuthorID AuthorID
}

func GetAuthorName(b *Book) (string, error) {
	author, err := GetAuthor(b.AuthorID)
	if err != nil {
		return "", err
	}
	return author.Name, nil
}

func main() {
	GetAuthor(0)
	GetAuthorName(&Book{AuthorID: 0})
}
