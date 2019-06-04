package sources

import (
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

// RedisInstance returns the singleton instance of Redis
func RedisInstance() *RedisSource {
	if redisInstance == nil {
		redisInstance = newRedisInstance()
	}
	return redisInstance
}

func newRedisInstance() *RedisSource {

	source := &RedisSource{
		client: redis.NewClient(&redis.Options{
			Addr: config.Instance().Redis.Addr,
			DB:   12,
		}),
	}

	_, err := source.client.Ping().Result()

	if err != nil {
		panic(err)
	}

	return source
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
