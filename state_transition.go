package gofsm

import "github.com/rs/zerolog/log"

type Handler[C EventContext] interface {
	Execute(eventCtx *C)
}

type StateTransition[C EventContext] struct {
	FromState string
	EventName string
	ToState   string
	Handler   Handler[C]
}

func NewStateTransition[C EventContext](fromState string, eventName string, toState string, handler Handler[C]) *StateTransition[C] {
	if &fromState == nil {
		log.Error().Msg("fromState nil")
		return nil
	}
	if &eventName == nil || len(eventName) < 1 {
		log.Error().Msg("eventName empty")
		return nil
	}
	if &toState == nil {
		log.Error().Msg("toState nil")
		return nil
	}
	if &handler == nil {
		log.Error().Msg("handler nil")
		return nil
	}

	return &StateTransition[C]{
		FromState: fromState,
		EventName: eventName,
		ToState:   toState,
		Handler:   handler,
	}
}
