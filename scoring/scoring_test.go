package scoring_test

import (
	"Prioritized/v0/scoring"
	"testing"
	"time"
)

func BenchmarkScore(b *testing.B) {
	tests := []struct{
		want		float64
		duration 	time.Duration
	}{	
		{
			// 5 minutes
			duration: time.Duration(300000000000),
		},
		{
			// 30 minutes
			duration: time.Duration(1800000000000),
		},
		{
			// 1 hour
			duration: time.Duration(3600000000000),
		},
		{
			// 1 hour 30 minutes
			duration: time.Duration(5400000000000),
		},
		{
			// 2 hours
			duration: time.Duration(7200000000000),
		},
		{
			// 2 hours 30 minutes
			duration: time.Duration(9000000000000),
		},
	}

	for _, tt := range(tests) {
		ans := scoring.GiveScore(tt.duration, 30, 1.2, 1.2) 
		b.Logf("t: %v got score %v with time preference of: %v\n", tt.duration, ans, 30)	
	}
	
}