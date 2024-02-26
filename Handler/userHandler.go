package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"exam-api/model"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var users []model.User
	db.Find(&users)

	respondWithJson(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user model.User
	db.First(&user, id)

	respondWithJson(w, http.StatusOK, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	db.Create(&user)

	respondWithJson(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user model.User
	json.NewDecoder(r.Body).Decode(&user)

	db.First(&user, id)
	db.Save(&user)

	respondWithJson(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user model.User
	db.First(&user, id)
	db.Delete(&user)

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}