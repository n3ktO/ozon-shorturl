package services

import (
	"log"
	"shorturl/src/database"
	"shorturl/src/database/models"
	"shorturl/src/grpcgen"
	"testing"
)

func TestUrlService_Create(t *testing.T) {
	dbInstance := &database.DBInstance{}
	dbInstance.Connect()
	urlModel := &models.UrlModel{DBInstance: dbInstance}
	service := UrlService{UrlModel: urlModel}

	t.Run("Right url test", func(t *testing.T) {
		_, err := service.Create(nil, &grpcgen.Url{Url: "https://twitter.com/"})
		if err != nil {
			t.Errorf("Create() error = %v", err)
		}
	})

	t.Run("Not url test", func(t *testing.T) {
		key, err := service.Create(nil, &grpcgen.Url{Url: "hello"})
		if err == nil {
			t.Errorf("Create() got = %v, wantErr = %v", key.Key, true)
		}
	})
}

func TestUrlService_Get(t *testing.T) {
	dbInstance := &database.DBInstance{}
	dbInstance.Connect()
	urlModel := &models.UrlModel{DBInstance: dbInstance}
	service := UrlService{UrlModel: urlModel}

	t.Run("Right key test", func(t *testing.T) {
		um, err := urlModel.GetByOriginal("https://google.com/")
		if err != nil {
			log.Fatalln(err)
		}

		_, err = service.Get(nil, &grpcgen.Key{Key: um.Key})
		if err != nil {
			t.Errorf("Get() error = %v", err)
		}
	})

	t.Run("Wrong key test", func(t *testing.T) {
		url, err := service.Get(nil, &grpcgen.Key{Key: "Hello"})
		if err == nil {
			t.Errorf("Get() got = %v, wantErr = %v", url.Url, true)
		}
	})
}
