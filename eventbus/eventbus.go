package eventbus

import (
	"github.com/sebpsdev/golibs/generic/arrays"
	"github.com/sebpsdev/golibs/generic/maps"
	"time"
)

type Subscriber interface {
	HandleEvent(e Event) (success bool, status string, message string)
}

type Event struct {
	name      string
	payload   interface{}
	publisher interface{}
	timestamp int64
}

type Handling struct {
	event      *Event
	subscriber Subscriber
	success    bool
	status     string
	message    string
}

type EventBus struct {
	subscriptions map[string][]Subscriber
	eventStack    []*Event
	handlingStack []*Handling
	// TODO: future events
	// eventQueue    []*Event
}

func (eb *EventBus) Subscribe(eventName string, subscriber Subscriber) (bool, error) {
	if eb.subscriptions == nil {
		eb.subscriptions = make(map[string][]Subscriber)
	}
	if found, _ := maps.FindKey(eventName, eb.subscriptions); found {
		eb.subscriptions[eventName] = append(eb.subscriptions[eventName], subscriber)
	}

	eb.subscriptions[eventName] = []Subscriber{subscriber}

	return true, nil
}

func (eb *EventBus) Unsubscribe(eventName string, subscriber Subscriber) (bool, error) {
	if eb.subscriptions == nil {
		return true, nil
	}
	if foundEvent, _ := maps.FindKey(eventName, eb.subscriptions); foundEvent {
		if index, foundSubscriber, _ := arrays.Find(subscriber, eb.subscriptions[eventName]); foundSubscriber {
			tmp := make([]Subscriber, 0)
			tmp = append(tmp, eb.subscriptions[eventName][:index]...)
			eb.subscriptions[eventName] = append(tmp, eb.subscriptions[eventName][index+1:]...)
		}
	}

	return true, nil
}

func (eb *EventBus) Publish(eventName string, payload interface{}, publisher interface{}, timestamp int64) (bool, error) {
	// TODO: future events - Implement a scheduled publishing queue using timestamp argument
	event := &Event{eventName, payload, publisher, time.Now().Unix()}
	eb.eventStack = append(eb.eventStack, event)

	if found, _ := maps.FindKey(eventName, eb.subscriptions); found {
		for _, subscriber := range eb.subscriptions[eventName] {
			success, status, message := subscriber.HandleEvent(*event)
			eb.handlingStack = append(eb.handlingStack, &Handling{
				event, subscriber, success, status, message,
			})
		}
	}

	return true, nil
}
