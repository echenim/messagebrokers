package utils

import "github.com/streadway/amqp"

//ConnectRabbitMQ connection function
func ConnectRabbitMQ() (*amqp.Connection, error) {

	
	return amqp.Dial("amqp://devng:#########@****.****.***:5672")
}
