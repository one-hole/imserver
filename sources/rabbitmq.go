package sources

/*
	1. 这里打算使用 RabbitMQ 的 Direct Exchange
	2. 按照我们之后的数据产生量(赔率调节)、这里其实使用 唯一一个 Queue 来消费就足够了
	3. 如果之后发现需要多个 Queue & 那么生成多个 Channel 即可 & 并且使用 Fanout类型即可
	4. 或者 go Run() 之后分发给不同的 Manager 也可以
*/

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/w-zengtao/socket-server/config"
	"github.com/w-zengtao/socket-server/sockets"
	"github.com/w-zengtao/socket-server/utils"
)

var instance *RabbitSource

const (
	exchangeName = "rw-hz-odds-direct"
	exchangeType = "direct"
	routingKey   = "rw-hz-odds-routing"
)

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

// RunRabbit will call in goroutines
func RunRabbit(manager *sockets.ClientManager) {
	channel, err := RabbitInstance().conn.Channel()
	utils.FailOnError(err, "Failed to open a Channel")
	defer channel.Close()

	err = channel.ExchangeDeclare(exchangeName, exchangeType, false, false, false, false, nil)
	utils.FailOnError(err, "Failed to declare a Exchange")

	queue, err := channel.QueueDeclare("", false, false, true, false, nil)
	utils.FailOnError(err, "Failed to declare a Queue")

	err = channel.QueueBind(queue.Name, routingKey, exchangeName, false, nil)
	utils.FailOnError(err, "Failed to bind a queue")

	msgs, err := channel.Consume(queue.Name, "", true, false, false, false, nil)

	utils.FailOnError(err, "Failed to consume the queue")

	fmt.Println("Before of datasource Run()")

	// 这里需要写入 Manger
	for d := range msgs {
		manager.Broadcast <- d.Body
	}

	fmt.Println("End of datasource Run()")

}

// Private
func newInstance() *RabbitSource {
	conn, _ := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:5672/", config.Instance().Rabbit.User, config.Instance().Rabbit.Password, config.Instance().Rabbit.Host))
	return &RabbitSource{
		conn: conn,
	}
}
