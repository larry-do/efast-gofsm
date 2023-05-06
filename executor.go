package statemachine

import (
	"github.com/pkg/errors"
)

type Executor[C EventContext] struct {
	stateEventConfigs map[string]map[string]*StateTransition[C]
}

func NewExecutor[C EventContext](transitions ...*StateTransition[C]) *Executor[C] {
	stateEventConfigs := make(map[string]map[string]*StateTransition[C])
	for i := 0; i < len(transitions); i++ {
		var eventTransitions map[string]*StateTransition[C]
		var ok bool
		if eventTransitions, ok = stateEventConfigs[transitions[i].FromState]; !ok {
			eventTransitions = make(map[string]*StateTransition[C])
			stateEventConfigs[transitions[i].FromState] = eventTransitions
		}
		eventTransitions[transitions[i].EventName] = transitions[i]
	}

	return &Executor[C]{
		stateEventConfigs: stateEventConfigs,
	}
}

func (executor *Executor[C]) GetStateTransition(fromState string, event string) (*StateTransition[C], error) {
	if ss, ok := executor.stateEventConfigs[fromState]; ok {
		if sss, oke := ss[event]; oke {
			return sss, nil
		}
	}
	return nil, errors.Errorf("not accept event %s at state %s", event, fromState)
}
