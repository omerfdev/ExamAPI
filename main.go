package main

import (
	"time"
"C:\Users\omerf\OneDrive\Desktop\ExamAPI\Handler"
	"github.com/gorilla/mux"
)

type BaseEntity struct {
	ID         int       `json:"id"`
	CreateDate time.Time `json:"create_date"`
	UpdateDate time.Time `json:"update_date,omitempty"`
	DeleteDate time.Time `json:"delete_date,omitempty"`
	IsActive   bool      `json:"is_active"`
}

type Category struct {
	BaseEntity
	CategoryName string `json:"category_name"`
	Exams        []Exam `json:"exams,omitempty"`
}

type Question struct {
	BaseEntity
	QuestionType    QuestionType     `json:"question_type"`
	QuestionText    string           `json:"question_text"`
	AnswerText      string           `json:"answer_text,omitempty"`
	QuestionOptions []QuestionOption `json:"question_options,omitempty"`
	Exam            Exam             `json:"exam,omitempty"`
}

type QuestionOption struct {
	BaseEntity
	OptionText   string   `json:"option_text"`
	IsTrueOption bool     `json:"is_true_option"`
	Question     Question `json:"question,omitempty"`
}

type User struct {
	BaseEntity
	Name      string     `json:"name"`
	Surname   string     `json:"surname"`
	Email     string     `gorm:"type:varchar(100);unique_index" json:"email"`
	Title     string     `json:"title"`
	UserExams []UserExam `json:"user_exams,omitempty"`
	Exam      Exam       `json:"exam,omitempty"`
}

type UserExam struct {
	BaseEntity
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Score     int       `json:"score,omitempty"`
	User      User      `json:"user,omitempty"`
	Exam      Exam      `json:"exam,omitempty"`
}

type Exam struct {
	BaseEntity
	ExamName         string     `json:"exam_name"`
	Duration         int        `json:"duration"`
	Description      string     `json:"description"`
	AchievementScore int        `json:"achievement_score,omitempty"`
	Category         Category   `json:"category,omitempty"`
	Questions        []Question `json:"questions,omitempty"`
	Users            []User     `json:"users,omitempty"`
	UserExams        []UserExam `json:"user_exams,omitempty"`
}

type QuestionType int

const (
	MultipleChoice QuestionType = iota
	OpenEnded
)

func main() {
	r := mux.NewRouter()

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
