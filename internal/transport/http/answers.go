package http

import (
	"context"
	"encoding/json"
	"github.com/spayder/answers-rest-api/internal/answer"
	"log"
	"net/http"
)

type AnswerService interface {
	PostAnswer(context.Context, answer.Answer) (answer.Answer, error)
	GetAnswer(context.Context, string) (answer.Answer, error)
	UpdateAnswer(context.Context, string, answer.Answer) (answer.Answer, error)
	DeleteAnswer(context.Context, string) error
}

func (h *Handler) PostAnswer(w http.ResponseWriter, r *http.Request) {
	var ans answer.Answer
	if err := json.NewDecoder(r.Body).Decode(&ans); err != nil {
		return
	}

	ans, err := h.Service.PostAnswer(r.Context(), ans)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(ans); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateAnswer(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {

}
