package deck

type Deck struct {
	Name  string `json:"name"`
	Cards []Card `json:"cards"`
}

func (d *Deck) DueCards() []*Card {
	var due []*Card
	for i := range d.Cards {
		if d.Cards[i].IsDue() {
			due = append(due, &d.Cards[i])
		}
	}
	return due
}

func (d *Deck) Stats() (total, due, new int) {
	total = len(d.Cards)
	for i := range d.Cards {
		if d.Cards[i].IsDue() {
			due++
		}
		if d.Cards[i].Reps == 0 {
			new++
		}
	}
	return
}
