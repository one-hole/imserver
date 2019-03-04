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
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "10.10.5.79:6379",
		DB:       12,
		PoolSize: 20,
	})
	var _, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	return &RedisSource{
		client: redisClient,
	}
}

// RunRedis will called in goroutines
func RunRedis(manager *sockets.ClientManager) {

	pubSub := RedisInstance().client.Subscribe(pushChannelName)
	defer pubSub.Close()

	ch := pubSub.Channel()

	for msg := range ch {
		manager.Broadcast <- []byte(msg.Payload)
	}
}
