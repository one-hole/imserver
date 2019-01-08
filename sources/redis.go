package sources

import "github.com/go-redis/redis"

var redisInstance *RedisSource

// RedisSource connected to redis & sub the messages
type RedisSource struct {
	Client *redis.Client
}

// RedisInstance returns the singleton instance of Redis
func RedisInstance() *RedisSource {
	if redisInstance == nil {
		redisInstance = newRedisInstance()
	}
	return redisInstance
}

func newRedisInstance() *RedisSource {
	return &RedisSource{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}
