package main

import (
	"ginchat/utils/rabbitmq"
	"log"
	"strconv"
)

const MQURL = "amqp://tiesheng:123456@192.168.0.201:5672/"

func main() {
	example3()
}
func example3() {
	rabbitmq1, err1 := rabbitmq.NewRabbitMqSubscription("exchange.t1", MQURL)
	defer rabbitmq1.Destroy()
	if err1 != nil {
		log.Println(err1)
	}
	//rabbitmq2, err2 := rabbitmq.NewRabbitMqSubscription("exchange.t1", MQURL)
	//defer rabbitmq2.Destroy()
	//if err2 != nil {
	//	log.Println(err2)
	//}

	go func() {
		for i := 0; i < 10000; i++ {
			rabbitmq1.Publish("消息：" + strconv.Itoa(i))
		}
	}()

	forever := make(chan bool)
	<-forever
}
