package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mhayk/GO-Expert-temp-by-cep/api/handler"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENVIRONMENT") != "production" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	router := http.NewServeMux()

	handler.NewWeatherHandler(router)


	// log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":8080", router))
}
