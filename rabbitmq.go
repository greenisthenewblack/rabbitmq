package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

//RabbitQueue - a configuration object detailing a rabbit queue
// type RabbitQueue struct {
// 	name         string
// 	durable      bool
// 	deleteUnused bool
// 	exclusive    bool
// 	noWait       bool
// }

//RabbitConnection - a configuration object for the connection to rabbit
type RabbitConnection struct {
	URL  string
	Port string
	User string
	Pass string
}

//Rabbit - create rabbit object
type Rabbit struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      *amqp.Queue
}

//Create - create connection
func Create(rbtc RabbitConnection) Rabbit {
	var connStr string
	if rbtc.User!="" 
		connStr = fmt.Sprintf("amqp://%v:%v@%v:%v/", rbtc.User, rbtc.Pass, rbtc.URL, rbtc.Port)
	}else{
		connStr = fmt.Sprintf("%v:%v/", rbtc.URL, rbtc.Port)
	}
	
	conn, err := amqp.Dial(connStr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	rbt := Rabbit{Connection: conn}
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	rbt.Channel = ch
	defer ch.Close()

	return rbt
}

//Create - create queue
// func (rbtq *RabbitQueue) Create(ch *amqp.Channel) amqp.Queue {
// 	q, err := ch.QueueDeclare(
// 		"hello", // name
// 		false,   // durable
// 		false,   // delete when unused
// 		false,   // exclusive
// 		false,   // no-wait
// 		nil,     // arguments
// 	)
// 	failOnError(err, "Failed to declare a queue")
// 	return q
// }
