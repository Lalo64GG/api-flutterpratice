package main

import (
	"api/cmd/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("HOST_SERVER")
	port := os.Getenv("PORT_SERVER")
	useHTTPS := os.Getenv("USE_HTTPS") == "true"
	certFile := os.Getenv("SSL_CERT_FILE")
	keyFile := os.Getenv("SSL_KEY_FILE")

	srv := server.NewServer(host, port, useHTTPS, certFile, keyFile)
	srv.Run()
}
