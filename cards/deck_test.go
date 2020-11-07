package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // t is test handler
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubes" {
		t.Errorf("Expected last card of Four of Clubes , but got %v", d[len(d)-1])

	}

}

// Testing saveToFile and newDeckFromFile function/methods

// Create a deck => sav to file ==> file created==> Attempt to load file ==> crash
// Delete any files with _decktesting => create a deck ==> save to file ==> Load from file ==> Assert deck length
// ==> Delete any file with _decktesting
//  To remove file  os package  func Remove(name string)error

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
