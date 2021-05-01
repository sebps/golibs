package eventbus

import (
	"github.com/sebps/golibs/generic/arrays"
	"github.com/sebps/golibs/generic/maps"
	"time"
)

type Subscriber interface {
	HandleEvent(e Event) (success bool, status string, message string)
}

type Event struct {
	Name      string
	Payload   interface{}
	Publisher interface{}
	Timestamp int64
}

type Handling struct {
	Event      *Event
	Subscriber Subscriber
	Success    bool
	Status     string
	Message    string
}

type EventBus struct {
	Subscriptions map[string][]Subscriber
	EventStack    []*Event
	HandlingStack []*Handling
	// TODO: future events
	// eventQueue    []*Event
}

func (eb *EventBus) Subscribe(eventName string, subscriber Subscriber) (bool, error) {
	if eb.Subscriptions == nil {
		eb.Subscriptions = make(map[string][]Subscriber)
	}
	if found, _ := maps.FindKey(eventName, eb.Subscriptions); found {
		eb.Subscriptions[eventName] = append(eb.Subscriptions[eventName], subscriber)
	}

	eb.Subscriptions[eventName] = []Subscriber{subscriber}

	return true, nil
}

func (eb *EventBus) Unsubscribe(eventName string, subscriber Subscriber) (bool, error) {
	if eb.Subscriptions == nil {
		return true, nil
	}
	if foundEvent, _ := maps.FindKey(eventName, eb.Subscriptions); foundEvent {
		if index, foundSubscriber, _ := arrays.Find(subscriber, eb.Subscriptions[eventName]); foundSubscriber {
			tmp := make([]Subscriber, 0)
			tmp = append(tmp, eb.Subscriptions[eventName][:index]...)
			eb.Subscriptions[eventName] = append(tmp, eb.Subscriptions[eventName][index+1:]...)
		}
	}

	return true, nil
}

func (eb *EventBus) Publish(eventName string, payload interface{}, publisher interface{}, timestamp int64) (bool, error) {
	// TODO: future events - Implement a scheduled publishing queue using timestamp argument
	event := &Event{eventName, payload, publisher, time.Now().Unix()}
	eb.EventStack = append(eb.EventStack, event)

	if found, _ := maps.FindKey(eventName, eb.Subscriptions); found {
		for _, subscriber := range eb.Subscriptions[eventName] {
			success, status, message := subscriber.HandleEvent(*event)
			eb.HandlingStack = append(eb.HandlingStack, &Handling{
				event, subscriber, success, status, message,
			})
		}
	}

	return true, nil
}
