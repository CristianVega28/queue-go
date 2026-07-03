package main

import (
	"fmt"
	"queue-go/server"
)

func main() {
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
