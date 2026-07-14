package main

import (
	"log/slog"
	"queue-go/logger"
	"queue-go/models"
	"queue-go/server"
)

func main() {
	logger.Setup()

	err := models.Setup()
	if err != nil {
		slog.Error("fallo al inicializar la base de datos", "err", err)
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
		if err := mgt.Migrate(v); err != nil {
			slog.Error("fallo en la migración", "err", err)
		}
	}

	slog.Info("servidor iniciado", "url", "http://localhost:"+srv.Port)

	if err := srv.Start(); err != nil {
		slog.Error("el servidor se detuvo", "err", err)
	}
}
