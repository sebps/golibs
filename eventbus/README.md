# Eventbus
Some generic event bus implementation providing event publishing / subscribing capabilities.

## Usage
The event bus allows any structure to subscribe to a labelled event as long as it implements the following Subscriber interface.
No specific interface is required for a structure to publish to the event bus.

```go
type Subscriber interface {
	HandleEvent(e Event) (success bool, status string, message string)
}
```

## Example

```go
  package main

  import (
    "fmt"
    "github.com/sebpsdev/golibs/eventbus"
    "time"
  )

  type Publisher struct{}
  type Subscriber struct{}

  func (s *Subscriber) HandleEvent(e eventbus.Event) (bool, string, string) {
    fmt.Printf("event received : %s\n", e.Name)
    return true, "success", "event processed"
  }

  func main() {
    eb := &eventbus.EventBus{}
    subscriber := &Subscriber{}
    publisher := &Publisher{}

    eb.Subscribe("EventName", subscriber)
    eb.Publish("EventName", nil, publisher, time.Now().Unix())
  }
  // expect log : "event received : EventName"
}
```