package app

import (
	"fmt"
	"net/http"
	"offgrid/internal/database"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port int
	db   *pgx.Conn
}

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Renderer = Renderer()
	e.Static("/static", "web/static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.HomePage)
	e.GET("/hello", s.HelloWorldHandler)

	return e
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
