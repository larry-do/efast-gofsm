package gofsm

import (
	"github.com/rs/zerolog/log"
)

const (
	StateUnknown   = "FSM_UNKNOWN"
	StateAny       = "FSM_ANY"
	StateKeepState = "FSM_KEEP_STATE"
)

type IActor[C EventContext] interface {
	GetCurrentState() string

	setCurrentState(state string)

	FireEvent(event string, ctx *C)
}

type Actor[C EventContext] struct {
	IActor[C]
	id       string
	Executor *Executor[C]
}

func NewActor[C EventContext](executor *Executor[C], id string) *Actor[C] {
	if executor == nil {
		log.Error().Msg("executor null")
		return nil
	}
	var actor = Actor[C]{
		id:       id,
		Executor: executor,
	}
	return &actor
}

func (actor *Actor[C]) GetId() string {
	return actor.id
}

func (actor *Actor[C]) GetExecutor() *Executor[C] {
	return actor.Executor
}

func (actor *Actor[C]) GetCurrentState() string {
	panic("Method GetCurrentState not supported. Please implement it.")
}

func (actor *Actor[C]) setCurrentState(state string) {
	if state == StateKeepState {
		return
	}
	panic("Method setCurrentState not supported. Please implement it.")
}

func (actor *Actor[C]) FireEvent(event string, ctx *C) {
	transition, err := actor.Executor.GetStateTransition(actor.GetCurrentState(), event)
	if err != nil {
		log.Error().Stack().Err(err).
			Str("actor_id", actor.GetId()).
			Str("current_state", actor.GetCurrentState()).
			Msg("Caught error")
		return
	}

	if transition.Handler == nil {
		log.Error().
			Str("actor_id", actor.GetId()).
			Str("current_state", actor.GetCurrentState()).
			Msg("Not found handler for event.")
	}
	transition.Handler.Execute(ctx)

	actor.setCurrentState(transition.ToState)
}
