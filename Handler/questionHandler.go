package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"C:\Users\omerf\OneDrive\Desktop\ExamAPI\Types\models.go"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetQuestions(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var questions []Question
	db.Find(&questions)

	respondWithJson(w, http.StatusOK, questions)
}

func GetQuestion(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var question Question
	db.First(&question, id)

	respondWithJson(w, http.StatusOK, question)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var question Question
	json.NewDecoder(r.Body).Decode(&question)

	db.Create(&question)

	respondWithJson(w, http.StatusCreated, question)
}

func UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var question Question
	json.NewDecoder(r.Body).Decode(&question)

	db.First(&question, id)
	db.Save(&question)

	respondWithJson(w, http.StatusOK, question)
}

func DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var question Question
	db.First(&question, id)
	db.Delete(&question)

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}