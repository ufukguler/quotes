package routes

import (
	"github.com/labstack/echo/v4"
	"quotes/handlers"
)

func UseRoute(e *echo.Echo) {
	api := e.Group("/api")

	category := api.Group("/category")
	{
		category.GET("", handlers.GetCategories)
		category.GET("/:id", handlers.GetCategoryById)
	}

	quote := api.Group("/quote")
	{
		quote.GET("", handlers.GetQuote)
		quote.GET("/:id", handlers.GetQuoteById)
		quote.POST("/:id/like", handlers.LikeQuoteById)
		quote.POST("/:id/dislike", handlers.DisLikeQuoteById)
		quote.GET("/category/:id", handlers.GetQuoteByCategoryId)
	}

	user := api.Group("/user")
	{
		user.GET("", handlers.GetCurrentUser)
		user.GET("/quotes", handlers.GetCurrentUserLikeList)
	}

	// auth
	auth := api.Group("/auth")
	{
		auth.POST("/login", handlers.Login)
	}

	// admin
	admin := api.Group("/admin")
	{
		admin.POST("/category", handlers.SaveCategory)
		admin.POST("/quote", handlers.SaveQuote)
	}

	stats := api.Group("/stats")
	{
		stats.GET("", handlers.GetStats)
	}

	version := api.Group("/version")
	{
		version.GET("", handlers.Version)
	}
}
