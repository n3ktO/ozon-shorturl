package models

import (
	"shorturl/src/database"
)

type UrlModel struct {
	DBInstance *database.DBInstance
}

type Url struct {
	Id       int    `db:"id"`
	Key      string `db:"key"`
	Original string `db:"original"`
}

func (um *UrlModel) Create(original string, key string) error {
	query := "INSERT INTO url (original, key) VALUES ($1, $2)"
	_, err := um.DBInstance.DB.Exec(query, original, key)
	if err != nil {
		return err
	}
	return nil
}

func (um *UrlModel) GetByKey(key string) (*Url, error) {
	url := &Url{}
	query := "SELECT * FROM url WHERE key=$1"
	err := um.DBInstance.DB.Get(url, query, key)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (um *UrlModel) GetByOriginal(original string) (*Url, error) {
	url := &Url{}
	query := "SELECT * FROM url WHERE original=$1"
	err := um.DBInstance.DB.Get(url, query, original)
	if err != nil {
		return nil, err
	}
	return url, nil
}
