package main

import (
	"database/sql"
	"fmt"
	"log"

	"go-template-microservice/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig("deployments/")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, 5432, config.DBUsername, config.DBPassword, config.DBName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
