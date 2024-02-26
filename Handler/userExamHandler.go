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

type UserExamHandler struct {
	Repo *repository.UserExamRepository
}

func (h *UserExamHandler) CreateUserExam(w http.ResponseWriter, r *http.Request) {
	var userExam model.UserExam
	err := json.NewDecoder(r.Body).Decode(&userExam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	err = h.Repo.InsertOne(ctx, userExam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserExamHandler) UpdateUserExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user exam ID", http.StatusBadRequest)
		return
	}

	var userExam model.UserExam
	err = json.NewDecoder(r.Body).Decode(&userExam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	updated, err := h.Repo.UpdateOne(ctx, uint(id), userExam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !updated {
		http.Error(w, "User exam not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UserExamHandler) DeleteUserExam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user exam ID", http.StatusBadRequest)
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

func (h *UserExamHandler) GetUserExamByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user exam ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	userExam, err := h.Repo.GetByID(ctx, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if userExam == nil {
		http.Error(w, "User exam not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(userExam)
}

func (h *UserExamHandler) GetAllUserExams(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userExams, err := h.Repo.GetAll(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userExams)
}
