package fsm

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

	return &StateTransition[C]{
		FromState: fromState,
		EventName: eventName,
		ToState:   toState,
		Handler:   handler,
	}
}
