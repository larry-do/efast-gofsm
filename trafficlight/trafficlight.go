package trafficlight

import (
	"github.com/rs/zerolog/log"
	"statemachine"
)

type SwitchOffTransitionHandler struct {
	handler statemachine.Handler
}

func (t *SwitchOffTransitionHandler) Execute(eventCtx statemachine.EventContext) {
	log.Info().Msgf("The light has been switched off. Data %s\n", eventCtx.GetData())
}

type SwitchOnTransitionHandler struct {
	handler statemachine.Handler
}

func (t *SwitchOnTransitionHandler) Execute(eventCtx statemachine.EventContext) {
	log.Info().Msgf("The light has been switched on. Data %s\n", eventCtx.GetData())
}

type EventCtx struct {
	statemachine.EventContext
	Data string
}

func (e *EventCtx) GetData() string {
	return e.Data
}

var (
	fromUnknownToOff = statemachine.NewStateTransition("UNKNOWN", "START", "OFF", nil)
	fromOnToOff      = statemachine.NewStateTransition("ON", "SWITCH_OFF", "OFF", &SwitchOffTransitionHandler{})
	fromOffToOn      = statemachine.NewStateTransition("OFF", "SWITCH_ON", "ON", &SwitchOnTransitionHandler{})
)

var (
	executor = statemachine.NewExecutor(fromUnknownToOff, fromOffToOn, fromOnToOff)
)

var (
	TrafficLight = statemachine.NewActor(executor)
)
