package main

import (
	"github.com/echenim/messagebrokers/rabbitmq/utils"
	"github.com/streadway/amqp"
)

func main() {

	//initiate connection open
	con, er := amqp.Dial("amqp://devng:103Solutionx@192.168.1.74:5672")
	utils.ErrorHandler(er, " Failed to connect to Broker ")
	//defer con.Close()

	//open channel
	ch, er := con.Channel()
	utils.ErrorHandler(er, "unable to open channel")
	//defer ch.Close()

	q, er := ch.QueueDeclare("hello world", false, false, false, false, nil)
	utils.ErrorHandler(er, "Queue declearation failed")
	body := "Hello k36tee its time to feed"
	er = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

}
