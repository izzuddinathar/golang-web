package routes

import (
	"golang-web/internal/handlers"
	"golang-web/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Static("/static", "./static")

	// Public routes
	router.GET("/", handlers.ShowLoginPage)
	router.POST("/login", handlers.Login)
	router.POST("/logout", handlers.Logout)

	// Protected routes
	auth := router.Group("/")
	auth.Use(middlewares.RequireLogin)
	{
		auth.GET("/dashboard", handlers.ShowDashboard)
		auth.GET("/users", handlers.ListUsers)
		auth.GET("/users/create", handlers.ShowCreateUser)
		auth.POST("/users/create", handlers.CreateUser)
		auth.GET("/users/edit/:id", handlers.ShowEditUser)
		auth.POST("/users/edit/:id", handlers.UpdateUser)
		auth.POST("/users/delete/:id", handlers.DeleteUser)
	}
}
