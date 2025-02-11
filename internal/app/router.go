package app

import (
	"offgrid/internal/app/views"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo, repo *repository.Repository) {
	user := views.UserHandler{Repo: repo}
	userGroup := e.Group("/")
	userGroup.GET("", user.HomePage)
	userGroup.GET("/contact", user.ContactPage)

	api := views.ApiHandler{Repo: repo}
	apiGroup := e.Group("/api")
	apiGroup.GET("/hello", api.HelloWorldHandler)

	admin := views.AdminHandler{Repo: repo}
	adminGroup := e.Group("/admin")
	adminGroup.GET("", admin.HomePage)
}
