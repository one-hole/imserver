package sources

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/one-hole/imserver/config"
	"github.com/one-hole/imserver/sockets"
)

var redisInstance *RedisSource

var (
	pushChannelName = "tenant-websocket"
)

// RedisSource connected to redis & sub the messages
type RedisSource struct {
	client *redis.Client
}

// RedisPing send ping message to redis
func RedisPing() error {
	_, err := redisInstance.client.Ping().Result()
	return err
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
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port),
			DB:       12,
		}),
	}
}

// RunRedis will called in goroutines
func RunRedis(manager *sockets.ClientManager, channelName string) {
	if "" == channelName {
		return
	}

	pubsub := RedisInstance().client.Subscribe(channelName)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		manager.Broadcast <- []byte(msg.Payload)
	}
}
