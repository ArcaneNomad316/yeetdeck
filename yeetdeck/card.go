package deck

import "time"

type Card struct {
	ID         string    `json:"id"`
	Front      string    `json:"front"`
	Back       string    `json:"back"`
	Interval   float64   `json:"interval"`
	EaseFactor float64   `json:"ease_factor"`
	Reps       int       `json:"reps"`
	NextDue    time.Time `json:"next_due"`
}

func NewCard(id, front, back string) Card {
	return Card{
		ID:         id,
		Front:      front,
		Back:       back,
		Interval:   0,
		EaseFactor: 2.5,
		Reps:       0,
		NextDue:    time.Now(),
	}
}

func (c *Card) IsDue() bool {
	return !time.Now().Before(c.NextDue)
}
