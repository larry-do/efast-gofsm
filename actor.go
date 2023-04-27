package statemachine

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"sync"
)

const UnknownState = "UNKNOWN"

type IActor interface {
	GetCurrentState() string

	setCurrentState(state string)

	FireEvent(event string, eventCtx EventContext)
}

type Actor struct {
	iactor       IActor
	mutex        sync.Mutex
	id           string
	currentState string
	executor     *Executor
}

func NewActor(executor *Executor) *Actor {
	if executor == nil {
		log.Error().Msg("executor null")
		return nil
	}
	var actor = Actor{
		id:           uuid.New().String(),
		currentState: UnknownState,
		executor:     executor,
	}
	return &actor
}

func (actor *Actor) GetCurrentState() string {
	return actor.currentState
}

func (actor *Actor) setCurrentState(newState string) {
	log.Info().Str("actor_id", actor.id).
		Str("previous_state", actor.currentState).
		Str("new_state", newState).
		Msg("Actor %s changed state from %s to %s\n")
	actor.currentState = newState
}

func (actor *Actor) FireEvent(event string, eventCtx EventContext) {
	actor.mutex.Lock()
	defer actor.mutex.Unlock()

	transition, err := actor.executor.getStateTransition(actor.currentState, event)
	if err != nil {
		log.Error().Stack().Err(err).
			Str("actor_id", actor.id).
			Str("current_state", actor.currentState).
			Msg("Caught error")
		return
	}

	actor.setCurrentState(transition.toState)

	if transition.handler != nil {
		transition.handler.Execute(eventCtx)
	}
}
