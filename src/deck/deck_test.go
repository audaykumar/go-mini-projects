package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	val := 52
	if len(d) != val {
		t.Errorf("Expected deck of size %d, but got %v", val, len(d))
	}
}

func TestSaveToDeckAndNewDeckTestTestFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()

	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
