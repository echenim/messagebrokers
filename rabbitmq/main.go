package main

import (
	//sumer "github.com/echenim/messagebrokers/rabbitmq/consumers"
	pub "github.com/echenim/messagebrokers/rabbitmq/publisher"
)

func main() {

	pub.HelloWorld()
	pub.AvailableProducts()

	//sumer.HelloWorld()

}
