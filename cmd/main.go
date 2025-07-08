package main

import (
	"fmt"
	"os"

	"github.com/AmadoMuerte/WaybarCodeStats/internal/lang"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/models"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/timeconv"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/wakatimeapi"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/waybar"
)

func main() {
	args := os.Args
	var argLang string
	if len(args) > 1 {
		argLang = args[1]
	} else {
		argLang = "en"
	}

	if len(args) <= 2 || args[2] == "" {
		fmt.Println("wakatime auth token is not exists")
		os.Exit(1)
	}
	
	token := args[2]

	api := wakatimeapi.Api{AuthToken: token}

	res := api.GetData()
	dailySeconds := res.Data.GrandTotal.TotalSeconds
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
