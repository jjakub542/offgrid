package views

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"offgrid/internal/app/session"
	"offgrid/internal/domain"
	"offgrid/internal/repository"
	"os"

	"github.com/google/uuid"
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

func (h *AdminHandler) ArticleCreate(c echo.Context) error {
	article := &domain.Article{
		Title:       c.FormValue("title"),
		Description: c.FormValue("desc"),
		Public:      false,
	}
	err := h.Repo.Article.CreateOne(article)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/articles")
}

func (h *AdminHandler) ArticleDelete(c echo.Context) error {
	err := h.Repo.Article.DeleteOneById(c.Param("article_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/articles")
}

func (h *AdminHandler) ArticleUpdate(c echo.Context) error {
	article := &domain.Article{
		Title:       c.FormValue("title"),
		Description: c.FormValue("desc"),
		Content:     c.FormValue("content"),
	}
	if c.FormValue("public") == "on" {
		article.Public = true
	} else {
		article.Public = false
	}
	err := h.Repo.Article.UpdateOneById(article, c.Param("article_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/articles")
}

func (h *AdminHandler) ArticleAttachImage(c echo.Context) error {
	filename := uuid.NewString() + ".png"
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad request")
	}
	defer src.Close()
	dst, err := os.Create("web/static/uploads/" + filename)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	image := &domain.Image{Filename: filename}
	err = h.Repo.Article.AttachImage(image, c.Param("article_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/articles/edit/"+c.Param("article_id"))
}

func (h *AdminHandler) ArticleEditPage(c echo.Context) error {
	article, err := h.Repo.Article.GetOneById(c.Param("article_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	images, err := h.Repo.Article.GetArticleImages(c.Param("article_id"))
	if err != nil {
		log.Fatal(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	for _, image := range images {
		article.ImageUrls = append(article.ImageUrls, image.GetUrl())
	}
	return c.Render(http.StatusOK, "admin/article.html", article)
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
