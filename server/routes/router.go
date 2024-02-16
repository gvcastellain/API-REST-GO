package routes

import (
	"api_rest/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		movies := main.Group("movies")
		{
			movies.GET("/", controllers.ShowMovies)
		}
	}
	return router
}
