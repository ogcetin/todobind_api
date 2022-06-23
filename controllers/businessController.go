package controllers

import (
	"api/models"
	"net/http"

	"strconv"

	"api/database"

	"github.com/gin-gonic/gin"
)

type Business = models.Business

func BusinessAll(c *gin.Context) {

	rows, err := database.DB.Query("select * from business")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
		return
	}
	defer rows.Close()

	var businessList []Business
	for rows.Next() {
		var b Business
		err := rows.Scan(&b.BusinessID, &b.Name, &b.Status)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
			return
		}
		businessList = append(businessList, b)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": businessList})
}

func BusinessOne(c *gin.Context) {
	business_id := c.Param("business_id")
	BusinessID, _ := strconv.Atoi(business_id)

	row := database.DB.QueryRow("select business_id,name,status from business where business_id = ?;", BusinessID)

	var business Business
	err := row.Scan(&business.BusinessID, &business.Name, &business.Status)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Something went wrong", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": business})
}

func BusinessAdd(c *gin.Context) {
	type BusinessAddForm struct {
		Name string `json:"name" binding:"required,min=3"`
	}
	var business BusinessAddForm
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	result, err := database.DB.Exec("INSERT INTO business (name, status) VALUES (?, ?)", business.Name, "active")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	business_id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": business_id})
}

func BusinessUpdate(c *gin.Context) {
	business_id := c.Param("business_id")
	BusinessID, _ := strconv.Atoi(business_id)

	type BusinessUpdateForm struct {
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		Name       string `json:"name" binding:"required"`
		Status     string `json:"status" binding:"required"`
	}

	var business BusinessUpdateForm
	business.BusinessID = BusinessID
	if err := c.ShouldBindJSON(&business); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	_, err := database.DB.Exec("update business set name=?, status=? where business_id=?", business.Name, business.Status, BusinessID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": business_id})
}

func BusinessDelete(c *gin.Context) {
	business_id := c.Param("business_id")
	BusinessID, _ := strconv.Atoi(business_id)
	_, err := database.DB.Exec("delete from business where business_id=?", BusinessID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": business_id})

}
