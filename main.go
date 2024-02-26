package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"exam-api/configs"
	"exam-api/repository"
	"exam-api/handler"
)

func main() {
	// MongoDB yapılandırması
	config := configs.Configs["local-qa"]
	clientOptions := options.Client().ApplyURI(config.URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(config.Database)
	categoryCollection := db.Collection(config.Collection)

	categoryRepo := &repository.CategoryRepository{DB: categoryCollection}
	categoryHandler := handler.NewCategoryHandler(categoryRepo)

	r := mux.NewRouter()
	r.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", categoryHandler.GetCategoryByID).Methods("GET")

	http.Handle("/", r)

	// Sunucuyu başlat
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
