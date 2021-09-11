package scoringparams

import (
	"math"
	"time"
)

// Logistic function for modelling task difficulty over time, max at 2037
func GiveTimeScore(t time.Duration) float64 {
	return float64(2000 / ( 1 + math.Exp((2 * t.Minutes() - 500) * -0.01)) + 37)
}