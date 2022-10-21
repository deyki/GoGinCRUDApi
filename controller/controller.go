package controller

import (
	"net/http"
	"strconv"

	"github.com/crudapigin/deyki/v2/service"
	"github.com/gin-gonic/gin"
)


func addUserHandler(c *gin.Context) {

	var user service.User

	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	userResponseModel, errorMessage := service.AddUser(&user)
	if errorMessage != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"ErrorMessage": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"User": userResponseModel})
}


func getUserByIdHandler(c *gin.Context) {

	pathVariable := c.Param("userID")

	userID, err := strconv.Atoi(pathVariable)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": "Failed to convert pathVariable to userID(int)"})
		return
	}

	userResponseModel, errorMessage := service.GetUserById(userID)
	if errorMessage != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"ErrorMessage": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"User": userResponseModel})
}


func getUsersHandler(c *gin.Context) {

	usersResponseModels, errorMessage := service.GetUsers()
	if errorMessage != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"ErrorMessage": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, usersResponseModels)
}


func deleteUserByIdHanlder(c *gin.Context) {

	pathVariable := c.Param("userID")

	userID, err := strconv.Atoi(pathVariable)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": "Failed to convert pathVariable to userID(int)"})
		return
	}

	errorMessage := service.DeleteUserById(userID)
	if errorMessage != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"ErrorMessage": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Message": "User with id " + pathVariable + " deleted!"})
}


func updateUserByIdHandler(c *gin.Context) {

	pathVariable := c.Param("userID")

	var user service.User

	err := c.BindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Bad request"})
		return
	}

	userID, err := strconv.Atoi(pathVariable)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": "Failed to convert pathVariable to userID(int)"})
		return
	}

	userResponseModel, errorMessage := service.UpdateUserById(userID, &user)
	if errorMessage != nil {
		c.IndentedJSON(errorMessage.HttpStatus, gin.H{"Error": errorMessage})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"Updated user": userResponseModel})
}

func GinRouter() {

	router := gin.Default()
	router.POST("/addUser", addUserHandler)
	router.GET("/getUser/:userID", getUserByIdHandler)
	router.GET("/users", getUsersHandler)
	router.DELETE("/deleteUser/:userID", deleteUserByIdHanlder)
	router.PUT("/updateUser/:userID", updateUserByIdHandler)
	router.Run()
}

