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

func (r *CategoryPostgreRepository) InsertOne(ctx context.Context, insertModel model.Category) error {
	err := r.DB.WithContext(ctx).Create(&insertModel).Error
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}
	return nil
}

func (r *CategoryPostgreRepository) DeleteOne(ctx context.Context, id uint) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.Category{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}
	return nil
}

func (r *CategoryPostgreRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category
	err := r.DB.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	return categories, nil
}

func (r *CategoryPostgreRepository) GetByID(ctx context.Context, id uint) (*model.Category, error) {
	var category model.Category
	err := r.DB.WithContext(ctx).First(&category, id).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}
	return &category, nil
}

func (r *CategoryPostgreRepository) GetFirstOrDefault(ctx context.Context, filter interface{}) (*model.Category, error) {
	var category model.Category
	err := r.DB.WithContext(ctx).Where(filter).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}
	return &category, nil
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

func (r *QuestionPostgreRepository) InsertOne(ctx context.Context, insertModel model.Question) error {
	err := r.DB.WithContext(ctx).Create(&insertModel).Error
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}
	return nil
}

func (r *QuestionPostgreRepository) DeleteOne(ctx context.Context, id uint) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.Question{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}
	return nil
}

func (r *QuestionPostgreRepository) GetAll(ctx context.Context) ([]model.Question, error) {
	var questions []model.Question
	err := r.DB.WithContext(ctx).Find(&questions).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	return questions, nil
}

func (r *QuestionPostgreRepository) GetByID(ctx context.Context, id uint) (*model.Question, error) {
	var question model.Question
	err := r.DB.WithContext(ctx).First(&question, id).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}
	return &question, nil
}

func (r *QuestionPostgreRepository) GetFirstOrDefault(ctx context.Context, filter interface{}) (*model.Question, error) {
	var question model.Question
	err := r.DB.WithContext(ctx).Where(filter).First(&question).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}
	return &question, nil
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

func (r *QuestionOptionPostgreRepository) InsertOne(ctx context.Context, insertModel model.QuestionOption) error {
	err := r.DB.WithContext(ctx).Create(&insertModel).Error
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}
	return nil
}

func (r *QuestionOptionPostgreRepository) DeleteOne(ctx context.Context, id uint) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.QuestionOption{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}
	return nil
}

func (r *QuestionOptionPostgreRepository) GetAll(ctx context.Context) ([]model.QuestionOption, error) {
	var questionOptions []model.QuestionOption
	err := r.DB.WithContext(ctx).Find(&questionOptions).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	return questionOptions, nil
}

func (r *QuestionOptionPostgreRepository) GetByID(ctx context.Context, id uint) (*model.QuestionOption, error) {
	var questionOption model.QuestionOption
	err := r.DB.WithContext(ctx).First(&questionOption, id).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}
	return &questionOption, nil
}

func (r *QuestionOptionPostgreRepository) GetFirstOrDefault(ctx context.Context, filter interface{}) (*model.QuestionOption, error) {
	var questionOption model.QuestionOption
	err := r.DB.WithContext(ctx).Where(filter).First(&questionOption).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}
	return &questionOption, nil
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

func (r *UserPostgreRepository) InsertOne(ctx context.Context, insertModel model.User) error {
	err := r.DB.WithContext(ctx).Create(&insertModel).Error
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}
	return nil
}

func (r *UserPostgreRepository) DeleteOne(ctx context.Context, id uint) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.User{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}
	return nil
}

func (r *UserPostgreRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := r.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	return users, nil
}

func (r *UserPostgreRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}
	return &user, nil
}

func (r *UserPostgreRepository) GetFirstOrDefault(ctx context.Context, filter interface{}) (*model.User, error) {
	var user model.User
	err := r.DB.WithContext(ctx).Where(filter).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}
	return &user, nil
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
func (r *UserExamPostgreRepository) InsertOne(ctx context.Context, insertModel model.UserExam) error {
	err := r.DB.WithContext(ctx).Create(&insertModel).Error
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}
	return nil
}

func (r *UserExamPostgreRepository) DeleteOne(ctx context.Context, id uint) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.UserExam{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}
	return nil
}

func (r *UserExamPostgreRepository) GetAll(ctx context.Context) ([]model.UserExam, error) {
	var userExams []model.UserExam
	err := r.DB.WithContext(ctx).Find(&userExams).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	return userExams, nil
}

func (r *UserExamPostgreRepository) GetByID(ctx context.Context, id uint) (*model.UserExam, error) {
	var userExam model.UserExam
	err := r.DB.WithContext(ctx).First(&userExam, id).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}
	return &userExam, nil
}

func (r *UserExamPostgreRepository) GetFirstOrDefault(ctx context.Context, filter interface{}) (*model.UserExam, error) {
	var userExam model.UserExam
	err := r.DB.WithContext(ctx).Where(filter).First(&userExam).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}
	return &userExam, nil
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

func (r *ExamPostgreRepository) InsertOne(ctx context.Context, insertModel model.Exam) error {
	err := r.DB.WithContext(ctx).Create(&insertModel).Error
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}
	return nil
}

func (r *ExamPostgreRepository) DeleteOne(ctx context.Context, id uint) error {
	err := r.DB.WithContext(ctx).Where("id = ?", id).Delete(&model.Exam{}).Error
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}
	return nil
}

func (r *ExamPostgreRepository) GetAll(ctx context.Context) ([]model.Exam, error) {
	var exams []model.Exam
	err := r.DB.WithContext(ctx).Find(&exams).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	return exams, nil
}

func (r *ExamPostgreRepository) GetByID(ctx context.Context, id uint) (*model.Exam, error) {
	var exam model.Exam
	err := r.DB.WithContext(ctx).First(&exam, id).Error
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}
	return &exam, nil
}

func (r *ExamPostgreRepository) GetFirstOrDefault(ctx context.Context, filter interface{}) (*model.Exam, error) {
	var exam model.Exam
	err := r.DB.WithContext(ctx).Where(filter).First(&exam).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}
	return &exam, nil
}
