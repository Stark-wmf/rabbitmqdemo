package messagequeue

import (
	"encoding/json"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"rabbitmqdemoProject/model"
	"time"
)
func failError(err error, msg string) {
	if err != nil {

		log.Fatalf("%s: %s", msg, err)
	}
}
//func OpenCreater() {
	//router := gin.Default()
	//router.POST("/order", func(c *gin.Context) {
	//	userid:= c.Query("userid")
	//	goodid := c.Query("goodid")
	//	shopid := c.Query("shopid")
	//	number := c.Query("number")
	//	Order(userid,shopid,goodid,number)
	//	go model.C(userid,shopid,goodid,number)
	//	c.JSON(http.StatusOK, gin.H{
	//		"status":  gin.H{
	//			"status_code": http.StatusOK,
	//			"status":      "ok",
	//		},
	//		"user-id": userid,
	//		"good-id": goodid,
	//		"shop-id": shopid,
	//		"number":  number,
	//	})
	//})
	//
	//router.Run()
	//}

	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//failError(err, "Can't connect to MQ")
	//defer conn.Close()
	//amqpChannel, err := conn.Channel()
	//failError(err, "Can't create a Channel")
	//defer amqpChannel.Close()
	//queue, err := amqpChannel.QueueDeclare("goodList", true, false, false,
	//	false, nil)
	//failError(err, "Could not declare queue")
	//rand.Seed(time.Now().UnixNano())
	//good := Order{Userid: rand.Intn(100), Shopid:1,
	//	Number:rand.Intn(99999), Goodid:2}
	//body, err := json.Marshal(good)
	//if err != nil {
	//	failError(err, "Error encoding JSON")
	//}
	//err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
	//	DeliveryMode: amqp.Persistent,
	//	ContentType: "text/plain",
	//	Body: body,
	//})
	//if err != nil {
	//	log.Fatalf("Error publishing message: %s", err)
	//}
	//log.Printf("AddGood: %s", string(body))
//}
func Order(userid string,shopid string,goodid string,number string){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	 failError(err, "Can't connect to MQ")
	defer conn.Close()
	amqpChannel, err := conn.Channel()
	failError(err, "Can't create a Channel")
	defer amqpChannel.Close()
	queue, err := amqpChannel.QueueDeclare("goodList", true, false, false,
		false, nil)
	failError(err, "Could not declare queue")
	rand.Seed(time.Now().UnixNano())
	good := model.Order{Userid:userid , Shopid:shopid,
		Number:number, Goodid:goodid}
	body, err := json.Marshal(good)
	if err != nil {
		failError(err, "Error encoding JSON")
	}
	err = amqpChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType: "text/plain",
		Body: body,
	})
	if err != nil {
		log.Fatalf("Error publishing message: %s", err)
	}
	log.Printf("AddGood: %s", string(body))
}



