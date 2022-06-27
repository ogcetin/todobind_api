package controllers

import (
	"api/models"
	"net/http"
	"strconv"

	"api/core"
	"api/database"

	"github.com/gin-gonic/gin"
)

type User = models.User

var userList = []User{
	{UserID: 1, BusinessID: 1, Name: "Sabri Alışık", Status: "active"},
	{UserID: 2, BusinessID: 1, Name: "Emin Alışık", Status: "active"},
}

func UserLogin(c *gin.Context) {
	type UserForm struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=3,max=24"`
	}
	var user UserForm
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	row := database.DB.QueryRow("select user_id from user where email=? and password=?;", user.Email, user.Password)

	var user_id int
	err := row.Scan(&user_id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Something went wrong", "data": err})
		return
	}

	jwt_parsed := core.JWT{Email: user.Email, Password: user.Password}

	token, err := core.CreateToken(jwt_parsed)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Token not created", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": token})
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

func UserAdd(c *gin.Context) {
	type UserAddForm struct {
		Name       string `json:"name" binding:"required,min=3"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		Status     string `json:"status" binding:"required,min=3"`
		Email      string `json:"email" binding:"required,min=3"`
		Phone      string `json:"phone" binding:"required,min=3"`
	}
	var user UserAddForm
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	result, err := database.DB.Exec("INSERT INTO user (name, business_id, status, email, phone) VALUES (?, ?, ?, ?, ?)", user.Name, user.BusinessID, user.Status, user.Email, user.Phone)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	user_id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": user_id})
}

func UserUpdate(c *gin.Context) {
	user_id := c.Param("user_id")
	UserID, _ := strconv.Atoi(user_id)

	type UserUpdateForm struct {
		UserID     int    `json:"user_id" binding:"required,numeric"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		Name       string `json:"name" binding:"required"`
		Status     string `json:"status" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required"`
	}

	var user UserUpdateForm
	user.UserID = UserID
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	_, err := database.DB.Exec("update user set business_id, name=?, status=?, email=?, password=? where user_id=?", user.BusinessID, user.Name, user.Status, user.Email, user.Password, UserID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": user_id})
}

func UserDelete(c *gin.Context) {
	business_id := c.Param("business_id")
	BusinessID, _ := strconv.Atoi(business_id)
	_, err := database.DB.Exec("delete from business where business_id=?", BusinessID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": business_id})

}
