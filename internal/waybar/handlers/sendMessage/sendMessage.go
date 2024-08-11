package sendMessage

import (
	"encoding/json"
	"fmt"
	"github.com/AmadoMuerte/WakaTimeModule/internal/waybar/models"
	"os"
)

func New(msg models.Message) {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsonBytes))
}
