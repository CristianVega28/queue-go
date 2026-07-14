package server

import (
	"net/http"
	"queue-go/models"
	"queue-go/utils"
)

func Routes() {
	Http.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	Http.Get("/sp500/columns", func(w http.ResponseWriter, r *http.Request) {

		model := models.Model{}
		columns, err := model.Columns(models.SP500{})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		utils.ResponseJson(map[string]any{
			"success": true,
			"data":    columns,
		}, w, r)
	})
}
