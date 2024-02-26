package main

import (
	"encoding/json"
	"exam-api/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var categories []model.Category
	db.Find(&categories)

	respondWithJson(w, http.StatusOK, categories)
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var category model.Category
	db.First(&category, id)

	respondWithJson(w, http.StatusOK, category)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	var category model.Category
	json.NewDecoder(r.Body).Decode(&category)

	db.Create(&category)

	respondWithJson(w, http.StatusCreated, category)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var category model.Category
	json.NewDecoder(r.Body).Decode(&category)

	db.First(&category, id)
	db.Save(&category)

	respondWithJson(w, http.StatusOK, category)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*gorm.DB)

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var category model.Category
	db.First(&category, id)
	db.Delete(&category)

	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
