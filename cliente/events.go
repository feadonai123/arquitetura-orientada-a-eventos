package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"

	utils "cliente/utils"
	broker "cliente/broker"
	mapper "cliente/mapper"
)

const WHO_EVENTS = "events"

type Event struct {
	EventName      string
	PublicKey string
}

var eventsMap = make(map[string]broker.EventHandler)

func populateEventsMap(filename string) {
	events, err := loadEventsFromFile(filename)
	utils.FailOnError(err, "Erro ao carregar os eventos")

	for _, event := range events {
		eventsMap[event.EventName] = func(content string) {
			executeGenericEvent(content, publicKeys[event.PublicKey], event.EventName)
		}
	}
}

func executeGenericEvent(content string, publicKey string, eventName string) {
	utils.LogInfo(fmt.Sprintf("Executando %s", eventName), WHO_EVENTS)
	message, err := mapper.WordToMessage(content, publicKey)
	if err != nil {
		utils.LogError(err, "Erro ao decodificar a mensagem", WHO_EVENTS)
		return
	}

	utils.WriteFile(message, fmt.Sprintf("%s/%s.txt", dirInbox, utils.Timestamp()))
}


func loadEventsFromFile(filename string) ([]Event, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var events []Event
	err = json.Unmarshal(bytes, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}