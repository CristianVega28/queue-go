// Package logger configura el logger global de la aplicación usando log/slog
// con salida coloreada en consola vía tint.
package logger

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// Setup inicializa el logger por defecto de slog con formato coloreado.
// Después de llamarlo, se puede usar slog.Info, slog.Error, etc. desde
// cualquier paquete.
func Setup() {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))

	slog.SetDefault(logger)
}
