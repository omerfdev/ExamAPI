package main

import (
	"time"
	"C:\Users\omerf\OneDrive\Desktop\exam-api\Handler"
	"github.com/gorilla/mux"
	"C:\Users\omerf\OneDrive\Desktop\exam-api\pkg\Utils\helper.go"
)



func main() {
	r := mux.NewRouter()
	env := utils.GetGoEnv()

	fmt.Println("Exam API running on \"" + env + "\" environment.")
	producer.Execute(env)

	r.HandleFunc("/questions", getQuestions).Methods("GET")
	r.HandleFunc("/questions/{id}", getQuestion).Methods("GET")
	r.HandleFunc("/questions", createQuestion).Methods("POST")
	r.HandleFunc("/questions/{id}", updateQuestion).Methods("PUT")
	r.HandleFunc("/questions/{id}", deleteQuestion).Methods("DELETE")

	r.HandleFunc("/categories", getCategories).Methods("GET")
	r.HandleFunc("/categories/{id}", getCategory).Methods("GET")
	r.HandleFunc("/categories", createCategory).Methods("POST")
	r.HandleFunc("/categories/{id}", updateCategory).Methods("PUT")
	r.HandleFunc("/categories/{id}", deleteCategory).Methods("DELETE")

	r.HandleFunc("/questionoptions", GetQuestionOptions).Methods("GET")
	r.HandleFunc("/questionoptions/{id}", GetQuestionOption).Methods("GET")
	r.HandleFunc("/questionoptions", CreateQuestionOption).Methods("POST")
	r.HandleFunc("/questionoptions/{id}", UpdateQuestionOption).Methods("PUT")
	r.HandleFunc("/questionoptions/{id}", DeleteQuestionOption).Methods("DELETE")

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	r.HandleFunc("/userexams", GetUserExams).Methods("GET")
	r.HandleFunc("/userexams/{id}", GetUserExam).Methods("GET")
	r.HandleFunc("/userexams", CreateUserExam).Methods("POST")
	r.HandleFunc("/userexams/{id}", UpdateUserExam).Methods("PUT")
	r.HandleFunc("/userexams/{id}", DeleteUserExam).Methods("DELETE")

	r.HandleFunc("/exams", GetExams).Methods("GET")
	r.HandleFunc("/exams/{id}", GetExam).Methods("GET")
	r.HandleFunc("/exams", CreateExam).Methods("POST")
	r.HandleFunc("/exams/{id}", UpdateExam).Methods("PUT")
	r.HandleFunc("/exams/{id}", DeleteExam).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
