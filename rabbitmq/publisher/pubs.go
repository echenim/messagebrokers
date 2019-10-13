package publisher

import (
	"encoding/json"
	"log"

	ent "github.com/echenim/messagebrokers/rabbitmq/entity"
	util "github.com/echenim/messagebrokers/rabbitmq/utils"
	"github.com/streadway/amqp"
)

// type product struct {
// 	ID       int
// 	Thumb    []byte
// 	Name     string
// 	Group    string
// 	Quantity int
// 	Price    float64
// }

//HelloWorld simple publisher for hello world
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
	body := "Hello k36tee its time to feed"
	er = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

}

//AvailableProducts function for publishing available product
func AvailableProducts() {
	con, er := util.ConnectRabbitMQ()
	util.ErrorHandler(er, "Failed to open connection")

	//defer con.Close()

	ch, er := con.Channel()
	util.ErrorHandler(er, "Failed to open channel")
	ch.Confirm(true)

	//defer ch.Close()

	q, er := ch.QueueDeclare("products", false, false, false, false, nil)
	util.ErrorHandler(er, "failed to declear Query")

	//create data defination
	var content ent.Product
	content.ID = 5
	content.Name = "Ford"
	content.Group = "Exotic Cars"
	content.Price = 104500.45

	er = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "",
		Body:        toJSON(content),
	})
}

func toJSON(p ent.Product) []byte {
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("%s : %s", err, "Fail to convert to json")
	}
	return b
}
