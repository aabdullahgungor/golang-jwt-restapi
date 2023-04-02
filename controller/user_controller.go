package controller

import (
	"net/http"

	"github.com/aabdullahgungor/golang-jwt-restapi/database"
	"github.com/aabdullahgungor/golang-jwt-restapi/model"
	"github.com/gin-gonic/gin"
)

func RegisterUSer(context *gin.Context) {
	var user model.User
	err := context.ShouldBindJSON(&user)
	if  err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	err = user.HashPassword(user.Password)
	if  err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username})
}