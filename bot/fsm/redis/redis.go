package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Storage struct {
	client *redis.Client
	ttl    time.Duration
}

func New(client *redis.Client) *Storage {
	return &Storage{
		client: client,
		ttl:    24 * time.Hour,
	}
}

func (s *Storage) SetState(chatID int, state string) error {

}
