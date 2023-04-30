package statemachine

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type EventContext interface {
	GetData() any
}

type EventCtx struct {
	EventContext
	Data any
}

func (e *EventCtx) GetData() any {
	return e.Data
}

type Handler interface {
	Execute(eventCtx EventContext)
}

type StateTransition struct {
	fromState string
	eventName string
	toState   string
	handler   Handler
}

func NewStateTransition(fromState string, eventName string, toState string, handler Handler) *StateTransition {
	if &fromState == nil || len(fromState) < 1 {
		log.Error().Msg("fromState empty")
		return nil
	}
	if &eventName == nil || len(eventName) < 1 {
		log.Error().Msg("eventName empty")
		return nil
	}
	if &toState == nil || len(toState) < 1 {
		log.Error().Msg("toState empty")
		return nil
	}
	if &handler == nil {
		log.Error().Msg("handler empty")
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
	stateEventConfigs map[string]map[string]*StateTransition
}

func NewExecutor(transitions ...*StateTransition) *Executor {
	stateEventConfigs := make(map[string]map[string]*StateTransition)
	for i := 0; i < len(transitions); i++ {
		var eventTransitions map[string]*StateTransition
		var ok bool
		if eventTransitions, ok = stateEventConfigs[transitions[i].fromState]; !ok {
			eventTransitions = make(map[string]*StateTransition)
			stateEventConfigs[transitions[i].fromState] = eventTransitions
		}
		eventTransitions[transitions[i].eventName] = transitions[i]
	}

	return &Executor{
		stateEventConfigs: stateEventConfigs,
	}
}

func (executor *Executor) getStateTransition(fromState string, event string) (*StateTransition, error) {
	if ss, ok := executor.stateEventConfigs[fromState]; ok {
		if sss, oke := ss[event]; oke {
			return sss, nil
		}
	}
	return nil, errors.Errorf("not accept event %s at state %s", event, fromState)
}
