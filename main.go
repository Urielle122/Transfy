package main

import (
	"log"
	"net/http"
	"transfy/core"
	"transfy/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		// log.Fatalf("Error loading .env file: %s", err)
	}
	core.InitConnection()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /add", routes.AddUser)
	log.Println("Serveur démarré sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
