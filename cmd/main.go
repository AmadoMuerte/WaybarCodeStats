package main

import (
	"fmt"
	"github.com/AmadoMuerte/WakaTimeModule/internal/lang"
	"github.com/AmadoMuerte/WakaTimeModule/internal/models"
	"github.com/AmadoMuerte/WakaTimeModule/internal/timeconv"
	"github.com/AmadoMuerte/WakaTimeModule/internal/wakatimeapi"
	"github.com/AmadoMuerte/WakaTimeModule/internal/waybar"
	"os"
)

func main() {
	args := os.Args
	var argLang string
	if len(args) > 1 {
		argLang = args[1]
	} else {
		argLang = "en"
	}

	api := wakatimeapi.Api{
		AuthToken: "TOKEN",
	}

	res := api.GetData()
	dailySeconds := res.Data.DailyAverage
	time := timeconv.SecondToTime(dailySeconds)
	translation := lang.GetTranslation(argLang)

	msg := models.Message{
		Text: fmt.Sprintf(
			"%d %s %d %s",
			time.Hour, translation.HourText, time.Minute, translation.MinuteText),
		Class: "custom-wakatime",
		Alt:   translation.Alt,
	}

	waybar.SetMessage(msg)
}
