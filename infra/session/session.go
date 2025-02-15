package session

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/redis/go-redis/v9"
	"github.com/yusuke-takatsu/fishing-api-server/errors"
	"log"
	"net/http"
	"os"
	"time"
)

type Manager struct {
	redisClient *redis.Client
	cookieName  string
	sessionTLS  time.Duration
}

func NewSessionManager(redisClient *redis.Client, cookieName string, sessionTLS time.Duration) *Manager {
	return &Manager{
		redisClient: redisClient,
		cookieName:  cookieName,
		sessionTLS:  sessionTLS,
	}
}

func (m *Manager) RegenerateSession(ctx context.Context, w http.ResponseWriter, userID string) error {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		log.Printf("err: %v", err)
		return errors.InternalErr.Wrap("予期せぬエラーが発生しました。", err)
	}
	sessionID := hex.EncodeToString(b)

	if err := m.redisClient.Set(ctx, sessionID, userID, m.sessionTLS).Err(); err != nil {
		log.Printf("redis can not set session err: %v", err)
		return errors.InternalErr.Wrap("予期せぬエラーが発生しました。", err)
	}

	isSecure := true
	if os.Getenv("APP_ENV") == "local" {
		isSecure = false
	}
	cookie := &http.Cookie{
		Name:     m.cookieName,
		Value:    sessionID,
		Path:     "/",
		HttpOnly: true,
		Secure:   isSecure,
		Expires:  time.Now().Add(m.sessionTLS),
	}

	http.SetCookie(w, cookie)

	return nil
}
