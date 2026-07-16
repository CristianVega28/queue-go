package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"queue-go/models"
	"queue-go/service"
	"queue-go/utils"
	"time"
)

func Routes() {
	Http.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	Http.Get("/sp500/columns", func(w http.ResponseWriter, r *http.Request) {

		var response any
		cacheKey := "columns:sp500"
		rsv := service.RedisSrvc{}
		ctx := context.Background()

		//TODO Test if it works

		if rsv.Has(ctx, cacheKey) {
			response = rsv.Get(ctx, cacheKey, "nothing")

		} else {
			model := models.Model{}
			columns, err := model.Columns(models.SP500{})

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			response = map[string]any{
				"success": true,
				"data":    columns,
			}

			data, _ := json.Marshal(response)

			err = rsv.Set(ctx, cacheKey, data, time.Hour)

			fmt.Println(err.Error())
		}

		utils.ResponseJson(response, w, r)
	})
}
