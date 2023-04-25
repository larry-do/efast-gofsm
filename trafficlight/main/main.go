package main

import (
	"statemachine/trafficlight"
)

func main() {
	trafficlight.TrafficLight.FireEvent("START", nil)
	trafficlight.TrafficLight.FireEvent("SWITCH_ON", &trafficlight.EventCtx{Data: "lumix"})
}
