package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateCategoryRequest struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateQuestionRequest struct {
	QuestionText string `json:"question_text" binding:"required"`
	QuestionType int    `json:"question_type" binding:"required"`
	CategoryID   uint   `json:"category_id" binding:"required"`
}

type UpdateQuestionRequest struct {
	ID           uint   `json:"id" binding:"required"`
	QuestionText string `json:"question_text" binding:"required"`
	QuestionType int    `json:"question_type" binding:"required"`
	CategoryID   uint   `json:"category_id" binding:"required"`
}

type CreateQuestionOptionRequest struct {
	OptionText string `json:"option_text" binding:"required"`
	IsCorrect  bool   `json:"is_correct" binding:"required"`
	QuestionID  uint   `json:"question_id" binding:"required"`
}

type UpdateQuestionOptionRequest struct {
	ID        uint   `json:"id" binding:"required"`
	OptionText string `json:"option_text" binding:"required"`
	IsCorrect bool   `json:"is_correct" binding:"required"`
	QuestionID uint   `json:"question_id" binding:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

type UpdateUserRequest struct {
	ID       uint   `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Title    string `json:"title" binding:"required"`
}

type CreateUserExamRequest struct {
	StartDate time.Time `json:"start_date" binding:"required"`
	UserID    uint      `json:"user_id" binding:"required"`
	ExamID    uint      `json:"exam_id" binding:"required"`
}

type UpdateUserExamRequest struct {
	ID       uint   `json:"id" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required"`
	Score     int       `json:"score" binding:"required"`
	UserID    uint      `json:"user_id" binding:"required"`
	ExamID    uint      `json:"exam_id" binding:"required"`
}

type CreateExamRequest struct {
	ExamName         string     `json:"exam_name" binding:"required"`
	Duration         int        `json:"duration" binding:"required"`
	Description      string     `json:"description" binding:"required"`
	AchievementScore int        `json:"achievement_score" binding:"required"`
	CategoryID       uint       `json:"category_id" binding:"required"`
}

type UpdateExamRequest struct {
	ID           uint   `json:"id" binding:"required"`
	ExamName         string     `json:"exam_name" binding:"required"`
	Duration         int        `json:"duration" binding:"required"`
	Description      string     `json:"description" binding:"required"`
	AchievementScore int        `json:"achievement_score" binding:"required"`
	CategoryID       uint       `json:"category_id" binding:"required"`
}