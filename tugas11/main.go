package main

import (
	"log"

	_ "tesjwt.go/docs"

	"tesjwt.go/database"
	"tesjwt.go/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	log.Println("starting app...")
	r.Run(":5000")
}
