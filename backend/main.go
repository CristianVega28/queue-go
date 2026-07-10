package main

import (
	"fmt"
	"queue-go/models"
	"queue-go/server"
)

func main() {
	err := models.Setup()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	server.Http.New()
	server.Routes()

	srv := server.Server{
		Port:     "8080",
		Handlers: server.Http.ServerMux(),
	}

	mgt := models.Model{}

	structMigrate := []models.ModelI{models.SP500{}, models.Queue{}}

	for _, v := range structMigrate {
		mgt.Migrate(v)
	}

	err = srv.Start()

	if err != nil {
		fmt.Println(err.Error())
	}
}
