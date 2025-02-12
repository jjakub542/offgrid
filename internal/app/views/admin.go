package views

import (
	"fmt"
	"net/http"
	"offgrid/internal/app/session"
	"offgrid/internal/domain"
	"offgrid/internal/repository"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Repo *repository.Repository
}

func (h *AdminHandler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/home.html", nil)
}

func (h *AdminHandler) ArticlesPage(c echo.Context) error {
	articles, err := h.Repo.Article.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Render(http.StatusOK, "admin/articles.html", articles)
}

func (h *AdminHandler) ArticleEditPage(c echo.Context) error {
	return c.Render(http.StatusOK, "admin/article.html", nil)
}

func (h *AdminHandler) LoginPage(c echo.Context) error {
	sessionID := c.Get("sessionID").(string)
	store := c.Get("sessionStore").(*session.Store)

	if c.Request().Method == http.MethodGet {
		role, ok := store.Get(sessionID, "role")
		if ok || role == "admin" {
			return c.Redirect(http.StatusSeeOther, "/admin")
		}
		return c.Render(http.StatusOK, "admin/login.html", nil)
	}

	requestUser := domain.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	requestUser.CreatePasswordHash()

	user, err := h.Repo.User.GetOneByEmail(c.FormValue("email"))

	fmt.Println(user.Email)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User with this email does not exist")
	}

	if user.PasswordHash != requestUser.PasswordHash {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password")
	}

	if !user.IsSuperuser {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	store.Set(sessionID, "authenticated", true)
	store.Set(sessionID, "role", "admin")

	return c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *AdminHandler) LogoutPage(c echo.Context) error {
	sessionID := c.Get("sessionID").(string)
	store := c.Get("sessionStore").(*session.Store)

	store.Delete(sessionID)
	return c.Render(http.StatusOK, "admin/logout.html", nil)
}
