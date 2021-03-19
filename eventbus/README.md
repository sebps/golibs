# Eventbus
Some generic event bus implementation.

## Example
```go
package main

import (
	"fmt"
	"github.com/sebpsdev/golibs/generic/eventbus"
)

func main() {
  type Publisher struct{}
  type Subscriber struct{}
  
  func (s *Subscriber) HandleEvent(e eventbus.Event) (bool, string, string) {
    fmt.Println("event received : %s", e.name)
  }

  eb := &eventbus.Eventbus{}
  subscriber := &Subscriber{}
  publisher := &Publisher{}

  eb.Subscribe("EventName", subscriber)
  eb.Publish("EventName", publisher)

  // expect log : "event received : EventName"
}
```