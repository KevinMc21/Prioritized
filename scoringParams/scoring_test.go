package scoringparams_test

import (
	scoringparams "Prioritized/v0/scoringParams"
	"Prioritized/v0/utils"
	"fmt"
	"testing"
	"time"
)

func TestTimeScore(t *testing.T) {
	tests := []struct{
		want		float64
		duration 	time.Duration
	}{	
		{
			// 5 minutes
			duration: time.Duration(300000000000),
			want: 51,
		},
		{
			// 30 minutes
			duration: time.Duration(1800000000000),
			want: 61,
		},
		{
			// 1 hour
			duration: time.Duration(3600000000000),
			want: 80,
		},
		{
			// 1 hour 30 minutes
			duration: time.Duration(5400000000000),
			want: 115,
		},
		{
			// 2 hours
			duration: time.Duration(7200000000000),
			want: 	175.0,
		},
		{
			// 2 hours 30 minutes
			duration: time.Duration(9000000000000),
			want: 275.0,
		},
	}

	for _, tt := range(tests) {
		
		ans := scoringparams.GiveTimeScore(tt.duration) 
		fmt.Printf("t: %v got score %v\n", tt.duration, ans)	
		if !utils.BetweenFloat64(tt.want - 5, tt.want + 5, ans) {
			t.Errorf("expected %v, got %v\n", tt.want, ans)
		}
	}
	
}