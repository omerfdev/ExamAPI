package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"exam-api/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateQuestionOption(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var questionOption model.QuestionOption
	json.NewDecoder(r.Body).Decode(&questionOption)

	db.Create(&questionOption)

	respondWithJson(w, http.StatusCreated, questionOption)
}

func GetQuestionOption(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var questionOption model.QuestionOption
	db.First(&questionOption, id)

	respondWithJson(w, http.StatusOK, questionOption)
}

func GetQuestionOptions(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var questionOptions []model.QuestionOption
	db.Find(&questionOptions)

	respondWithJson(w, http.StatusOK, questionOptions)
}

func UpdateQuestionOption(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var questionOption model.QuestionOption
	json.NewDecoder(r.Body).Decode(&questionOption)

	db.First(&questionOption, id)
	db.Save(&questionOption)

	respondWithJson(w, http.StatusOK, questionOption)
}

func DeleteQuestionOption(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var questionOption model.QuestionOption
	db.First(&questionOption, id)
	db.Delete(&questionOption)

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
 