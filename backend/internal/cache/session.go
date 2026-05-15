package cache

import (
	"context"
	"fmt"
	"time"
)

const (
	SessionTTL       = 1 * time.Hour
	SessionKeyPrefix = "session:%s"
)

type SessionStore struct {
	cache Service
}

func NewSessionStore(cache Service) *SessionStore {
	return &SessionStore{cache: cache}
}

func (s *SessionStore) SetSession(ctx context.Context, token string, data string) error {
	key := fmt.Sprintf(SessionKeyPrefix, token)
	return s.cache.Set(ctx, key, data, SessionTTL)
}

func (s *SessionStore) GetSession(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf(SessionKeyPrefix, token)
	return s.cache.Get(ctx, key)
}

func (s *SessionStore) DeleteSession(ctx context.Context, token string) error {
	key := fmt.Sprintf(SessionKeyPrefix, token)
	return s.cache.Del(ctx, key)
}
