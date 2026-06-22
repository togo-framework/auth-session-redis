// Package redissession is a Redis-backed session store for togo auth
// (SESSION_DRIVER=redis). Install: `togo install togo-framework/auth-session-redis`.
package redissession

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/togo-framework/auth"
	"github.com/togo-framework/togo"
)

const prefix = "togo:session:"

func init() {
	auth.RegisterSessionStore("redis", func(k *togo.Kernel) (auth.SessionStore, error) {
		url := envOr("REDIS_URL", "redis://localhost:6379/0")
		opt, err := redis.ParseURL(url)
		if err != nil {
			return nil, err
		}
		return &store{rdb: redis.NewClient(opt)}, nil
	})
}

type store struct{ rdb *redis.Client }

func (s *store) Put(ctx context.Context, sid, token string, ttl time.Duration) error {
	return s.rdb.Set(ctx, prefix+sid, token, ttl).Err()
}

func (s *store) Get(ctx context.Context, sid string) (string, bool, error) {
	v, err := s.rdb.Get(ctx, prefix+sid).Result()
	if err == redis.Nil {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	return v, true, nil
}

func (s *store) Delete(ctx context.Context, sid string) error {
	return s.rdb.Del(ctx, prefix+sid).Err()
}

func envOr(k, d string) string {
	if v := getenv(k); v != "" {
		return v
	}
	return d
}
