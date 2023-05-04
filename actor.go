package statemachine

import (
	"github.com/rs/zerolog/log"
	goLog "log"
)

const UnknownState = "UNKNOWN"

type IActor interface {
	GetCurrentState() string

	setCurrentState(state string)

	FireEvent(event string, eventCtx EventContext)
}

type Actor struct {
	iactor   IActor
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
	goLog.Fatal("Not implemented GetCurrentState")
	return UnknownState
}

func (actor *Actor) setCurrentState(newState string) {
	log.Info().Str("actor_id", actor.GetId()).
		Str("previous_state", actor.GetCurrentState()).
		Str("new_state", newState).
		Msg("Actor state changed")

	goLog.Fatal("Not implemented setCurrentState")
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
