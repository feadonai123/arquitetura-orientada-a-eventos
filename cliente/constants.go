package main

import (
	"flag"

	utils "cliente/utils"
)

const dirAction = "./files/action"
const dirRead = "./files/actionRead"
const dirError = "./files/actionEror"
const dirInbox = "./files/inbox"

var inputActions = "./assets/input/actions.json"
var inputEvents = "./assets/input/events.json"

const brokerExchange = "logs_direct"
const brokerURL = "amqp://guest:guest@localhost:5672/"

var publicKeys = map[string]string{
	"CLIENTE": utils.ReadFile("./assets/key/cliente/public.pem"),
	"MOTOBOY": utils.ReadFile("./assets/key/motoboy/public.pem"),
	"RESTAURANTE": utils.ReadFile("./assets/key/restaurante/public.pem"),
	"TESOURARIA": utils.ReadFile("./assets/key/tesouraria/public.pem"),
}

var privateKeys = map[string]string{
	"CLIENTE": utils.ReadFile("./assets/key/cliente/private.pem"),
	"MOTOBOY": utils.ReadFile("./assets/key/motoboy/private.pem"),
	"RESTAURANTE": utils.ReadFile("./assets/key/restaurante/private.pem"),
	"TESOURARIA": utils.ReadFile("./assets/key/tesouraria/private.pem"),
}

func loadFlags() {
	_inputActions := flag.String("a", "./assets/input/actions.json", "Caminho para o arquivo de ações")
	_inputEvents := flag.String("e", "./assets/input/events.json", "Caminho para o arquivo de eventos")

	flag.Parse()

	inputActions = *_inputActions
	inputEvents = *_inputEvents

	utils.LogInfo("Flags carregadas", "main")
	utils.LogInfo("Arquivo de ações: " + inputActions, "main")
	utils.LogInfo("Arquivo de eventos: " + inputEvents, "main")
}