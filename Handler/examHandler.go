package main

import (
	"encoding/json"
	"net/http"
	"strconv"
"C:\Users\omerf\OneDrive\Desktop\ExamAPI\Types"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateExam(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var exam Exam
	json.NewDecoder(r.Body).Decode(&exam)

	db.Create(&exam)

	respondWithJson(w, http.StatusCreated, exam)
}

func GetExam(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var exam Exam
	db.First(&exam, id)

	respondWithJson(w, http.StatusOK, exam)
}

func GetExams(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var exams []Exam
	db.Find(&exams)

	respondWithJson(w, http.StatusOK, exams)
}

func UpdateExam(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var exam Exam
	json.NewDecoder(r.Body).Decode(&exam)

	db.First(&exam, id)
	db.Save(&exam)

	respondWithJson(w, http.StatusOK, exam)
}

func DeleteExam(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var exam Exam
	db.First(&exam, id)
	db.Delete(&exam)

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
