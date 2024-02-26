package repository

import (
	"context"
	"errors"
	"exam-api/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	DB *mongo.Collection
}

func (r *CategoryRepository) InsertOne(ctx context.Context, insertModel model.Category) error {
	_, err := r.DB.InsertOne(ctx, insertModel)
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}

	return nil
}

func (r *CategoryRepository) UpdateOne(ctx context.Context, id uint, updateModel model.Category) (bool, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateModel}

	result, err := r.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}

	if result.ModifiedCount == 0 {
		return false, errors.New("No document found to update")
	}

	return true, nil
}

func (r *CategoryRepository) DeleteOne(ctx context.Context, id uint) error {
	filter := bson.M{"_id": id}

	result, err := r.DB.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}

	if result.DeletedCount == 0 {
		return errors.New("No document found to delete")
	}

	return nil
}

func (r *CategoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var categories []model.Category

	cursor, err := r.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &categories); err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}

	return categories, nil
}

func (r *CategoryRepository) GetByID(ctx context.Context, id uint) (*model.Category, error) {
	var category model.Category

	filter := bson.M{"_id": id}
	err := r.DB.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}

	return &category, nil
}
func (r *CategoryRepository) GetFirstOrDefault(ctx context.Context, filter bson.M) (*model.Category, error) {
	var category model.Category
	err := r.DB.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		// MongoDB'de belirtilen kriterlere uyan belge bulunamadığında nil döndürülebilir.
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}

	return &category, nil
}

type QuestionRepository struct {
	DB *mongo.Collection
}

func (r *QuestionRepository) InsertOne(ctx context.Context, insertModel model.Question) error {
	_, err := r.DB.InsertOne(ctx, insertModel)
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}
	return nil
}

func (r *QuestionRepository) DeleteOne(ctx context.Context, id uint) error {
	filter := bson.M{"_id": id}
	result, err := r.DB.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}
	if result.DeletedCount == 0 {
		return errors.New("No document found to delete")
	}
	return nil
}

func (r *QuestionRepository) UpdateOne(ctx context.Context, id uint, updateModel model.Question) (bool, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateModel}

	result, err := r.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}
	if result.ModifiedCount == 0 {
		return false, errors.New("No document found to update")
	}
	return true, nil
}

func (r *QuestionRepository) GetAll(ctx context.Context) ([]model.Question, error) {
	var questions []model.Question

	cursor, err := r.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &questions); err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}

	return questions, nil
}

func (r *QuestionRepository) GetByID(ctx context.Context, id uint) (*model.Question, error) {
	var question model.Question

	filter := bson.M{"_id": id}
	err := r.DB.FindOne(ctx, filter).Decode(&question)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}

	return &question, nil
}
func (r *QuestionRepository) GetFirstOrDefault(ctx context.Context, filter bson.M) (*model.Question, error) {
	var question model.Question
	err := r.DB.FindOne(ctx, filter).Decode(&question)
	if err != nil {
		// MongoDB'de belirtilen kriterlere uyan belge bulunamadığında nil döndürülebilir.
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}

	return &question, nil
}

type QuestionOptionRepository struct {
	DB *mongo.Collection
}

func (r *QuestionOptionRepository) InsertOne(ctx context.Context, insertModel model.QuestionOption) error {
	_, err := r.DB.InsertOne(ctx, insertModel)
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}

	return nil
}
func (r *QuestionOptionRepository) UpdateOne(ctx context.Context, id uint, updateModel model.QuestionOption) (bool, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateModel}

	result, err := r.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}

	if result.ModifiedCount == 0 {
		return false, errors.New("No document found to update")
	}

	return true, nil
}
func (r *QuestionOptionRepository) DeleteOne(ctx context.Context, id uint) error {
	filter := bson.M{"_id": id}

	result, err := r.DB.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}

	if result.DeletedCount == 0 {
		return errors.New("No document found to delete")
	}

	return nil
}
func (r *QuestionOptionRepository) GetAll(ctx context.Context) ([]model.QuestionOption, error) {
	var questionOptions []model.QuestionOption

	cursor, err := r.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &questionOptions); err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}

	return questionOptions, nil
}
func (r *QuestionOptionRepository) GetByID(ctx context.Context, id uint) (*model.QuestionOption, error) {
	var questionOption model.QuestionOption

	filter := bson.M{"_id": id}
	err := r.DB.FindOne(ctx, filter).Decode(&questionOption)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}

	return &questionOption, nil
}
func (r *QuestionOptionRepository) GetFirstOrDefault(ctx context.Context, filter bson.M) (*model.QuestionOption, error) {
	var questionOption model.QuestionOption
	err := r.DB.FindOne(ctx, filter).Decode(&questionOption)
	if err != nil {
		// MongoDB'de belirtilen kriterlere uyan belge bulunamadığında nil döndürülebilir.
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}

	return &questionOption, nil
}

