package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rabbitmqdemoProject/messagequeue"
	"rabbitmqdemoProject/model"
)
func main(){
	openOrderApi()

}
func openOrderApi(){
	router := gin.Default()
	router.POST("/order", func(c *gin.Context) {
		userid:= c.Query("userid")
		goodid := c.Query("goodid")
		shopid := c.Query("shopid")
		number := c.Query("number")
		messagequeue.Order(userid,shopid,goodid,number)
		go model.C(userid,shopid,goodid,number)
		c.JSON(http.StatusOK, gin.H{
			"status":  gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"user-id": userid,
			"good-id": goodid,
			"shop-id": shopid,
			"number":  number,
		})
	})

	router.Run()
}
