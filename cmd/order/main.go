package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/luizhenriquees/go-order-manager/internal/infra/database"
	"github.com/luizhenriquees/go-order-manager/internal/usecase"
	"github.com/luizhenriquees/go-order-manager/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // espera tudo rodar e depois executa o close
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // escutando/lendo a fila do rabbitmq e alimenta um go channel numa T2
	rabbitmqWorker(msgRabbitmqChannel, uc)      // worker que também lê o channel, executa o usecase e salva no banco na T1
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitMQ worker...")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)

		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false) // Diz pro RabbitMQ que a mensagem foi processada
		fmt.Println("Mensagem processada e salva no banco:", output)
	}
}