type UserRepository struct {
	DB *mongo.Collection
}

func (r *UserRepository) InsertOne(ctx context.Context, insertModel model.User) error {
	_, err := r.DB.InsertOne(ctx, insertModel)
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}

	return nil
}
func (r *UserRepository) UpdateOne(ctx context.Context, id uint, updateModel model.User) (bool, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateModel}

	result, err := r.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}

	if result.ModifiedCount == 0 {
		return false, errors.New("No document found to update")
	}

	return true, nil
}
func (r *UserRepository) DeleteOne(ctx context.Context, id uint) error {
	filter := bson.M{"_id": id}

	result, err := r.DB.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}

	if result.DeletedCount == 0 {
		return errors.New("No document found to delete")
	}

	return nil
}
func (r *UserRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User

	cursor, err := r.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &users); err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}

	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User

	filter := bson.M{"_id": id}
	err := r.DB.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}

	return &user, nil
}
func (r *UserRepository) GetFirstOrDefault(ctx context.Context, filter bson.M) (*model.User, error) {
	var user model.User

	err := r.DB.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Belirtilen filtre ile eşleşen belge bulunamadı
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}

	return &user, nil
}

type UserExamRepository struct {
	DB *mongo.Collection
}

func (r *UserExamRepository) InsertOne(ctx context.Context, insertModel model.UserExam) error {
	_, err := r.DB.InsertOne(ctx, insertModel)
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}

	return nil
}
func (r *UserExamRepository) UpdateOne(ctx context.Context, id uint, updateModel model.UserExam) (bool, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateModel}

	result, err := r.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}

	if result.ModifiedCount == 0 {
		return false, errors.New("No document found to update")
	}

	return true, nil
}
func (r *UserExamRepository) DeleteOne(ctx context.Context, id uint) error {
	filter := bson.M{"_id": id}

	result, err := r.DB.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}

	if result.DeletedCount == 0 {
		return errors.New("No document found to delete")
	}

	return nil
}
func (r *UserExamRepository) GetAll(ctx context.Context) ([]model.UserExam, error) {
	var userExams []model.UserExam

	cursor, err := r.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &userExams); err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}

	return userExams, nil
}

func (r *UserExamRepository) GetByID(ctx context.Context, id uint) (*model.UserExam, error) {
	var userExam model.UserExam

	filter := bson.M{"_id": id}
	err := r.DB.FindOne(ctx, filter).Decode(&userExam)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}

	return &userExam, nil
}
func (r *UserExamRepository) GetFirstOrDefault(ctx context.Context, filter bson.M) (*model.UserExam, error) {
	var userExam model.UserExam

	err := r.DB.FindOne(ctx, filter).Decode(&userExam)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Belirtilen filtre ile eşleşen belge bulunamadı
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}

	return &userExam, nil
}

type ExamRepository struct {
	DB *mongo.Collection
}

func (r *ExamRepository) InsertOne(ctx context.Context, insertModel model.Exam) error {
	_, err := r.DB.InsertOne(ctx, insertModel)
	if err != nil {
		return errors.New(fmt.Sprintf("InsertOne failed: %v", err))
	}

	return nil
}
func (r *ExamRepository) UpdateOne(ctx context.Context, id uint, updateModel model.Exam) (bool, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateModel}

	result, err := r.DB.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, errors.New(fmt.Sprintf("UpdateOne failed: %v", err))
	}

	if result.ModifiedCount == 0 {
		return false, errors.New("No document found to update")
	}

	return true, nil
}

func (r *ExamRepository) DeleteOne(ctx context.Context, id uint) error {
	filter := bson.M{"_id": id}

	result, err := r.DB.DeleteOne(ctx, filter)
	if err != nil {
		return errors.New(fmt.Sprintf("DeleteOne failed: %v", err))
	}

	if result.DeletedCount == 0 {
		return errors.New("No document found to delete")
	}

	return nil
}
func (r *ExamRepository) GetAll(ctx context.Context) ([]model.Exam, error) {
	var exams []model.Exam

	cursor, err := r.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &exams); err != nil {
		return nil, errors.New(fmt.Sprintf("GetAll failed: %v", err))
	}

	return exams, nil
}
func (r *ExamRepository) GetByID(ctx context.Context, id uint) (*model.Exam, error) {
	var exam model.Exam

	filter := bson.M{"_id": id}
	err := r.DB.FindOne(ctx, filter).Decode(&exam)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("GetByID failed: %v", err))
	}

	return &exam, nil
}
func (r *ExamRepository) GetFirstOrDefault(ctx context.Context, filter bson.M) (*model.Exam, error) {
	var exam model.Exam

	err := r.DB.FindOne(ctx, filter).Decode(&exam)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Belirtilen filtre ile eşleşen belge bulunamadı
		}
		return nil, errors.New(fmt.Sprintf("GetFirstOrDefault failed: %v", err))
	}

	return &exam, nil
}
