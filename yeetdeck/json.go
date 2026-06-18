package storage

import (
	"encoding/json"
	"os"

	"ankigo/deck"
)

func Load(path string) (*deck.Deck, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var d deck.Deck
	if err := json.Unmarshal(data, &d); err != nil {
		return nil, err
	}
	return &d, nil
}

func Save(path string, d *deck.Deck) error {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
