package app

import (
	"errors"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"

	_ "github.com/joho/godotenv/autoload"
)

type templateRegistry struct {
	templates map[string]*template.Template
}

func (t *templateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)

		return err
	}

	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func Renderer() *templateRegistry {
	templates := make(map[string]*template.Template)

	templates["home.html"] = template.Must(template.ParseFiles("web/templates/user/home.html", "web/templates/user/base.html"))
	templates["contact.html"] = template.Must(template.ParseFiles("web/templates/user/contact.html", "web/templates/user/base.html"))
	templates["about.html"] = template.Must(template.ParseFiles("web/templates/user/about.html", "web/templates/user/base.html"))
	templates["login.html"] = template.Must(template.ParseFiles("web/templates/user/login.html", "web/templates/user/base.html"))
	templates["logout.html"] = template.Must(template.ParseFiles("web/templates/user/logout.html", "web/templates/user/base.html"))
	templates["admin/home.html"] = template.Must(template.ParseFiles("web/templates/admin/home.html", "web/templates/admin/base.html"))

	return &templateRegistry{
		templates: templates,
	}
}
