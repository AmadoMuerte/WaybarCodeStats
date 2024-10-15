package main

import (
	"fmt"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/filereader"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/lang"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/models"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/timeconv"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/wakatimeapi"
	"github.com/AmadoMuerte/WaybarCodeStats/internal/waybar"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	var argLang string
	if len(args) > 1 {
		argLang = args[1]
	} else {
		argLang = "en"
	}

	homeDir, err := os.UserHomeDir()
	file, err := filereader.FindFile(filepath.Join(homeDir, ".wakatime.cfg"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	token, err := filereader.ReadToken(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
