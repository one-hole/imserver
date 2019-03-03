package sources

import (
	"github.com/go-redis/redis"
	"github.com/w-zengtao/socket-server/sockets"
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
	redisdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       12,
	})
	_, err := redisdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	return &RedisSource{
		client: redisdb,
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
