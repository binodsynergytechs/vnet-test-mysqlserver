package main

import (
	"log"

	"vnet-mysql/config"
	"vnet-mysql/pkg/db"
	"vnet-mysql/pkg/routes"
)

func main() {
	config.LoadDotEnv()
	db, err := db.InitDB()
	if err != nil {
		log.Panicf("Error in connecting db %v", err.Error())
	}
	log.Println("Connected to db", db)
	routes.InitRoutes(db)
}
