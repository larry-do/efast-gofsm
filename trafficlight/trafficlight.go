package trafficlight

import (
	"fmt"
	"statemachine"
)

type SwitchOffTransitionHandler struct {
	handler statemachine.Handler
}

func (t *SwitchOffTransitionHandler) Execute(eventCtx statemachine.EventContext) {
	fmt.Printf("The light has been switched off. Data %s", eventCtx.GetData())
}

type SwitchOnTransitionHandler struct {
	handler statemachine.Handler
}

func (t *SwitchOnTransitionHandler) Execute(eventCtx statemachine.EventContext) {
	fmt.Printf("The light has been switched on. Data %s", eventCtx.GetData())
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
