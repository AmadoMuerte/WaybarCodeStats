package lang

import "github.com/AmadoMuerte/WaybarCodeStats/internal/models"

func GetTranslation(lang string) models.Translation {
	switch lang {
	case "en":
		return models.Translation{
			HourText:   "h",
			MinuteText: "m",
			Alt:        " of coding today",
		}
	case "ru":
		return models.Translation{
			HourText:   "ч",
			MinuteText: "м",
			Alt:        " кодинга сегодня",
		}
	}

	return models.Translation{}
}
