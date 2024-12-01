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

		auth.GET("/menus", handlers.ListMenus)
		auth.GET("/menus/create", handlers.ShowCreateMenu)
		auth.POST("/menus/create", handlers.CreateMenu)
		auth.GET("/menus/edit/:id", handlers.ShowEditMenu)
		auth.POST("/menus/edit/:id", handlers.UpdateMenu)
		auth.POST("/menus/delete/:id", handlers.DeleteMenu)

		auth.GET("/tables", handlers.ListTables)
		auth.GET("/tables/create", handlers.ShowCreateTable)
		auth.POST("/tables/create", handlers.CreateTable)
		auth.GET("/tables/edit/:id", handlers.ShowEditTable)
		auth.POST("/tables/edit/:id", handlers.UpdateTable)
		auth.POST("/tables/delete/:id", handlers.DeleteTable)

		auth.GET("/orders", handlers.ListOrders)
		auth.GET("/orders/create", handlers.ShowCreateOrder)
		auth.POST("/orders/create", handlers.CreateOrder)
		auth.GET("/orders/edit/:id", handlers.ShowEditOrder)
		auth.POST("/orders/edit/:id", handlers.UpdateOrder)
		auth.POST("/orders/delete/:id", handlers.DeleteOrder)

		auth.GET("/payments", handlers.ListPayments)
		auth.GET("/payments/create", handlers.ShowCreatePayment)
		auth.POST("/payments/create", handlers.CreatePayment)
		auth.GET("/payments/edit/:id", handlers.ShowEditPayment)
		auth.POST("/payments/edit/:id", handlers.UpdatePayment)
		auth.POST("/payments/delete/:id", handlers.DeletePayment)

	}
}
