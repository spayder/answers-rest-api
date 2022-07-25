package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

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

	h.Router.HandleFunc("/api/answers", h.PostAnswer).Methods("POST")
	h.Router.HandleFunc("/api/answers/{id}", h.GetAnswer).Methods("GET")
	h.Router.HandleFunc("/api/answers/{id}", h.UpdateAnswer).Methods("PUT")
	h.Router.HandleFunc("/api/answers/{id}", h.DeleteAnswer).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutdown gracefully")
	return nil
}
