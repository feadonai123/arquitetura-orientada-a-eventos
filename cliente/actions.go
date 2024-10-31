package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"

	utils "cliente/utils"
	watcher "cliente/watcher"
	mapper "cliente/mapper"
)

const WHO_ACTIONS = "actions"

type Action struct {
	ActionName       string
	EventName      string
	PrivateKey string
}

var actionsMap = make(map[string]watcher.FileHandler)

func populateActionsMap(filename string) {
	actions, err := loadActionsFromFile(filename)
	utils.FailOnError(err, "Erro ao carregar as ações")

	for _, action := range actions {
		actionsMap[action.ActionName] = func(content string) {
			executeGenericAction(content, privateKeys[action.PrivateKey], action.ActionName, action.EventName)
		}
	}
}

func executeGenericAction(content string, privateKey string, actionName string, eventName string) {
	utils.LogInfo(fmt.Sprintf("Executando %s", actionName), WHO_ACTIONS)
	message := mapper.MessageToWord(content, privateKey)
	publisher.Publish(eventName, message)
}

func loadActionsFromFile(filename string) ([]Action, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var actions []Action
	err = json.Unmarshal(bytes, &actions)
	if err != nil {
		return nil, err
	}

	return actions, nil
}