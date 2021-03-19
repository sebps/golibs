package eventbus

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type SubscribingStruct struct {
	receivedEvents []Event
}

func (s *SubscribingStruct) HandleEvent(e Event) (success bool, status string, message string) {
	s.receivedEvents = append(s.receivedEvents, e)
	return true, "processed", "event processed"
}

func TestPublish(t *testing.T) {
	var tests = []struct {
		have struct {
			eb *EventBus
		}
		want []*Event
	}{
		{
			have: struct {
				eb *EventBus
			}{
				eb: &EventBus{},
			},
			want: []*Event{
				&Event{
					name:      "event",
					payload:   nil,
					publisher: nil,
					timestamp: time.Now().Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%#v", tt.have)
		t.Run(testname, func(t *testing.T) {
			ok, err := tt.have.eb.Publish("event", nil, nil, time.Now().Unix())
			if err != nil {
				t.Errorf("got publish error %d", err)
			}
			if !ok {
				t.Errorf("publish not ok")
			}

			if !reflect.DeepEqual(tt.have.eb.eventStack, tt.want) {
				t.Errorf("got %#v, want %#v", tt.have.eb.eventStack, tt.want)
			}
		})
	}
}

func TestSubscribe(t *testing.T) {
	var tests = []struct {
		have struct {
			eb         *EventBus
			subscriber *SubscribingStruct
		}
		want []Event
	}{
		{
			have: struct {
				eb         *EventBus
				subscriber *SubscribingStruct
			}{
				eb:         &EventBus{},
				subscriber: &SubscribingStruct{},
			},
			want: []Event{
				Event{
					name:      "event",
					payload:   nil,
					publisher: nil,
					timestamp: time.Now().Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%#v", tt.have)
		t.Run(testname, func(t *testing.T) {
			ok, err := tt.have.eb.Subscribe("event", tt.have.subscriber)
			if err != nil {
				t.Errorf("got subscribe error %d", err)
			}
			if !ok {
				t.Errorf("subscribe not ok")
			}

			ok, err = tt.have.eb.Publish("event", nil, nil, time.Now().Unix())
			if err != nil {
				t.Errorf("got publish error %d", err)
			}
			if !ok {
				t.Errorf("publish not ok")
			}

			if !reflect.DeepEqual(tt.have.subscriber.receivedEvents, tt.want) {
				t.Errorf("got %#v, want %#v", tt.have.subscriber.receivedEvents, tt.want)
			}
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	var tests = []struct {
		have struct {
			eb         *EventBus
			subscriber *SubscribingStruct
		}
		want []Event
	}{
		{
			have: struct {
				eb         *EventBus
				subscriber *SubscribingStruct
			}{
				eb:         &EventBus{},
				subscriber: &SubscribingStruct{},
			},
			want: []Event{
				Event{
					name:      "event",
					payload:   nil,
					publisher: nil,
					timestamp: time.Now().Unix(),
				},
			},
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%#v", tt.have)
		t.Run(testname, func(t *testing.T) {
			ok, err := tt.have.eb.Subscribe("event", tt.have.subscriber)
			if err != nil {
				t.Errorf("got subscribe error %d", err)
			}
			if !ok {
				t.Errorf("subscribe not ok")
			}

			ok, err = tt.have.eb.Publish("event", nil, nil, time.Now().Unix())
			if err != nil {
				t.Errorf("got publish error %d", err)
			}
			if !ok {
				t.Errorf("publish not ok")
			}

			ok, err = tt.have.eb.Unsubscribe("event", tt.have.subscriber)
			if err != nil {
				t.Errorf("got unsubscribe error %d", err)
			}
			if !ok {
				t.Errorf("unsubscribe not ok")
			}

			ok, err = tt.have.eb.Publish("event", nil, nil, time.Now().Unix())
			if err != nil {
				t.Errorf("got publish error %d", err)
			}
			if !ok {
				t.Errorf("publish not ok")
			}

			if !reflect.DeepEqual(tt.have.subscriber.receivedEvents, tt.want) {
				t.Errorf("got %#v, want %#v", tt.have.subscriber.receivedEvents, tt.want)
			}
		})
	}
}
