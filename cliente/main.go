package main

import (
	watcher "cliente/watcher"
	broker "cliente/broker"
)

var publisher *broker.Publisher

func main(){
	loadFlags()
	populateActionsMap(inputActions)
	populateEventsMap(inputEvents)

	done := make(chan bool)

	subscriber := broker.NewSubscriber(brokerURL, brokerExchange)
	publisher = broker.NewPublisher(brokerURL, brokerExchange)
	
	defer publisher.Close()
	defer subscriber.Close()

	for key, handler := range eventsMap {
		go subscriber.Subscribe(key, handler)
	}

	go watcher.Run(dirAction, dirRead, dirError, actionsMap)

	<-done
}