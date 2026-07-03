package server

import "net/http"

type (
	_Http struct {
		mux *http.ServeMux
	}
)

var Http _Http

func (h *_Http) New() *_Http {
	Http = _Http{
		mux: http.NewServeMux(),
	}
	return &Http
}

func (h *_Http) Get(path string, f http.HandlerFunc) {
	h.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		f(w, r)
	})
}

func (h *_Http) Post(path string, f http.HandlerFunc) {
	h.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		f(w, r)
	})
}

func (h *_Http) ServerMux() *http.ServeMux {
	return h.mux
}
