package main

import (
	"encoding/json"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	ProductName string `json:"product_name"`
	Customer    string `json:"customer"`
	Quantity    int    `json:"quantity"`
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("RabbitMQ Bağlanti Hatasi:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Kanal Hatasi:", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"siparis_kuyrugu", 
		false, false, false, false, nil,
	)

	msgs, err := ch.Consume(
		q.Name, "", true, false, false, false, nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var order Order
			json.Unmarshal(d.Body, &order)

			log.Printf(" YENİ SİPARİŞ GELDİ: %s (%d adet) - %s", order.ProductName, order.Quantity, order.Customer)
			
			log.Println("  Fatura kesiliyor")
			time.Sleep(2 * time.Second) 
			
			log.Println("  Kargoya veriliyor")
			time.Sleep(2 * time.Second)

			log.Println("  İşlem Tamamlandi")
			log.Println("----------------------------------")
		}
	}()

	log.Printf("  İsci hazir. Siparis bekleniyor")
	<-forever
}