package sources

import (
	"github.com/go-redis/redis"
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
			Addr:     "localhost:6379",
			Password: "",
			DB:       12,
		}),
	}
}

// RunRedis will called in goroutines
func RunRedis(manager *sockets.ClientManager) {

	pubsub := RedisInstance().client.Subscribe(pushChannelName)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		manager.Broadcast <- []byte(msg.Payload)
	}
}
