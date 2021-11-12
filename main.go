package main

import (
	"URLCUT/app"
	_ "URLCUT/docs"
	"URLCUT/infrastucture/httphandler"
	"URLCUT/infrastucture/localservices"
	"URLCUT/infrastucture/repo"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	//password = "********"
	password = "Xcritical79"
	dbname   = "ShortUrl"
)

// @title ShortURL API
// @version 1.0
// @description Accepts a link and returns an abbreviated version

// host localhost:8000
// @BasePath /

func main() {
	e := echo.New()
	db := connectDB()
	defer db.Close()
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repository := repo.NewKeySaver(db)
	keyGeneratorService := localservices.NewKeyGenerator()
	service := app.NewURLCutterService(keyGeneratorService, repository)
	handler := httphandler.NewUrlGenerator(service)
	if handler != nil {

	}
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.POST("/urlcut", handler.UrlCutter)
	e.GET("/:key", handler.GetUrl)
	e.Logger.Fatal(e.Start(viper.GetString("port")))
}

func connectDB() (db *sql.DB) {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//viper.GetString("db.host"),viper.GetString("db.port"),viper.GetString("db.user"),password,viper.GetString("db.dbname"))
	// open database
	db, err := sql.Open("postgres", psqlconn)
	checkError(err)
	// check db
	err = db.Ping()
	checkError(err)
	return
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
