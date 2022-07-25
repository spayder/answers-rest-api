package http

import (
	"context"
	"github.com/spayder/answers-rest-api/internal/answer"
	"net/http"
)

type AnswerService interface {
	GetAnswers(context.Context) ([]answer.Answer, error)
	PostAnswer(context.Context, answer.Answer) (answer.Answer, error)
	GetAnswer(context.Context, string) (answer.Answer, error)
	UpdateAnswer(context.Context, string, answer.Answer) (answer.Answer, error)
	DeleteAnswer(context.Context, answer.Answer) error
}

func (h *Handler) GetAnswers(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) PostAnswer(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateAnswer(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {

}
