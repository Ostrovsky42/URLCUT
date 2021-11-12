package httphandler

import (
	"URLCUT/app"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type urlGenerator struct {
	service app.UrlCutterServ
}

func NewUrlGenerator(service app.UrlCutterServ) *urlGenerator {
	return &urlGenerator{service: service}
}

// GetUrl @Summary GetUrl
// @Description get url by key
// @ID get-url-by-key
// @Accept  json
// @Produce  json
// @Param key path string true "url"
// @Success 301
// Header 200 {string} Token "qwerty"
// Failure 400,404 {object} string
// Failure 500 {object} string
// Failure default {object} string
// @Router /{key} [get]
func (u *urlGenerator) GetUrl(ctx echo.Context) error {
	key := ctx.Param("key")
	if key == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("key invalid"))
	}
	url, err := u.service.GetURL(key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.Redirect(http.StatusMovedPermanently, url)
	//return ctx.JSON(http.StatusOK, url)"key invalid"
}

// UrlCutter @Summary CutURL
// @Description cut url and get key
// @ID cut-url-and-get-key
// @Accept  json
// @Produce  json
// @Param longUrl body UrlToCut true "cut url"
// @Success 200 {object} string "ok"
// Header 200 {string} Token "qwerty"
// Failure 400,404 {object} string
// Failure 500 {object} string
// Failure default {object} string
// @Router /urlcut [post]
func (u *urlGenerator) UrlCutter(ctx echo.Context) error {
	var inputUrl UrlToCut
	err := ctx.Bind(&inputUrl)
	if err != nil {
		log.Print("bind returned error")
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	key, err := u.service.MakeKey(inputUrl.LongUrl)
	if err != nil {
		log.Print("makeKey error")
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, key)
}

type UrlToCut struct {
	LongUrl string `json:"long_url"`
}
