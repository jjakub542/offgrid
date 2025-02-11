package views

import (
	"net/http"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Repo *repository.Repository
}

func (h *UserHandler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func (h *UserHandler) ContactPage(c echo.Context) error {
	return c.Render(http.StatusOK, "contact.html", nil)
}

func (h *UserHandler) AboutPage(c echo.Context) error {
	return c.Render(http.StatusOK, "about.html", nil)
}
