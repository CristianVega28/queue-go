package utils

import (
	"encoding/json"
	"net/http"
)

// ResponseJson crea una respuesta parseada y con los headers de json
func ResponseJson(data any, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
