package main

import (
	"flag"
	"fmt"
	"os"

	"ankigo/storage"
	"ankigo/ui"
)

func main() {
	deckPath := flag.String("deck", "deck.json", "Path to deck file")
	flag.Parse()

	d, err := storage.Load(*deckPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading deck: %v\n", err)
		os.Exit(1)
	}

	ui.RunReview(d)

	if err := storage.Save(*deckPath, d); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving deck: %v\n", err)
		os.Exit(1)
	}
}
