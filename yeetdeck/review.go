package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"ankigo/deck"
	"ankigo/scheduler"
)

var reader = bufio.NewReader(os.Stdin)

func pause() {
	reader.ReadString('\n')
}

func prompt(msg string) string {
	fmt.Print(msg)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func RunReview(d *deck.Deck) {
	total, due, new_ := d.Stats()
	fmt.Printf("\n=== %s ===\n", d.Name)
	fmt.Printf("Total: %d | Due: %d | New: %d\n\n", total, due, new_)

	cards := d.DueCards()
	if len(cards) == 0 {
		fmt.Println("No cards due. Come back later!")
		return
	}

	for i, card := range cards {
		fmt.Printf("--- Card %d/%d ---\n", i+1, len(cards))
		fmt.Printf("FRONT:\n  %s\n\n", card.Front)
		fmt.Print("[ Press Enter to reveal ]")
		pause()

		fmt.Printf("BACK:\n  %s\n\n", card.Back)

		var quality int
		for {
			input := prompt("Rate (0=fail 1=hard 3=ok 5=easy): ")
			q, err := strconv.Atoi(input)
			if err == nil && q >= 0 && q <= 5 {
				quality = q
				break
			}
			fmt.Println("  Invalid input, enter 0-5")
		}

		newInterval, newEF, newReps, nextDue := scheduler.Review(
			card.Interval, card.EaseFactor, card.Reps, quality,
		)
		card.Interval = newInterval
		card.EaseFactor = newEF
		card.Reps = newReps
		card.NextDue = nextDue

		fmt.Printf("  → Next review in %.0f day(s)\n\n", newInterval)
	}

	fmt.Printf("Session complete! Reviewed %d card(s).\n", len(cards))
}
