package controllers

import (
	"api/models"
	"net/http"

	"strconv"

	"api/database"

	"github.com/gin-gonic/gin"
)

type Team = models.Team

func TeamAll(c *gin.Context) {

	rows, err := database.DB.Query("select * from team")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
		return
	}
	defer rows.Close()

	var teamList []Team
	for rows.Next() {
		var te Team
		err := rows.Scan(&te.TeamID, &te.Name, &te.CreatorID, &te.BusinessID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Something went wrong", "data": err})
			return
		}
		teamList = append(teamList, te)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": teamList})
}

func TeamOne(c *gin.Context) {
	team_id := c.Param("team_id")
	TeamID, _ := strconv.Atoi(team_id)

	row := database.DB.QueryRow("select team_id,business_id,name,creator_id from team where team_id = ?;", TeamID)

	var team Team
	err := row.Scan(&team.BusinessID, &team.Name, &team.TeamID, &team.CreatorID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "", "message": "Something went wrong", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": team})
}

func TeamAdd(c *gin.Context) {
	type TeamAddForm struct {
		Name       string `json:"name" binding:"required,min=3"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		CreatorID  int    `json:"creator_id" binding:"required,numeric"`
	}
	var team TeamAddForm
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	result, err := database.DB.Exec("INSERT INTO team (name, creator_id, business_id) VALUES (?, ?, ?)", team.Name, team.CreatorID, team.BusinessID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	team_id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": team_id})
}

func TeamUpdate(c *gin.Context) {
	team_id := c.Param("team_id")
	TeamID, _ := strconv.Atoi(team_id)

	type TeamUpdateForm struct {
		TeamID     int    `json:"team_id" binding:"required,numeric"`
		BusinessID int    `json:"business_id" binding:"required,numeric"`
		Name       string `json:"name" binding:"required"`
		CreatorID  string `json:"creator_id" binding:"required,numeric"`
	}

	var team TeamUpdateForm
	team.TeamID = TeamID
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
		return
	}

	_, err := database.DB.Exec("update team set name=?, business_id, creator_id where team_id=?", team.Name, team.BusinessID, team.CreatorID, TeamID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": team_id})
}

func TeamDelete(c *gin.Context) {
	team_id := c.Param("team_id")
	TeamID, _ := strconv.Atoi(team_id)
	_, err := database.DB.Exec("delete from team where team_id=?", TeamID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error(), "data": ""})
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "ok", "data": team_id})

}
