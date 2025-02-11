package app

import (
	"net/http"
	"offgrid/internal/app/views"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo, repo *repository.Repository) {

	e.GET("", HomePage)
	e.GET("/contact", ContactPage)

	//user := views.UserHandler{Repo: repo}
	//userGroup := e.Group("/user")

	api := views.ApiHandler{Repo: repo}
	apiGroup := e.Group("/api")
	apiGroup.GET("/hello", api.HelloWorldHandler)

	admin := views.AdminHandler{Repo: repo}
	adminGroup := e.Group("/admin")
	adminGroup.GET("", admin.HomePage)
}

func HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func ContactPage(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func AboutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", nil)
}

func LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

func LogoutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "logout.html", nil)
}
