package routes

import (
	"github.com/AmineGoirech/gin-auth/controllers"
	"github.com/AmineGoirech/gin-auth/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/login", controllers.Login)
	r.POST("/register",controllers.Register)
	// r.GET("/logout",controllers.Logout)
	// r.GET("/refreshtoken",controllers.RefreshToken)

	private := r.Group("/private")
	private.Use(middleware.Authenticate)

	private.GET("/refreshtoken",controllers.RefreshToken)

}
