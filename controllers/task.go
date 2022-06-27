package controllers

import (
	"api/models"
	"net/http"

	"strconv"

	"api/database"

	"github.com/gin-gonic/gin"
)

type Task = models.Task

func TaskAll(c *gin.Context) {

	rows, err := database.DB.Query("select * from task")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
		return
	}
	defer rows.Close()

	var taskList []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.TaskID, &t.ProjectID, &t.BusinessID, &t.CreatorID, &t.AttendantID, &t.CreationDate, &t.DueDate, &t.Detail)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
			return
		}
		taskList = append(taskList, t)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": taskList})
}

func TaskOne(c *gin.Context) {
	task_id := c.Param("task_id")
	TeamID, _ := strconv.Atoi(task_id)

	row := database.DB.QueryRow("select task_id,project_id,business_id,creator_id,attendant_id,creation_date,due_date,detail from task where task_id = ?;", TeamID)

	var task Task
	err := row.Scan(&task.TaskID, &task.ProjectID, &task.BusinessID, &task.CreatorID, &task.AttendantID, &task.CreationDate, &task.DueDate, &task.Detail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Something went wrong", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": task})
}

func TaskAdd(c *gin.Context) {
	type TaskAddForm struct {
		ProjectID    int    `json:"project_id" binding:"required,numeric"`
		BusinessID   int    `json:"business_id" binding:"required,numeric"`
		CreatorID    int    `json:"creator_id" binding:"required,numeric"`
		AttendantID  int    `json:"attendant_id" binding:"required,numeric"`
		CreationDate string `json:"creation_date" binding:"required"`
		DueDate      string `json:"due_date" binding:"required"`
		Detail       string `json:"detail" binding:"required"`
	}
	var task TaskAddForm
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	result, err := database.DB.Exec("INSERT INTO task (project_id,business_id,creator_id,attendant_id,creation_date,due_date,detail) VALUES (?, ?, ?)", task.ProjectID, task.BusinessID, task.CreatorID, task.AttendantID, task.CreationDate, task.DueDate, task.Detail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	task_id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": task_id})
}

func TaskUpdate(c *gin.Context) {
	task_id := c.Param("task_id")
	TaskID, _ := strconv.Atoi(task_id)

	type TaskUpdateForm struct {
		TaskID       int    `json:"task_id" binding:"required,numeric"`
		ProjectID    int    `json:"project_id" binding:"required,numeric"`
		BusinessID   int    `json:"business_id" binding:"required,numeric"`
		CreatorID    int    `json:"creator_id" binding:"required,numeric"`
		AttendantID  int    `json:"attendant_id" binding:"required,numeric"`
		CreationDate string `json:"creation_date" binding:"required"`
		DueDate      string `json:"due_date" binding:"required"`
		Detail       string `json:"detail" binding:"required"`
	}

	var task TaskUpdateForm
	task.TaskID = TaskID
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	_, err := database.DB.Exec("update task set project_id,business_id,creator_id,attendant_id,creation_date,due_date,detail where task_id=?", task.ProjectID, task.BusinessID, task.CreatorID, task.AttendantID, task.CreationDate, task.DueDate, task.Detail)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": task_id})
}

func TaskDelete(c *gin.Context) {
	task_id := c.Param("task_id")
	TaskID, _ := strconv.Atoi(task_id)
	_, err := database.DB.Exec("delete from task where task_id=?", TaskID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": task_id})
}
