package controllers

import (
	"api/models"
	"net/http"

	"strconv"

	"api/database"

	"github.com/gin-gonic/gin"
)

type Project = models.Project

func ProjectAll(c *gin.Context) {

	rows, err := database.DB.Query("select * from project")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
		return
	}
	defer rows.Close()

	var projectList []Project
	for rows.Next() {
		var p Project
		err := rows.Scan(&p.ProjectID, &p.Name, &p.TeamID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
			return
		}
		projectList = append(projectList, p)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": projectList})
}

func ProjectOne(c *gin.Context) {
	project_id := c.Param("project_id")
	ProjectID, _ := strconv.Atoi(project_id)

	row := database.DB.QueryRow("select project_id,business_id,name,team_id from project where project_id = ?;", ProjectID)

	var project Project
	err := row.Scan(&project.ProjectID, &project.Name, &project.TeamID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Something went wrong", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": project})
}

func ProjectAdd(c *gin.Context) {
	type ProjectAddForm struct {
		Name       string `json:"name" binding:"required,min=3"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		TeamID     int    `json:"team_id" binding:"required,numeric"`
	}
	var project ProjectAddForm
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	result, err := database.DB.Exec("INSERT INTO project (name, team_id, business_id) VALUES (?, ?, ?)", project.Name, project.TeamID, project.BusinessID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	project_id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": project_id})
}

func ProjectUpdate(c *gin.Context) {
	project_id := c.Param("project_id")
	ProjectID, _ := strconv.Atoi(project_id)

	type ProjectUpdateForm struct {
		ProjectID  int    `json:"project_id" binding:"required,numeric"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		TeamID     int    `json:"team_id" binding:"required,numeric"`
		Name       string `json:"name" binding:"required"`
	}

	var project ProjectUpdateForm
	project.ProjectID = ProjectID
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	_, err := database.DB.Exec("update project set name=?, team_id=?, business_id=? where project_id=?", project.Name, project.TeamID, project.BusinessID, ProjectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": project_id})
}

func ProjectDelete(c *gin.Context) {
	project_id := c.Param("project_id")
	ProjectID, _ := strconv.Atoi(project_id)
	_, err := database.DB.Exec("delete from project where project_id=?", ProjectID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": project_id})

}
