package scoring

import (
	"math"
	"time"
)

func GiveScore(t time.Duration, timePreference float64, taskCoef float64, groupingCoef float64) float64 {
	timeScore := giveTimeScore(t, timePreference)
	return math.Round((timeScore * taskCoef * groupingCoef) * 100)/ 100
}