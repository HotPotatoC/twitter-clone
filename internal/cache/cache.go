package cache

import "time"

type Cache interface {
	Ping() error
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) (int64, error)
}
