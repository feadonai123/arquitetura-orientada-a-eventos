package broker

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	utils "cliente/utils"
)

const WHO_SUB = "broker sub"

type Subscriber struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	exchange   string
}

type EventHandler func(string)

func NewSubscriber(url string, exchange string) *Subscriber {
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
	utils.LogInfo("Conexão com o broker estabelecida", WHO_SUB)
	return &Subscriber{
		connection: conn,
		channel:    ch,
		exchange:   exchange,
	}
}

func (s *Subscriber) Subscribe(routingKey string, callback EventHandler) {
	q, err := s.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	utils.FailOnError(err, "Erro ao declarar fila")

	err = s.channel.QueueBind(
		q.Name,     // queue name
		routingKey, // routing key
		s.exchange, // exchange
		false,
		nil,
	)
	utils.FailOnError(err, "Erro ao realizar bind da fila")

	msgs, err := s.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Erro ao consumir mensagens")

	func() {
		for d := range msgs {
			utils.LogInfo(fmt.Sprintf("Mensagem recebida no tópico %s", routingKey), WHO_SUB)
			go callback(string(d.Body))
		}
	}()
}

func (s *Subscriber) Close() {
	if s.channel != nil {
		s.channel.Close()
	}
	if s.connection != nil {
		s.connection.Close()
	}
	utils.LogInfo("Conexão com o broker encerrada", WHO_SUB)
}