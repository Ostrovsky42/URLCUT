package app

import (
	"URLCUT/infrastucture/localservices"
	"github.com/labstack/gommon/log"
)

type UrlCutterService struct {
	keyGenerator localservices.GenerateInterface
	repo         UrlSaver
}

func (u *UrlCutterService) MakeKey(url string) (string, error) {
	key := u.GetUniqueKey()
	err := u.repo.Save(url, key)
	if err != nil {
		log.Print("URL was not saved", err)
		return "", err
	}
	return key, nil
}

func (u *UrlCutterService) GetURL(key string) (string, error) {
	longUrl, err := u.repo.Get(key)
	if err != nil {
		log.Print("no such url", err)
		return "", err
	}
	return longUrl, nil
}

func NewURLCutterService(keyGenerator localservices.GenerateInterface, repo UrlSaver) *UrlCutterService {
	return &UrlCutterService{keyGenerator: keyGenerator, repo: repo}
}

func (u *UrlCutterService) GetUniqueKey() string {
	var key string
	for {
		key = u.keyGenerator.GenerateKey()
		_, err := u.repo.Get(key)
		if err == nil {
			break
		}
	}
	return key
}
