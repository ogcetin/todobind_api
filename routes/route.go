package routes

import (
	"api/controllers"
	"api/core"

	"github.com/gin-gonic/gin"
)

func Prepare(r *gin.Engine) {
	api := r.Group("/api")
	{
		//
		api.POST("/user/login", controllers.UserLogin)
		secured := api.Group("/").Use(core.JWTAuth())
		{
			secured.GET("/user", controllers.UserAll)
			secured.GET("/user/:user_id", controllers.UserOne)
			secured.POST("/user", controllers.UserAdd)
			secured.PUT("/user/:user_id", controllers.UserUpdate)
			secured.DELETE("/user/:user_id", controllers.UserDelete)

			secured.GET("/business", controllers.BusinessAll)
			secured.GET("/business/:business_id", controllers.BusinessOne)
			secured.POST("/business", controllers.BusinessAdd)
			secured.PUT("/business/:business_id", controllers.BusinessUpdate)
			secured.DELETE("/business/:business_id", controllers.BusinessDelete)

			secured.GET("/project", controllers.ProjectAll)
			secured.GET("/project/:project_id", controllers.ProjectOne)
			secured.POST("/project", controllers.ProjectAdd)
			secured.PUT("/project/:project_id", controllers.ProjectUpdate)
			secured.DELETE("/project/:project_id", controllers.ProjectDelete)

			secured.GET("/section", controllers.SectionAll)
			secured.GET("/section/:section_id", controllers.SectionOne)
			secured.POST("/section", controllers.SectionAdd)
			secured.PUT("/section/:section_id", controllers.SectionUpdate)
			secured.DELETE("/section/:section_id", controllers.SectionDelete)

			secured.GET("/task", controllers.TaskAll)
			secured.GET("/task/:task_id", controllers.TaskOne)
			secured.POST("/task", controllers.TaskAdd)
			secured.PUT("/task/:task_id", controllers.TaskUpdate)
			secured.DELETE("/task/:task_id", controllers.TaskDelete)

			secured.GET("/team", controllers.TeamAll)
			secured.GET("/team/:team_id", controllers.TeamOne)
			secured.POST("/team", controllers.TeamAdd)
			secured.PUT("/team/:team_id", controllers.TeamUpdate)
			secured.DELETE("/team/:team_id", controllers.TeamDelete)

		}

	}
}
