package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AmadoMuerte/WaybarCodeStats/internal/filereader"
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

	if args[2] == '' {
		fmt.Println("API key is not set.")
		return
	}
        apikey := args[2]

	// homeDir, err := os.UserHomeDir()
	// file, err := filereader.FindFile(filepath.Join(homeDir, ".wakatime.cfg"))
	// if err != nil {
	// 	fmt.Println("wakatime auth file is not exists", err)
	// 	os.Exit(1)
	// }

	token, err := apikey
	if err != nil {
		fmt.Println("wakatime auth token is not exists", err)
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
