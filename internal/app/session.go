package app

import (
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SessionData struct {
	Data      map[string]interface{}
	ExpiresAt time.Time
}

type SessionStore struct {
	sessions map[string]*SessionData
	mu       sync.RWMutex
}

func NewSessionStore() *SessionStore {
	return &SessionStore{
		sessions: make(map[string]*SessionData),
	}
}

func (s *SessionStore) Set(sessionID, key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if session, ok := s.sessions[sessionID]; ok {
		session.Data[key] = value
	} else {
		s.sessions[sessionID] = &SessionData{
			Data:      map[string]interface{}{key: value},
			ExpiresAt: time.Now().Add(30 * time.Minute),
		}
	}
}

func (s *SessionStore) Get(sessionID, key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if session, ok := s.sessions[sessionID]; ok {
		return session.Data[key], true
	}
	return nil, false
}

func (s *SessionStore) Delete(sessionID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, sessionID)
}

const sessionCookieName = "session_id"

func SessionMiddleware(store *SessionStore) echo.MiddlewareFunc {
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
