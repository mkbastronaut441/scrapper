package main

import (
	"fmt"
	"log"
	"net/http"
	db "scrapper/internal/adapter/persistence/db"
	"scrapper/internal/config"
)

func main() {
	//load config
	fmt.Println("running server...")

	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	dbconn := db.NewDbConnection(&conf.Db)
	fmt.Println(dbconn)
	_, err = dbconn.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Server started on port ", conf.Port)
	http.ListenAndServe(conf.Port, nil)
}
