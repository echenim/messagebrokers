package main

import (
	sumer "github.com/echenim/messagebrokers/rabbitmq/consumers"
	pub "github.com/echenim/messagebrokers/rabbitmq/publishers"
)

func main() {

	pub.HelloWorld()
	pub.HelloWorld()

	sumer.HelloWorld()

}
