package publshers

import (
	util "github.com/echenim/messagebrokers/rabbitmq/utils"
	"github.com/streadway/amqp"
)

//HelloWorld simple publisher for hello world
func HelloWorld() {

	con, er := util.ConnectRabbitMQ()
	util.ErrorHandler(er, " Failed to connect to Broker ")
	defer con.Close()

	//open channel
	ch, er := con.Channel()
	util.ErrorHandler(er, "unable to open channel")
	//defer ch.Close()

	q, er := ch.QueueDeclare("myron world", false, false, false, false, nil)
	util.ErrorHandler(er, "Queue declearation failed")
	body := "Hello k36tee its time to feed"
	er = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

}
