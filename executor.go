package statemachine

import (
	"fmt"
	"log"
)

type EventType string

type EventContext interface {
	GetData() string
}

type Handler interface {
	Execute(eventCtx EventContext)
}

type StateTransition struct {
	fromState StateType
	eventName EventType
	toState   StateType
	handler   Handler
}

func NewStateTransition(fromState StateType, eventName EventType, toState StateType, handler Handler) *StateTransition {
	if &fromState == nil || len(fromState) < 1 {
		log.Println("fromState empty")
		return nil
	}
	if &eventName == nil || len(eventName) < 1 {
		log.Println("eventName empty")
		return nil
	}
	if &toState == nil || len(toState) < 1 {
		log.Println("toState empty")
		return nil
	}
	if &handler == nil {
		log.Println("handler empty")
		return nil
	}

	return &StateTransition{
		fromState: fromState,
		eventName: eventName,
		toState:   toState,
		handler:   handler,
	}
}

type Executor struct {
	stateEventConfigs map[StateType]map[EventType]*StateTransition
}

func NewExecutor(transitions ...*StateTransition) *Executor {
	stateEventConfigs := make(map[StateType]map[EventType]*StateTransition)
	for i := 0; i < len(transitions); i++ {
		var eventTransitions map[EventType]*StateTransition
		var ok bool
		if eventTransitions, ok = stateEventConfigs[transitions[i].fromState]; !ok {
			eventTransitions = make(map[EventType]*StateTransition)
			stateEventConfigs[transitions[i].fromState] = eventTransitions
		}
		eventTransitions[transitions[i].eventName] = transitions[i]
	}

	return &Executor{
		stateEventConfigs: stateEventConfigs,
	}
}

func (executor *Executor) getStateTransition(fromState StateType, event EventType) (*StateTransition, error) {
	if ss, ok := executor.stateEventConfigs[fromState]; ok {
		if sss, oke := ss[event]; oke {
			return sss, nil
		}
	}
	return nil, fmt.Errorf("not accept event %s at state %s", event, fromState)
}
