package main

import (
	"github.com/AmadoMuerte/WakaTimeModule/internal/waybar/handlers/sendMessage"
	"github.com/AmadoMuerte/WakaTimeModule/internal/waybar/models"
)

func main() {
	msg := models.Message{
		Text:  "5 hours of coding ",
		Class: "custom-wakatime",
		Alt:   "<span foreground='#f1fa8c'>Golang</span>",
	}

	sendMessage.New(msg)
}
