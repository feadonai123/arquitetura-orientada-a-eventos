package broker

import (
	"context"
	"time"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	utils "cliente/utils"
)

const WHO_PUB = "broker pub"

type Publisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	exchange   string
}

func NewPublisher(url string, exchange string) *Publisher {
	conn, err := amqp.Dial(url)
	utils.FailOnError(err, "Erro ao conectar ao broker")

	ch, err := conn.Channel()
	utils.FailOnError(err, "Erro ao criar canal")

	err = ch.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	utils.FailOnError(err, "Erro ao declarar exchange")
	utils.LogInfo("Conexão com o broker estabelecida", WHO_PUB)
	return &Publisher{
		connection: conn,
		channel:    ch,
		exchange:   exchange,
	}
}

func (p *Publisher) Publish(routingKey string, body string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	err := p.channel.PublishWithContext(ctx,
		p.exchange, // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Erro ao publicar mensagem")
	utils.LogInfo(fmt.Sprintf("Mensagem publicada no tópico %s", routingKey), WHO_PUB)
}

func (p *Publisher) Close() {
	if p.channel != nil {
		p.channel.Close()
	}
	if p.connection != nil {
		p.connection.Close()
	}
	utils.LogInfo("Conexão com o broker encerrada", WHO_PUB)
}