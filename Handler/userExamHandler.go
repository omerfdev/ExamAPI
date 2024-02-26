package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"C:\Users\omerf\OneDrive\Desktop\ExamAPI\Types"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


func CreateUserExam(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var userExam UserExam
	json.NewDecoder(r.Body).Decode(&userExam)

	db.Create(&userExam)

	respondWithJson(w, http.StatusCreated, userExam)
}

func GetUserExam(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var userExam UserExam
	db.First(&userExam, id)

	respondWithJson(w, http.StatusOK, userExam)
}

func GetUserExams(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var userExams []UserExam
	db.Find(&userExams)

	respondWithJson(w, http.StatusOK, userExams)
}

func RemoveUserExam(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var userExam UserExam
	db.First(&userExam, id)
	db.Delete(&userExam)

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func GetExamsByStudentId(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	studentId, _ := strconv.Atoi(params["studentId"])

	var userExams []UserExam
	db.Find(&userExams, "user_id = ?", studentId)

	respondWithJson(w, http.StatusOK, userExams)
}

func UpdateExamUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var userExam UserExam
	json.NewDecoder(r.Body).Decode(&userExam)

	db.First(&userExam, id)
	db.Save(&userExam)

	respondWithJson(w, http.StatusOK, userExam)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
