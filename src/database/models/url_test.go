package models

import (
	"log"
	"shorturl/src/database"
	"shorturl/src/utils"
	"testing"
)

func TestUrlModel_Create(t *testing.T) {
	dbInstance := &database.DBInstance{}
	dbInstance.Connect()
	urlModel := &UrlModel{DBInstance: dbInstance}

	t.Run("Create url", func(t *testing.T) {
		original := "https://google.com/"
		key := utils.GenerateKey(10)
		err := urlModel.Create(original, key)
		if err != nil {
			t.Errorf("Create() error = %v", err)
		}
	})
}

func TestUrlModel_GetByKey(t *testing.T) {
	dbInstance := &database.DBInstance{}
	dbInstance.Connect()
	urlModel := &UrlModel{DBInstance: dbInstance}

	original := "https://youtube.com/"
	key := utils.GenerateKey(10)
	err := urlModel.Create(original, key)
	if err != nil {
		log.Fatalln(err)
	}

	t.Run("Get by short url", func(t *testing.T) {
		um, err := urlModel.GetByKey(key)
		if err != nil {
			t.Errorf("GetByKey() error = %v", err)
		} else {
			if um.Original != original {
				t.Errorf("GetByKey() got = %v, want = %v", um.Original, original)
			}
		}
	})

	t.Run("Get by non-existent short url", func(t *testing.T) {
		um, err := urlModel.GetByKey("qwertyuiop")
		if err == nil {
			t.Errorf("GetByKey() got = %v, wantErr = %v", um.Original, true)
		}
	})
}

func TestUrlModel_GetByOriginal(t *testing.T) {
	dbInstance := &database.DBInstance{}
	dbInstance.Connect()
	urlModel := &UrlModel{DBInstance: dbInstance}

	original := "https://reddit.com/"
	key := utils.GenerateKey(10)
	err := urlModel.Create(original, key)
	if err != nil {
		log.Fatalln(err)
	}

	t.Run("Get by original url", func(t *testing.T) {
		um, err := urlModel.GetByOriginal(original)
		if err != nil {
			t.Errorf("GetByOriginal() error = %v", err)
		} else {
			if um.Key != key {
				t.Errorf("GetByOriginal() got = %v, want = %v", um.Key, key)
			}
		}
	})

	t.Run("Get by non-existent url", func(t *testing.T) {
		um, err := urlModel.GetByOriginal("https://qwertyuiop.com")
		if err == nil {
			t.Errorf("GetByOriginal() got = %v, wantErr = %v", um.Key, true)
		}
	})
}
