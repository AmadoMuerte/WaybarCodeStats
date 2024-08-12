package timeconv

import "github.com/AmadoMuerte/WakaTimeModule/internal/models"

func SecondToTime(second int) models.Time {
	return models.Time{
		Hour:   second / 3600,
		Minute: (second % 3600) / 60,
		Second: second % 60,
	}
}
