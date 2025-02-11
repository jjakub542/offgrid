package views

import (
	"net/http"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

type ApiHandler struct {
	Repo *repository.Repository
}

func (h *ApiHandler) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}
