package main

import (
	"ginchat/utils/rabbitmq"
	"log"
)

func main() {
	rabbitmq1, err1 := rabbitmq.NewRabbitMqSubscription("exchange.t1", "amqp://tiesheng:123456@192.168.0.201:5672/")
	if err1 != nil {
		log.Println(err1)
	}
	defer rabbitmq1.Destroy()
	go func() {
		msgs, err3 := rabbitmq1.Consume()
		if err3 != nil {
			log.Println(err3)
		}
		for {
			select {
			case d := <-msgs:
				log.Printf("接受到了：%s", d.Body)
			}
		}

		//for d := range msgs {
		//	log.Printf("接受到了：%s", d.Body)
		//}
	}()
	forever := make(chan bool)
	<-forever
}
