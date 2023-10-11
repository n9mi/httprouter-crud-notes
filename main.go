package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/naomigrain/httprouter-crud-notes/app/database"
	"github.com/naomigrain/httprouter-crud-notes/app/router"
	"github.com/naomigrain/httprouter-crud-notes/config"
	"github.com/naomigrain/httprouter-crud-notes/exception"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConfig := config.GetDBConfig(true)
	db := database.NewDB(dbConfig)
	validate := validator.New()

	router := router.NewRouter(db, validate)
	router.PanicHandler = exception.PanicHandler

	serverConfig := config.GetServerConfig(true)
	server := http.Server{
		Addr:    serverConfig.Host + ":" + serverConfig.Port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println(dbConfig, db, validate)
}
