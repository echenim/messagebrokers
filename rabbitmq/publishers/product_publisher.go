package publishers

import (
	util "github.com/echenim/messagebrokers/rabbitmq/utils"
	"github.com/streadway/amqp"
)

type product struct {
	Id       int64
	Thumb    []byte
	Name     string
	Group    string
	Quantity int
	Price    float64
}

func AvailableProducts() {
	con, er := util.ConnectRabbitMQ()
	util.ErrorHandler(er, "Failed to open connection")
	defer con.Close()

	ch, er := con.Channel()
	util.ErrorHandler(er, "Failed to open channel")
	defer ch.Close()

	q, er := ch.QueueDeclare("products", false, false, false, false, nil)
	util.ErrorHandler(er, "failed to declear Query")

	//create data defination
	var content product
	content.Id = 1
	content.Name = "BMW 8i"
	content.Group = "Exotic Cars"
	content.Price = 24500.45

	er = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "",
		Body:        []byte(""),
	})

}
	})

}



//ToJSON function to convert string to json
func ToJSON(p product) []byte {
	b, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("%s : %s", err, "Fail to convert to json")
	}
	return b
}

