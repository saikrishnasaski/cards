package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	cards := newDeck()

	if len(cards) != 16 {

		t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {

		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
