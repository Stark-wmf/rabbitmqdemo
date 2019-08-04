package messagequeue

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"rabbitmqdemoProject/model"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
func OpenConsumer() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Can't connect to MQ")
	defer conn.Close()
	amqpChannel, err := conn.Channel()
	handleError(err, "Can't create a amqpChannel")
	defer amqpChannel.Close()
	queue, err := amqpChannel.QueueDeclare("goodList", true, false, false,
		false, nil)
	handleError(err, "Could not declare `add` queue")
	err = amqpChannel.Qos(1, 0, false)
	handleError(err, "Could not configure QoS")
	messageChannel, err := amqpChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
//	4. 建⽴发布者
	handleError(err, "Could not register consumer")
	stopChan := make(chan bool)
	go func() {
		log.Printf("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			log.Printf("Received a message: %s", string(d.Body))
			good := &model.Order{}
			err := json.Unmarshal(d.Body, good)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			log.Printf("Good: %s", string(d.Body))
			if err := d.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()
	//终⽌当前进程
	<-stopChan
}