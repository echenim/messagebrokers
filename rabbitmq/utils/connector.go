package utils

import "github.com/streadway/amqp"

//ConnectRabbitMQ connection function
func ConnectRabbitMQ() (*amqp.Connection, error) {

	//conn, err := amqp.Dial("amqp://devng:103Solutionx@192.168.1.74:5672")
	return amqp.Dial("amqp://devng:103Solutionx@192.168.1.74:5672")
}
