package model

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)
func failError(err error, msg string) {
	if err != nil {

		log.Fatalf("%s: %s", msg, err)
	}
}
type Good struct {
	 goodid int
	 goodname string
}

type User struct {
	userid int
}

type Shop struct {
	shopid int
	goodid int
	goodnum int
	id int `gorm:"primary_key"`
}

type Order struct {
	Userid string
	Goodid string
	Shopid string
	Number string
}

type DbOrder struct {
	Userid string
	Goodid string
	Shopid string
	Buynum string
	Buytime string
}

type Goodnum struct {
	goodnum int
}

func C(userid string,shopid string,goodid string,buynum string) {
	db, err := gorm.Open("mysql", "root:191513@/rabbitmqdemo?charset=utf8&parseTime=True&loc=Local")
	failError(err,"数据库链接失败")
	defer db.Close()
	var order =DbOrder{Userid: userid, Shopid: shopid,Goodid:goodid,Buynum:buynum,Buytime:time.Now().Format("2006-01-02 15:04:05")}
	db.Create(&order)
}

func QueryCap(shopid int,goodid int) {
	db, err := gorm.Open("mysql", "root:191513@/rabbitmqdemo?charset=utf8")
	failError(err,"数据库链接失败")
	defer db.Close()
	var shop Shop
	db=db.Where(" shopid= ? AND goodid=?", shopid,goodid).Find(&shop)
	fmt.Println(db)
}

func UpdateCap(shopid int,goodid int,goodnum int) {
	db, err := gorm.Open("mysql", "root:191513@/rabbitmqdemo?charset=utf8&parseTime=True&loc=Local")
	failError(err,"数据库链接失败")
	defer db.Close()
	var shop Shop
	db=db.Model(&shop).Where("shopid = ? AND goodid = ?", shopid,goodid).Update("goodnum", goodnum)
    fmt.Println(db)
}



