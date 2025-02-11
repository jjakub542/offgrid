package views

import (
	"net/http"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Repo *repository.Repository
}

func (h *AdminHandler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/home.html", nil)
}
