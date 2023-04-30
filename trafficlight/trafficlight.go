package trafficlight

import (
	"github.com/rs/zerolog/log"
	"statemachine"
)

type SwitchOffTransitionHandler struct {
	handler statemachine.Handler
}

func (t *SwitchOffTransitionHandler) Execute(eventCtx statemachine.EventContext) {
	log.Info().Any("data", eventCtx.GetData()).Msg("The light has been switched off.")
}

type SwitchOnTransitionHandler struct {
	handler statemachine.Handler
}

func (t *SwitchOnTransitionHandler) Execute(eventCtx statemachine.EventContext) {
	log.Info().Any("data", eventCtx.GetData()).Msg("The light has been switched on.")
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
