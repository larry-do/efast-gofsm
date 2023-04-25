package statemachine

import (
	"github.com/google/uuid"
	"log"
	"sync"
)

const UnknownState = "UNKNOWN"

type StateType string

type IActor interface {
	GetCurrentState() StateType

	setCurrentState(state StateType)

	FireEvent(event EventType, eventCtx EventContext)
}

type Actor struct {
	iactor       IActor
	mutex        sync.Mutex
	id           string
	currentState StateType
	executor     *Executor
}

func NewActor(executor *Executor) *Actor {
	if executor == nil {
		log.Println("executor null")
		return nil
	}
	var actor = Actor{
		id:           uuid.New().String(),
		currentState: UnknownState,
		executor:     executor,
	}
	return &actor
}

func (actor *Actor) GetCurrentState() StateType {
	return actor.currentState
}

func (actor *Actor) setCurrentState(newState StateType) {
	log.Printf("changed state from %s to %s", actor.currentState, newState)
	actor.currentState = newState
}

func (actor *Actor) FireEvent(event EventType, eventCtx EventContext) {
	actor.mutex.Lock()
	defer actor.mutex.Unlock()

	transition, err := actor.executor.getStateTransition(actor.currentState, event)
	if err != nil {
		log.Println(err)
		return
	}

	actor.setCurrentState(transition.toState)

	if transition.handler != nil {
		transition.handler.Execute(eventCtx)
	}
}
