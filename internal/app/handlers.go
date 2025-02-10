package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) UserHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "user/home.html", nil)
}

func (s *Server) AdminHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/home.html", nil)
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
