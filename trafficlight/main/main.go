package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
	"statemachine"
	"statemachine/trafficlight"
	"strings"
	"time"
)

func main() {
	configLogging(true)

	trafficlight.TrafficLight.FireEvent("START", nil)
	trafficlight.TrafficLight.FireEvent("SWITCH_ON", &statemachine.EventCtx{Data: "Turn on"})
	trafficlight.TrafficLight.FireEvent("SWITCH_OFF", &statemachine.EventCtx{Data: 1})
	trafficlight.TrafficLight.FireEvent("SWITCH_ON", &statemachine.EventCtx{Data: 2})
	trafficlight.TrafficLight.FireEvent("SWITCH_ON", &statemachine.EventCtx{Data: 3})
}

func configLogging(console bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	log.Logger = log.With().
		CallerWithSkipFrameCount(3).
		Int("pid", os.Getpid()).
		Logger()

	if console {
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
			NoColor:    false,
		}
		output.FormatLevel = func(i interface{}) string {
			switch i.(string) {
			case "info":
				return fmt.Sprintf("%5s:", color.Green.Render(strings.ToUpper(i.(string))))
			case "debug":
				return fmt.Sprintf("%5s:", color.Green.Render(strings.ToUpper(i.(string))))
			case "error":
				return fmt.Sprintf("%5s:", color.Red.Render(strings.ToUpper(i.(string))))
			case "fatal":
				return fmt.Sprintf("%5s:", color.Red.Render(strings.ToUpper(i.(string))))
			case "panic":
				return fmt.Sprintf("%5s:", color.Red.Render(strings.ToUpper(i.(string))))
			default:
				return fmt.Sprintf("%5s:", color.Green.Render(strings.ToUpper(i.(string))))
			}
		}
		output.FormatMessage = func(i interface{}) string {
			return fmt.Sprintf("| %s |", i)
		}
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s=", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s", i))
		}
		output.FormatTimestamp = func(i interface{}) string {
			unixTime, _ := i.(json.Number).Int64()
			return strings.ToUpper(fmt.Sprintf("%s",
				time.UnixMilli(unixTime).Format("2006-01-02 15:04:05.000")))
		}
		output.PartsExclude = []string{
			//zerolog.TimestampFieldName,
		}
		log.Logger = log.Output(output)
	}
}
