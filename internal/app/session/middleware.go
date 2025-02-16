package session

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const sessionCookieName = "session_id"

func Middleware(store *Store) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()

			cookie, err := req.Cookie(sessionCookieName)
			sessionID := ""
			if err == nil {
				sessionID = cookie.Value
			}

			if sessionID == "" {
				sessionID = generateSessionID()
				http.SetCookie(res, &http.Cookie{
					Name:     sessionCookieName,
					Value:    sessionID,
					Path:     "/",
					HttpOnly: true,
				})
			}

			c.Set("sessionID", sessionID)
			c.Set("sessionStore", store)

			return next(c)
		}
	}
}

func generateSessionID() string {
	return uuid.NewString()
}

func AdminAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionID := c.Get("sessionID").(string)
		store := c.Get("sessionStore").(*Store)

		role, ok := store.Get(sessionID, "role")
		if !ok || role != "admin" {
			return c.JSON(http.StatusForbidden, map[string]string{"error": "Access forbidden: Admin only"})
		}
		return next(c)
	}
}
