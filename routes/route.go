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

			secured.GET("/business", controllers.BusinessAll)
			secured.GET("/business/:business_id", controllers.BusinessOne)
			secured.POST("/business", controllers.BusinessAdd)
			secured.PUT("/business/:business_id", controllers.BusinessUpdate)
			secured.DELETE("/business/:business_id", controllers.BusinessDelete)
		}

	}
}
