package handler

import (
	"context"
	"encoding/json"
	"exam-api/model"
	"exam-api/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type QuestionOptionHandler struct {
	Repo *repository.QuestionOptionRepository
}

func (h *QuestionOptionHandler) CreateQuestionOption(w http.ResponseWriter, r *http.Request) {
	var option model.QuestionOption
	err := json.NewDecoder(r.Body).Decode(&option)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.Repo.InsertOne(ctx, option)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *QuestionOptionHandler) UpdateQuestionOption(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid option ID", http.StatusBadRequest)
		return
	}

	var option model.QuestionOption
	err = json.NewDecoder(r.Body).Decode(&option)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	updated, err := h.Repo.UpdateOne(ctx, uint(id), option)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !updated {
		http.Error(w, "Option not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *QuestionOptionHandler) DeleteQuestionOption(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid option ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.Repo.DeleteOne(ctx, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *QuestionOptionHandler) GetQuestionOptionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid option ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	option, err := h.Repo.GetByID(ctx, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if option == nil {
		http.Error(w, "Option not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(option)
}

func (h *QuestionOptionHandler) GetAllQuestionOptions(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	options, err := h.Repo.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(options)
}
