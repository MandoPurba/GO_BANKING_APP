package main

import (
	"github.com/MandoPurba/rest-api/apps"
	"github.com/MandoPurba/rest-api/config"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	// SETUP .ENV
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed add file env")
	}

	// CONNECT DATABASE
	db := config.InitDB()
	defer db.Close()

	// ADD ROUTER
	router := apps.InitRouter()
	addr := ":8000"

	log.Printf("Server started on %s\n", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(err)
	}
}
