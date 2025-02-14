package app

import (
	"offgrid/internal/app/session"
	"offgrid/internal/app/views"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, repo *repository.Repository) {
	user := views.UserHandler{Repo: repo}
	admin := views.AdminHandler{Repo: repo}

	userGroup := e.Group("")
	userGroup.GET("", user.HomePage)
	userGroup.GET("/about", user.AboutPage)
	userGroup.GET("/contact", user.ContactPage)

	adminGroup := e.Group("/admin")
	adminGroup.GET("", session.AdminAuthMiddleware(admin.HomePage))
	adminGroup.GET("/articles", session.AdminAuthMiddleware(admin.ArticlesPage))
	adminGroup.POST("/articles/create", session.AdminAuthMiddleware(admin.ArticleCreate))
	adminGroup.POST("/articles/delete/:article_id", session.AdminAuthMiddleware(admin.ArticleDelete))
	adminGroup.POST("/articles/update/:article_id", session.AdminAuthMiddleware(admin.ArticleUpdate))
	adminGroup.POST("/articles/attach-image/:article_id", session.AdminAuthMiddleware(admin.ArticleAttachImage))
	adminGroup.GET("/articles/edit/:article_id", session.AdminAuthMiddleware(admin.ArticleEditPage))
	adminGroup.Any("/login", admin.LoginPage)
	adminGroup.Any("/logout", admin.LogoutPage)
}
