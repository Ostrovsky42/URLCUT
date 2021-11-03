package httphandler

import (
	"URLCUT/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type urlGenerator struct {
	service app.KeyGenerator
}

func NewUrlGenerator(service app.KeyGenerator) *urlGenerator {
	return &urlGenerator{service: service}
}

func (u *urlGenerator) GetUrl(ctx echo.Context) error {
	key := ctx.Param("key")
	if key == "" {
		return ctx.JSON(http.StatusBadRequest, "key invalid")
	}

	url, err := u.service.GetURL(key)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.Redirect(http.StatusMovedPermanently, url)
}

func (u *urlGenerator) UrlCutter(c echo.Context) error {
	var urlToCut UrlToCut
	err := c.Bind(&urlToCut)
	if err != nil {
		log.Print("bind returned error")
		return c.String(http.StatusBadRequest, err.Error())
	}
	key, err := u.service.MakeKey(urlToCut.LongUrl)
	return c.JSON(http.StatusOK, key)
}

type UrlToCut struct {
	LongUrl string `json:"long_url"`
}
