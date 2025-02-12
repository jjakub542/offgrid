package app

import (
	"offgrid/internal/app/session"
	"offgrid/internal/app/views"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, repo *repository.Repository) {
	user := views.UserHandler{Repo: repo}
	userGroup := e.Group("")
	userGroup.GET("", user.HomePage)
	userGroup.GET("/about", user.AboutPage)
	userGroup.GET("/contact", user.ContactPage)

	api := views.ApiHandler{Repo: repo}
	apiGroup := e.Group("/api")
	apiGroup.GET("/hello", api.HelloWorldHandler)

	admin := views.AdminHandler{Repo: repo}
	adminGroup := e.Group("/admin")
	adminGroup.GET("", session.AdminAuthMiddleware(admin.HomePage))
	adminGroup.GET("/articles", session.AdminAuthMiddleware(admin.ArticlesPage))
	adminGroup.GET("/articles/:article_id", session.AdminAuthMiddleware(admin.ArticleEditPage))
	adminGroup.Any("/login", admin.LoginPage)
	adminGroup.GET("/logout", admin.LogoutPage)
}
