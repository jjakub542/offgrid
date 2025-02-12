package app

import (
	"fmt"
	"net/http"
	"offgrid/internal/app/session"
	"offgrid/internal/database"
	"offgrid/internal/repository"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port       int
	db         *pgxpool.Pool
	store      *session.Store
	repository *repository.Repository
}

func (s *Server) Handler() http.Handler {
	e := echo.New()
	e.Use(session.Middleware(s.store))
	e.Renderer = TemplateRenderer()
	e.Static("/static", "web/static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	RegisterRoutes(e, s.repository)

	return e
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.Connect()
	store := session.NewStore()
	NewServer := &Server{
		port:       port,
		db:         db,
		store:      store,
		repository: repository.New(db),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.Handler(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
