package scoring

import (
	"Prioritized/v0/utils"
	"time"
)

const MaximumTimeScore float64 = 1500

// Logistic function for modelling task difficulty over time, max at 2037
func giveTimeScore(t time.Duration, timePreference float64) float64 {
	minutesPassed := t.Minutes()

	if minutesPassed <= timePreference {
		score :=  minutesPassed*250/timePreference
		return score
	} else {
		score := minutesPassed*250/(timePreference/2) - timePreference*250/timePreference 
		return utils.MinF64([]float64{score, MaximumTimeScore})
	}
}
