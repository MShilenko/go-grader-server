package redis

import (
	"fmt"
	"time"

	redicsCl "github.com/go-redis/redis/v7"

	"github.com/MShilenko/go-grader-server/configs"
)

func GetRedisClient(cfg *configs.Config) (*redicsCl.Client, error) {
	cl := redicsCl.NewClient(&redicsCl.Options{
		Addr:         fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		DialTimeout:  time.Duration(cfg.Redis.DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.Redis.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Redis.WriteTimeout) * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
	})
	stat := cl.Ping()
	if stat.Err() != nil {
		return nil, stat.Err()
	}

	return cl, nil
}
