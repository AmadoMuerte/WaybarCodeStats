package timeconv

import "github.com/AmadoMuerte/WaybarCodeStats/internal/models"
import "math"

func SecondToTime(second float64) models.Time {
	hours := math.Floor(second / 3600)
	remainingSeconds := math.Mod(second, 3600)
	minutes := math.Floor(remainingSeconds / 60)
	seconds := math.Mod(remainingSeconds, 60)

	return models.Time{
		Hour:   int(hours),
		Minute: int(minutes),
		Second: int(seconds),
	}
}
