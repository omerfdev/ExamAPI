package handler

import (
	"encoding/json"
	"exam-api/model"
	"exam-api/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type QuestionHandler struct {
	Repo *repository.QuestionRepository
}

func NewQuestionHandler(repo *repository.QuestionRepository) *QuestionHandler {
	return &QuestionHandler{Repo: repo}
}

func (h *QuestionHandler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var question model.Question
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Repo.InsertOne(r.Context(), question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *QuestionHandler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	var question model.Question
	err = json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.Repo.UpdateOne(r.Context(), uint(id), question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !updated {
		http.Error(w, "No document found to update", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.DeleteOne(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *QuestionHandler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.Repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(questions)
}

func (h *QuestionHandler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid question ID", http.StatusBadRequest)
		return
	}

	question, err := h.Repo.GetByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if question == nil {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(question)
}
