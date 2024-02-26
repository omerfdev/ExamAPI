package main

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func getDb() *gorm.DB {
	db, err := gorm.Open("sqlserver", "Data Source=ALMALI\\OMERFDEV;Initial Catalog=ExamDB;User ID=sa;pwd=Omer34;")
	if err != nil {
		panic("Failed to connect to the database!")
	}

	db.AutoMigrate(&BaseEntity{})
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Question{})
	db.AutoMigrate(&QuestionOption{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&UserExam{})
	db.AutoMigrate(&Exam{})

	return db
}