package controllers

import (
	"api/models"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type User = models.User

var userList = []User{
	{UserID: 1, BusinessID: 1, Name: "Sabri Alışık", Status: "active"},
	{UserID: 2, BusinessID: 1, Name: "Emin Alışık", Status: "active"},
}

func UserAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "", "message": "", "data": userList})
}
func UserOne(c *gin.Context) {
	user_id := c.Param("user_id")
	UserID, _ := strconv.Atoi(user_id)

	for i := 0; i < len(userList); i++ {
		if userList[i].UserID == UserID {
			c.JSON(http.StatusOK, gin.H{"status": "", "message": "", "data": userList[i]})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "", "message": "Business not found", "data": ""})
}
