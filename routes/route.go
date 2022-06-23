package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func Prepare(r *gin.Engine) {

	r.GET("/user", controllers.UserAll)
	r.GET("/user/:user_id", controllers.UserOne)

	r.GET("/business", controllers.BusinessAll)
	r.GET("/business/:business_id", controllers.BusinessOne)
	r.POST("/business", controllers.BusinessAdd)
	r.PUT("/business/:business_id", controllers.BusinessUpdate)
	r.DELETE("/business/:business_id", controllers.BusinessDelete)
}
