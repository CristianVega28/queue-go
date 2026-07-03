package server

import "net/http"

type (
	Server struct {
		Port     string
		Handlers *http.ServeMux
	}
)

func (srv *Server) New() *Server {
	return &(Server{})
}

func (srv *Server) Start() error {
	serverHttp := &http.Server{
		Addr:    ":" + srv.Port,
		Handler: srv.Handlers,
	}

	return serverHttp.ListenAndServe()
}
