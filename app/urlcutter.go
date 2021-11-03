package app

import (
	"URLCUT/infrastucture/localservices"
	"github.com/labstack/gommon/log"
)

type urlCutterService struct {
	keyGenerator localservices.KeyGenerator
	repo         UrlSaver
}

func (u *urlCutterService) MakeKey(url string) (string, error) {
	key := GetUniqueKey(u)
	err := u.repo.Save(url, key)
	if err != nil {
		log.Print("URL was not saved", err)
		return "", err
	}
	return key, err
}

func (u *urlCutterService) GetURL(key string) (string, error) {
	longUrl, err := u.repo.Get(key)
	if err != nil {
		log.Print("no such url", err)
		return "", err
	}
	return longUrl, nil
}

func NewURLCutterService(keyGenerator localservices.KeyGenerator, repo UrlSaver) *urlCutterService {
	return &urlCutterService{keyGenerator: keyGenerator, repo: repo}
}

func GetUniqueKey(u *urlCutterService) string {
	keys := u.repo.GetKeys()
	key := u.keyGenerator.GenerateKey()
	for _, dbKey := range keys {
		if key == dbKey {
			key = u.keyGenerator.GenerateKey()
			//todo :logic
		}
	}
	return key
}
