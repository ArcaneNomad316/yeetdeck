package scheduler

import (
	"math"
	"time"
)

const (
	MinEaseFactor     = 1.3
	DefaultEaseFactor = 2.5
)

// Review applies the SM-2 algorithm and returns updated scheduling values.
// quality: 0=complete fail, 3=correct w/ difficulty, 5=perfect
func Review(interval, easeFactor float64, reps, quality int) (newInterval, newEF float64, newReps int, nextDue time.Time) {
	// Update ease factor
	newEF = easeFactor + (0.1 - float64(5-quality)*(0.08+float64(5-quality)*0.02))
	if newEF < MinEaseFactor {
		newEF = MinEaseFactor
	}

	if quality < 3 {
		// Failed — reset
		newReps = 0
		newInterval = 1
	} else {
		// Passed
		switch reps {
		case 0:
			newInterval = 1
		case 1:
			newInterval = 6
		default:
			newInterval = math.Round(interval * newEF)
		}
		newReps = reps + 1
	}

	nextDue = time.Now().AddDate(0, 0, int(newInterval))
	return
}
