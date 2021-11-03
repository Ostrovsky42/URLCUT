package main

import (
	"URLCUT/app"
	"URLCUT/infrastucture/httphandler"
	"URLCUT/infrastucture/localservices"
	"URLCUT/infrastucture/repo"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "********"
	dbname   = "ShortUrl"
)

func main() {
	e := echo.New()
	db := ConnectDB()
	defer db.Close()

	repository := repo.NewKeySaver(db)
	keyGeneratorService := localservices.NewKeyGenerator()
	service := app.NewURLCutterService(*keyGeneratorService, repository)
	handler := httphandler.NewUrlGenerator(service)
	e.POST("/urlcut", handler.UrlCutter)
	e.GET("/:key", handler.GetUrl)
	e.Logger.Fatal(e.Start(":8088"))
}

func ConnectDB() (db *sql.DB) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// check db
	err = db.Ping()
	CheckError(err)
	return
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
