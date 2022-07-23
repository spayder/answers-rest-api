package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type AnswerService interface {
}

type Handler struct {
	Router  *mux.Router
	Service AnswerService
	Server  *http.Server
}

func NewHandler(service AnswerService) *Handler {
	h := &Handler{
		Router:  mux.NewRouter(),
		Service: service,
	}

	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello")
	})
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
