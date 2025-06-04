package main

import (
	"api/cmd/server"
	"os"
)

var (
	HOST = os.Getenv("HOST_SERVER")
	PORT = os.Getenv("PORT_SERVER")

)

func main(){
	srv := server.NewServer(HOST, PORT)
	srv.Run()
}