package main

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	app.Post("/order", func(c *fiber.Ctx) error {
		order := new(Order)
		if err := c.BodyParser(order); err != nil {
			return c.Status(400).SendString("Hatali Veri")
		}

		body, _ := json.Marshal(order)

		err = ch.Publish(
			"",     
			q.Name, 
			false,  
			false,  
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})

		if err != nil {
			return c.Status(500).SendString("Kuyruk Hatasi")
		}

		return c.JSON(fiber.Map{
			"message": "Siparişiniz alindi, hazirlaniyor!",
			"status":  "Processing",
		})
	})

	log.Fatal(app.Listen(":3000"))
}