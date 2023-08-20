package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {

	cards := newDeck()

	if len(cards) != 16 {

		t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}
}
