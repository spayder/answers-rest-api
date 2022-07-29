package http

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/spayder/answers-rest-api/internal/answer"
	"log"
	"net/http"
)

type Response struct {
	Message string
}

type AnswerService interface {
	PostAnswer(context.Context, answer.Answer) (answer.Answer, error)
	GetAnswer(context.Context, string) (answer.Answer, error)
	UpdateAnswer(context.Context, string, answer.Answer) (answer.Answer, error)
	DeleteAnswer(context.Context, string) error
}

type PostAnswerRequest struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}

func ConvertPostAnswerRequestToAnswer(a PostAnswerRequest) answer.Answer {
	return answer.Answer{
		Key:   a.Key,
		Value: a.Value,
	}
}

func (h *Handler) PostAnswer(w http.ResponseWriter, r *http.Request) {
	var ans PostAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&ans); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(ans)
	if err != nil {
		http.Error(w, "not a valid answer", http.StatusBadRequest)
		return
	}

	convertedAnswer := ConvertPostAnswerRequestToAnswer(ans)

	postedAnswer, err := h.Service.PostAnswer(r.Context(), convertedAnswer)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(postedAnswer); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ans, err := h.Service.GetAnswer(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(ans); err != nil {
		log.Print(err)
		panic(err)
	}
}

func (h *Handler) UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var ans answer.Answer
	if err := json.NewDecoder(r.Body).Decode(&ans); err != nil {
		return
	}

	updatedAns, err := h.Service.UpdateAnswer(r.Context(), id, ans)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedAns); err != nil {
		log.Print(err)
		panic(err)
	}
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteAnswer(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		log.Print(err)
		panic(err)
	}
}
