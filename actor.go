package statemachine

import (
	"github.com/rs/zerolog/log"
)

const UnknownState = "UNKNOWN"

type IActor interface {
	GetCurrentState() string

	setCurrentState(state string)

	FireEvent(event string, eventCtx EventContext)
}

type Actor struct {
	IActor
	id       string
	Executor *Executor
}

func NewActor(executor *Executor, id string) *Actor {
	if executor == nil {
		log.Error().Msg("executor null")
		return nil
	}
	var actor = Actor{
		id:       id,
		Executor: executor,
	}
	return &actor
}

func (actor *Actor) GetId() string {
	return actor.id
}

func (actor *Actor) GetExecutor() *Executor {
	return actor.Executor
}

func (actor *Actor) GetCurrentState() string {
	panic("Method GetCurrentState not supported. Please implement it.")
}

func (actor *Actor) setCurrentState(newState string) {
	panic("Method setCurrentState not supported. Please implement it.")
}

func (actor *Actor) FireEvent(event string, eventCtx EventContext) {
	transition, err := actor.Executor.GetStateTransition(actor.GetCurrentState(), event)
	if err != nil {
		log.Error().Stack().Err(err).
			Str("actor_id", actor.GetId()).
			Str("current_state", actor.GetCurrentState()).
			Msg("Caught error")
		return
	}

	actor.setCurrentState(transition.ToState)

	if transition.Handler != nil {
		transition.Handler.Execute(eventCtx)
	}
}
