# Iniciar container rabbitMQ
docker run -d --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:4.0-management

# Rodar projeto
go run main.go actions.go constants.go events.go

# Buildar projeto
go build -o meu_programa main.go actions.go constants.go events.go