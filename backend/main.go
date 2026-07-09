package main

import (
	"fmt"
	"os"
	"queue-go/models"
	"queue-go/server"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		err := models.Setup()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("database queue-go.db created successfully")
		return
	}

	server.Http.New()
	server.Routes()

	srv := server.Server{
		Port:     "8080",
		Handlers: server.Http.ServerMux(),
	}

	err := srv.Start()

	if err != nil {
		fmt.Println(err.Error())
	}
}
