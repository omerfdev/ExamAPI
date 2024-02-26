package handler

import (
	"encoding/json"
	"exam-api/model"
	"exam-api/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ExamHandler struct {
	Repo *repository.ExamRepository
}

func NewExamHandler(repo *repository.ExamRepository) *ExamHandler {
	return &ExamHandler{Repo: repo}
}

func (h *ExamHandler) CreateExam(w http.ResponseWriter, r *http.Request) {
	var exam model.Exam
	err := json.NewDecoder(r.Body).Decode(&exam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Repo.InsertOne(r.Context(), exam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ExamHandler) UpdateExam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid exam ID", http.StatusBadRequest)
		return
	}

	var exam model.Exam
	err = json.NewDecoder(r.Body).Decode(&exam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.Repo.UpdateOne(r.Context(), uint(id), exam)
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

func (h *ExamHandler) DeleteExam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid exam ID", http.StatusBadRequest)
		return
	}

	err = h.Repo.DeleteOne(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ExamHandler) GetAllExams(w http.ResponseWriter, r *http.Request) {
	exams, err := h.Repo.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(exams)
}

func (h *ExamHandler) GetExamByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid exam ID", http.StatusBadRequest)
		return
	}

	exam, err := h.Repo.GetByID(r.Context(), uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exam == nil {
		http.Error(w, "Exam not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(exam)
}
