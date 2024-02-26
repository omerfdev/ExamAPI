package repository

import (
	"context"
	"errors"
	"exam-api/model"
	"fmt"

	"gorm.io/gorm"
)

type CategoryPostgreRepository struct {
	DB *gorm.DB
}

func (r *CategoryPostgreRepository) UpdateOne(ctx context.Context, id uint, updateModel model.Category) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&model.Category{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type QuestionPostgreRepository struct {
	DB *gorm.DB
}

func (r *QuestionPostgreRepository) UpdateOne(ctx context.Context, id uint, updateModel model.Question) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&model.Question{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type QuestionOptionPostgreRepository struct {
	DB *gorm.DB
}

func (r *QuestionOptionPostgreRepository) UpdateOne(ctx context.Context, id uint, updateModel model.QuestionOption) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&model.QuestionOption{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type UserPostgreRepository struct {
	DB *gorm.DB
}

func (r *UserPostgreRepository) UpdateOne(ctx context.Context, id uint, updateModel model.User) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&model.User{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type UserExamPostgreRepository struct {
	DB *gorm.DB
}

func (r *UserExamPostgreRepository) UpdateOne(ctx context.Context, id uint, updateModel model.UserExam) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&model.UserExam{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}

type ExamPostgreRepository struct {
	DB *gorm.DB
}

func (r *ExamPostgreRepository) UpdateOne(ctx context.Context, id uint, updateModel model.Exam) (bool, error) {
	err := r.DB.WithContext(ctx).Model(&model.Exam{}).Where("id = ?", id).Updates(updateModel).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	return true, nil
}
