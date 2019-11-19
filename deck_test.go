package main

import (
  "testing"
  "os"
)

func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 16 {
    t.Errorf("Expected deck lenght of 16, but got %v", len(d))
  }
  if d[0] != "Ace of Spades" {
    t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
  }

  if d[len(d) - 1] != "Four of Clubs" {
    t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d) - 1])
  }
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
  os.Remove("_decktesting")

  d := newDeck()
  d.saveToFile("_decktesting")

  loadedDeck := deckFromFile("_decktesting")

  if len(loadedDeck) != 16 {
    t.Errorf("Expected deck lenght of 16, but got %v", len(loadedDeck))
  }

  os.Remove("_decktesting")
}
