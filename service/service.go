package service

import (
	
	"github.com/crudapigin/deyki/v2/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username 	string	`json:"username"`
	Email		string	`json:"email"`
}


type UserResponseModel struct {
	Username	string	`json:"username"`
	Email		string	`json:"email"`
}


func ConnectDB() (*gorm.DB, *util.ErrorMessage) {

	dsn := "host=localhost user=postgres password=1234 dbname=crudapigin port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, util.ErrorMessage{}.FailedToOpenDB()
	}

	db.AutoMigrate(&User{})

	return db, nil
}


func AddUser(u *User) (*UserResponseModel, *util.ErrorMessage) {

	db, err := ConnectDB()
	if err != nil {
		return nil, util.ErrorMessage{}.FailedToOpenDB()
	}

	db.Create(&u)

	return &UserResponseModel{u.Username, u.Email}, nil
}


func GetUserById(userID int) (*UserResponseModel, *util.ErrorMessage) {

	db, err := ConnectDB()
	if err != nil {
		return nil, util.ErrorMessage{}.FailedToOpenDB()
	}

	var user User

	errorMessage := db.Where("id = ?", userID).First(&user).Error
	if errorMessage != nil {
		return nil, util.ErrorMessage{}.UserNotFound()
	}

	return &UserResponseModel{user.Username, user.Email}, nil
}


func GetUsers() (*[]UserResponseModel, *util.ErrorMessage) {

	db, err := ConnectDB()
	if err != nil {
		return nil, util.ErrorMessage{}.FailedToOpenDB()
	}

	var users []User
	var userResponseModels []UserResponseModel

	db.Find(&users)

	for _, user := range users {
		userResponseModels = append(userResponseModels, UserResponseModel{user.Username, user.Email})
	}

	return &userResponseModels, nil
}


func DeleteUserById(userID int) *util.ErrorMessage {

	db, err := ConnectDB()
	if err != nil {
		return util.ErrorMessage{}.FailedToOpenDB()
	}

	var user User

	exist := db.First(&user, userID).Error

	if exist != nil {
		return util.ErrorMessage{}.UserNotFound()
	} else {
		db.Delete(&user)
	}

	return nil
}