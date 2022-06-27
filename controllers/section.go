package controllers

import (
	"api/models"
	"net/http"

	"strconv"

	"api/database"

	"github.com/gin-gonic/gin"
)

type Section = models.Section

func SectionAll(c *gin.Context) {

	rows, err := database.DB.Query("select * from section")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
		return
	}
	defer rows.Close()

	var sectionList []Section
	for rows.Next() {
		var s Section
		err := rows.Scan(&s.SectionID, &s.Name, &s.ProjectID, &s.BusinessID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
			return
		}
		sectionList = append(sectionList, s)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": sectionList})
}

func SectionOne(c *gin.Context) {
	section_id := c.Param("section_id")
	SectionID, _ := strconv.Atoi(section_id)

	row := database.DB.QueryRow("select section_id,business_id,name,project_id from section where section_id = ?;", SectionID)

	var section Section
	err := row.Scan(&section.BusinessID, &section.Name, &section.SectionID, &section.ProjectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Something went wrong", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": section})
}

func SectionAdd(c *gin.Context) {
	type SectionAddForm struct {
		Name       string `json:"name" binding:"required,min=3"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		ProjectID  int    `json:"project_id" binding:"required,numeric"`
	}
	var section SectionAddForm
	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	result, err := database.DB.Exec("INSERT INTO section (name, project_id, business_id) VALUES (?, ?, ?)", section.Name, section.ProjectID, section.BusinessID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	section_id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": section_id})
}

func SectionUpdate(c *gin.Context) {
	section_id := c.Param("section_id")
	SectionID, _ := strconv.Atoi(section_id)

	type SectionUpdateForm struct {
		SectionID  int    `json:"section_id" binding:"required,numeric"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		Name       string `json:"name" binding:"required"`
		ProjectID  string `json:"project_id" binding:"required,numeric"`
	}

	var section SectionUpdateForm
	section.SectionID = SectionID
	if err := c.ShouldBindJSON(&section); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	_, err := database.DB.Exec("update section set name=?, business_id, project_id where business_id=?", section.Name, section.BusinessID, section.ProjectID, SectionID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": section_id})
}

func SectionDelete(c *gin.Context) {
	section_id := c.Param("section_id")
	SectionID, _ := strconv.Atoi(section_id)
	_, err := database.DB.Exec("delete from section where section_id=?", SectionID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": section_id})

}
