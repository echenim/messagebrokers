package consumers

import (
	"fmt"
	"log"

	util "github.com/echenim/messagebrokers/rabbitmq/utils"
)

//HelloWorld simple hello world consumer
func HelloWorld() {
	con, er := util.ConnectRabbitMQ()
	util.ErrorHandler(er, " Failed to connect to Broker ")
	defer con.Close()

	//open channel
	ch, er := con.Channel()
	util.ErrorHandler(er, "unable to open channel")
	defer ch.Close()

	q, er := ch.QueueDeclare("myron world", false, false, false, false, nil)
	util.ErrorHandler(er, "Queue declearation failed")
	msg, er := ch.Consume(q.Name, "", true, false, false, false, nil)
	util.ErrorHandler(er, "Failed to register a consumer")
	forever := make(chan bool)

	go func() {
		for d := range msg {
			fmt.Printf("\n %v\n", string(d.Body))
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
