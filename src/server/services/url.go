package services

import (
	"context"
	"errors"
	urllib "net/url"
	"shorturl/src/database/models"
	"shorturl/src/grpcgen"
	"shorturl/src/utils"
)

type UrlService struct {
	grpcgen.UnimplementedShortUrlServer
	UrlModel *models.UrlModel
}

func (s *UrlService) Create(ctx context.Context, url *grpcgen.Url) (*grpcgen.Key, error) {
	_, err := urllib.ParseRequestURI(url.Url)
	if err != nil {
		return nil, errors.New("only url accepted")
	}

	var key string
	if um, err := s.UrlModel.GetByOriginal(url.Url); err != nil {
		for {
			key = utils.GenerateKey(10)
			_, err = s.UrlModel.GetByKey(key)
			if err != nil {
				break
			}
		}

		err = s.UrlModel.Create(url.Url, key)
		if err != nil {
			return nil, errors.New("service error")
		}
	} else {
		key = um.Key
	}

	return &grpcgen.Key{Key: key}, nil
}

func (s *UrlService) Get(ctx context.Context, key *grpcgen.Key) (*grpcgen.Url, error) {
	um, err := s.UrlModel.GetByKey(key.Key)
	if err != nil {
		return nil, errors.New("url does not exist")
	}

	return &grpcgen.Url{Url: um.Original}, nil
}
