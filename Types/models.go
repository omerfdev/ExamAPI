package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BaseEntity struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type Category struct {
	BaseEntity
	Name string `gorm:"type:varchar(100);unique_index" json:"name"`
}

type Question struct {
	BaseEntity
	QuestionText string `json:"question_text"`
	QuestionType int    `json:"question_type"`
	Category     Category
	CategoryID   uint `json:"category_id"`
}

type QuestionOption struct {
	BaseEntity
	OptionText string `json:"option_text"`
	IsCorrect  bool   `json:"is_correct"`
	Question   Question
	QuestionID  uint `json:"question_id"`
}

type User struct {
	BaseEntity
	Name     string `gorm:"type:varchar(100);unique_index" json:"name"`
	Surname  string `gorm:"type:varchar(100);unique_index" json:"surname"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Title    string `json:"title"`
	UserExams []UserExam `json:"user_exams,omitempty"`
}

type UserExam struct {
	BaseEntity
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Score     int       `json:"score,omitempty"`
	User      User      `json:"user,omitempty"`
	UserID    uint      `json:"user_id"`
	Exam      Exam      `json:"exam,omitempty"`
	ExamID    uint      `json:"exam_id"`
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

