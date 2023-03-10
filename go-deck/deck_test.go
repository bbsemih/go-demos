package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Expected deck length of 12 but got : %v", len(d))
	}

	if d[0] != "Ace of\u00a0Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of\u00a0Hearts" {
		t.Errorf("Expected last card of Four of Hearts, but got %q", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Errorf("Expected 12 cards in deck, got: %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
