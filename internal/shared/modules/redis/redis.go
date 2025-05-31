package redis

import (
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/config"
	"github.com/cristiano-pacheco/goflix/pkg/redis"
)

func NewRedis(config config.Config) redis.Redis {
	return redis.NewRedis(config.Redis.Addr, config.Redis.Password, config.Redis.DB)
}
