package main

import (
	"C:\Users\omerf\OneDrive\Desktop\ExamAPI\Types\models.go"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func (r *CategoryRepository) UpdateOne(ctx context.Context, id uint, updateModel Category) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&Category{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type QuestionRepository struct {
	DB *gorm.DB
}

func (r *QuestionRepository) UpdateOne(ctx context.Context, id uint, updateModel Question) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&Question{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type QuestionOptionRepository struct {
	DB *gorm.DB
}

func (r *QuestionOptionRepository) UpdateOne(ctx context.Context, id uint, updateModel QuestionOption) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&QuestionOption{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) UpdateOne(ctx context.Context, id uint, updateModel User) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type UserExamRepository struct {
	DB *gorm.DB
}

func (r *UserExamRepository) UpdateOne(ctx context.Context, id uint, updateModel UserExam) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&UserExam{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type ExamRepository struct {
	DB *gorm.DB
}

func (r *ExamRepository) UpdateOne(ctx context.Context, id uint, updateModel Exam) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&Exam{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}