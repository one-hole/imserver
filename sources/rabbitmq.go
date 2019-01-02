package sources

/*
	1. 这里打算使用 RabbitMQ 的 Direct Exchange
	2. 按照我们之后的数据产生量(赔率调节)、这里其实使用 唯一一个 Queue 来消费就足够了
	3. 如果之后发现需要多个 Queue & 那么生成多个 Channel 即可
*/

import (
	"gitee.com/odd-socket/utils"
	"github.com/streadway/amqp"
)

var instance *RabbitSource

const (
	exchangeName = "rw-hz-odds-direct"
	exchangeType = "direct"
	routingKey   = "rw-hz-odds-routing"
)

func init() {
	go Run()
}

// RabbitSource 从 RabbitMQ 读取数据、写入 Socket Client 的 message channel
// 因为这里只有 连接是要被复用的 & 所以只需要保持连接的对象即可
type RabbitSource struct {
	conn *amqp.Connection
}

// RabbitInstance returns the singleton instance of RabbitMQ
func RabbitInstance() *RabbitSource {
	if instance == nil {
		instance = newInstance()
	}
	return instance
}

// Close will release resources
func Close() {

}

// Run will call in goroutines
func Run() {
	channel, err := RabbitInstance().conn.Channel()
	utils.FailOnError(err, "Failed to open a Channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(exchangeName, exchangeType, true, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a Queue")

	queue, err := channel.QueueDeclare("", false, false, true, false, nil)
	utils.FailOnError(err, "Failed to declare a Queue")

	err = channel.QueueBind(queue.Name, routingKey, exchangeName, false, nil)

}

// Private
func newInstance() *RabbitSource {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	return &RabbitSource{
		conn: conn,
	}
}
